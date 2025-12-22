-- Drop triggers and functions
DROP TRIGGER IF EXISTS trg_check_homework_deadline_on_link ON lesson_homeworks;
DROP FUNCTION IF EXISTS check_homework_deadline_on_link();
DROP TRIGGER IF EXISTS trg_check_homework_deadline ON homework;
DROP FUNCTION IF EXISTS check_homework_deadline();

-- Revert homework.deadline_time back to TIME
ALTER TABLE homework ALTER COLUMN deadline_time TYPE TIME USING 
    CASE 
        WHEN deadline_time IS NOT NULL THEN 
            deadline_time::TIME
        ELSE NULL
    END;

-- Remove check constraint from course
ALTER TABLE course DROP CONSTRAINT IF EXISTS check_start_date_not_past;

-- Add bonus_amount back to student table
ALTER TABLE student ADD COLUMN bonus_amount INT CHECK (bonus_amount >= 0);

