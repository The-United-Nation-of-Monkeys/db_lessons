-- Remove bonus_amount from student table
ALTER TABLE student DROP COLUMN IF EXISTS bonus_amount;

-- Add check constraint to prevent past dates for course start_date
ALTER TABLE course ADD CONSTRAINT check_start_date_not_past 
    CHECK (start_date >= CURRENT_DATE);

-- Change homework.deadline_time from TIME to TIMESTAMP to allow date comparison
ALTER TABLE homework ALTER COLUMN deadline_time TYPE TIMESTAMPTZ USING 
    CASE 
        WHEN deadline_time IS NOT NULL THEN 
            (CURRENT_DATE + deadline_time)::TIMESTAMPTZ
        ELSE NULL
    END;

-- Create function to check homework deadline against course end date
CREATE OR REPLACE FUNCTION check_homework_deadline()
RETURNS TRIGGER AS $$
DECLARE
    v_course_end_date DATE;
    v_deadline_date DATE;
BEGIN
    -- Only check if deadline_time is provided
    IF NEW.deadline_time IS NULL THEN
        RETURN NEW;
    END IF;

    v_deadline_date := DATE(NEW.deadline_time);

    -- Get the course end_date for all lessons associated with this homework
    -- Check all courses to ensure deadline is not after any course end date
    FOR v_course_end_date IN
        SELECT DISTINCT c.end_date
        FROM lesson_homeworks lh
        JOIN course_lessons cl ON cl.lesson_id = lh.lesson_id
        JOIN course c ON c.course_id = cl.course_id
        WHERE lh.homework_id = NEW.homework_id
    LOOP
        -- Check if deadline_time is after course end_date
        IF v_deadline_date > v_course_end_date THEN
            RAISE EXCEPTION 'Homework deadline (%) cannot be after course end date (%). Course ends on %', 
                v_deadline_date, v_course_end_date, v_course_end_date;
        END IF;
    END LOOP;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger to check homework deadline when homework is updated
CREATE TRIGGER trg_check_homework_deadline
    BEFORE INSERT OR UPDATE ON homework
    FOR EACH ROW
    WHEN (NEW.deadline_time IS NOT NULL)
    EXECUTE FUNCTION check_homework_deadline();

-- Create function to check homework deadline when linking homework to lesson
CREATE OR REPLACE FUNCTION check_homework_deadline_on_link()
RETURNS TRIGGER AS $$
DECLARE
    v_course_end_date DATE;
    v_deadline_time TIMESTAMPTZ;
BEGIN
    -- Get homework deadline
    SELECT deadline_time INTO v_deadline_time
    FROM homework
    WHERE homework_id = NEW.homework_id;

    -- If no deadline set, no need to check
    IF v_deadline_time IS NULL THEN
        RETURN NEW;
    END IF;

    -- Get the course end_date for the lesson
    SELECT c.end_date INTO v_course_end_date
    FROM course_lessons cl
    JOIN course c ON c.course_id = cl.course_id
    WHERE cl.lesson_id = NEW.lesson_id
    LIMIT 1;

    -- If course exists, check deadline
    IF v_course_end_date IS NOT NULL THEN
        IF DATE(v_deadline_time) > v_course_end_date THEN
            RAISE EXCEPTION 'Homework deadline (%) cannot be after course end date (%). Course ends on %', 
                DATE(v_deadline_time), v_course_end_date, v_course_end_date;
        END IF;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger to check homework deadline when linking to lesson
CREATE TRIGGER trg_check_homework_deadline_on_link
    BEFORE INSERT ON lesson_homeworks
    FOR EACH ROW
    EXECUTE FUNCTION check_homework_deadline_on_link();

