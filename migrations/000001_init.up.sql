-- Final database schema with all migrations applied
-- This file contains the final state of all tables, views, functions, procedures, and triggers

-- ============================================
-- TABLES
-- ============================================

CREATE TABLE student (
                         student_id   SERIAL PRIMARY KEY,
                         name         TEXT,
                         surname      TEXT,
                         email        TEXT UNIQUE,
                         password     TEXT
);

CREATE TABLE teacher (
    teacher_id SERIAL PRIMARY KEY,
                         name       TEXT,
                         surname    TEXT,
                         email      TEXT UNIQUE,
                         password   TEXT
);

CREATE TABLE admin (
                       admin_id SERIAL PRIMARY KEY,
                       name     TEXT,
                       surname  TEXT,
                       email    TEXT UNIQUE,
    password TEXT
);

CREATE TABLE category (
    category_id SERIAL PRIMARY KEY,
                          name        TEXT
);

CREATE TABLE currency (
    currency_id SERIAL PRIMARY KEY,
                          name        TEXT
);

CREATE TABLE course (
                        course_id   SERIAL PRIMARY KEY,
                        name        TEXT,
    description TEXT,
    category_id INT,
                        start_date  DATE,
                        end_date    DATE,
                        price       INT CHECK (price >= 0),
    currency_id INT,
                        CHECK (start_date <= end_date),
                        CHECK (start_date >= CURRENT_DATE)
);

CREATE TABLE teachers_courses (
                                  course_id  INT,
    teacher_id INT,
    PRIMARY KEY (course_id, teacher_id)
);

CREATE TABLE course_lessons (
    lesson_id INT PRIMARY KEY,
    course_id INT
);

CREATE TABLE lessons_materials (
                                   lesson_id   INT,
    material_id INT,
    PRIMARY KEY (lesson_id, material_id)
);

CREATE TABLE lesson (
                        lesson_id   SERIAL PRIMARY KEY,
                        name        TEXT,
    description TEXT,
                        open_time   TIMESTAMPTZ,
                        video_link  TEXT,
    lesson_text TEXT
);

CREATE TABLE lesson_homeworks (
                                  lesson_id   INT,
    homework_id INT,
    PRIMARY KEY (lesson_id, homework_id)
);

CREATE TABLE homework (
                          homework_id   SERIAL PRIMARY KEY,
                          name          TEXT,
                          description   TEXT,
                          deadline_time TIMESTAMPTZ
);

CREATE TABLE homeworks_tasks (
                                 task_id     INT,
    homework_id INT,
    PRIMARY KEY (task_id, homework_id)
);

CREATE TABLE level (
    level_id SERIAL PRIMARY KEY,
                       name     TEXT
);

CREATE TABLE subcategory (
    subcategory_id SERIAL PRIMARY KEY,
                             name           TEXT,
                             category_id    INT
);

CREATE TABLE task (
                      task_id        SERIAL PRIMARY KEY,
                      type           TEXT,
                      description    TEXT,
                      right_answer   TEXT,
                      points         INT CHECK (points >= 0),
                      level_id       INT,
                      category_id    INT,
    subcategory_id INT
);

CREATE TABLE status_homework (
    status_homework_id SERIAL PRIMARY KEY,
                                 name               TEXT
);

CREATE TABLE status_answer (
    status_answer_id SERIAL PRIMARY KEY,
                               name             TEXT
);

CREATE TABLE student_answer (
                                student_answer_id  SERIAL PRIMARY KEY,
                                student_id         INT,
                                task_id            INT,
                                answer             TEXT,
                                status_answer_id   INT
);

CREATE TABLE homework_result (
                                 homework_result_id  SERIAL PRIMARY KEY,
                                 student_answer_id   INT,
                                 student_id          INT,
                                 task_id             INT,
                                 status_homework_id  INT,
                                 points              INT CHECK (points >= 0)
);

CREATE TABLE status_transaction (
    status_transaction_id SERIAL PRIMARY KEY,
                                    name                  TEXT
);

CREATE TABLE "transaction" (
    transaction_id SERIAL PRIMARY KEY,
                               student_id     INT,
                               status_id      INT,
                               total_price    INT CHECK (total_price >= 0)
);

CREATE TABLE transactions_courses (
    transaction_id INT,
                                      course_id      INT,
    PRIMARY KEY (transaction_id, course_id)
);

CREATE TABLE material (
    material_id SERIAL PRIMARY KEY,
                          source      TEXT,
                          extension   VARCHAR(7) CHECK (extension ~ '^[a-zA-Z0-9]{1,7}$'),
                          size        INT CHECK (size >= 0)
);

CREATE TABLE transaction_history (
                                     history_id     BIGSERIAL PRIMARY KEY,
                                     transaction_id INT,
                                     student_id     INT,
                                     status_id      INT,
                                     total_price    INT,
                                     operation      TEXT NOT NULL,
                                     changed_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
                                     changed_by     TEXT,
                                     old_row        JSONB,
                                     new_row        JSONB
);

-- ============================================
-- FOREIGN KEYS
-- ============================================

ALTER TABLE course ADD FOREIGN KEY (category_id) REFERENCES category(category_id);
ALTER TABLE course ADD FOREIGN KEY (currency_id) REFERENCES currency(currency_id);

ALTER TABLE teachers_courses 
    ADD FOREIGN KEY (teacher_id) REFERENCES teacher(teacher_id),
    ADD FOREIGN KEY (course_id) REFERENCES course(course_id),
    ADD CONSTRAINT one_teacher_per_course UNIQUE (course_id);

ALTER TABLE course_lessons 
    ADD FOREIGN KEY (course_id) REFERENCES course(course_id);
ALTER TABLE course_lessons 
    ADD FOREIGN KEY (lesson_id) REFERENCES lesson(lesson_id);

ALTER TABLE lessons_materials 
    ADD FOREIGN KEY (lesson_id) REFERENCES lesson(lesson_id),
    ADD FOREIGN KEY (material_id) REFERENCES material(material_id),
    ADD CONSTRAINT one_lesson_per_material UNIQUE (material_id);

ALTER TABLE lesson_homeworks 
    ADD FOREIGN KEY (lesson_id) REFERENCES lesson(lesson_id),
    ADD FOREIGN KEY (homework_id) REFERENCES homework(homework_id),
    ADD CONSTRAINT one_lesson_per_homework UNIQUE (homework_id);

ALTER TABLE homeworks_tasks 
    ADD FOREIGN KEY (homework_id) REFERENCES homework(homework_id);
ALTER TABLE homeworks_tasks 
    ADD FOREIGN KEY (task_id) REFERENCES task(task_id);

ALTER TABLE subcategory ADD FOREIGN KEY (category_id) REFERENCES category(category_id);

ALTER TABLE task 
    ADD FOREIGN KEY (level_id) REFERENCES level(level_id);
ALTER TABLE task 
    ADD FOREIGN KEY (category_id) REFERENCES category(category_id);
ALTER TABLE task 
    ADD FOREIGN KEY (subcategory_id) REFERENCES subcategory(subcategory_id);

ALTER TABLE student_answer 
    ADD FOREIGN KEY (student_id) REFERENCES student(student_id),
    ADD FOREIGN KEY (task_id) REFERENCES task(task_id),
    ADD FOREIGN KEY (status_answer_id) REFERENCES status_answer(status_answer_id);

ALTER TABLE homework_result 
    ADD FOREIGN KEY (student_answer_id) REFERENCES student_answer(student_answer_id),
    ADD FOREIGN KEY (student_id) REFERENCES student(student_id),
    ADD FOREIGN KEY (task_id) REFERENCES task(task_id),
    ADD FOREIGN KEY (status_homework_id) REFERENCES status_homework(status_homework_id);

ALTER TABLE "transaction" 
    ADD FOREIGN KEY (student_id) REFERENCES student(student_id),
    ADD FOREIGN KEY (status_id) REFERENCES status_transaction(status_transaction_id);

ALTER TABLE transactions_courses 
    ADD FOREIGN KEY (transaction_id) REFERENCES "transaction"(transaction_id),
    ADD FOREIGN KEY (course_id) REFERENCES course(course_id);

-- ============================================
-- VIEWS
-- ============================================

CREATE OR REPLACE VIEW v_transactions_report AS
SELECT
    t.transaction_id,
    t.student_id,
    s.name    AS student_name,
    s.surname AS student_surname,
    t.total_price,
    st.name   AS status_name,
    string_agg(DISTINCT c.name, ', ' ORDER BY c.name) AS courses_names
FROM "transaction" t
         LEFT JOIN student s
                   ON s.student_id = t.student_id
         LEFT JOIN status_transaction st
                   ON st.status_transaction_id = t.status_id
         LEFT JOIN transactions_courses tc
                   ON tc.transaction_id = t.transaction_id
         LEFT JOIN course c
                   ON c.course_id = tc.course_id
GROUP BY
    t.transaction_id,
    t.student_id,
    s.name,
    s.surname,
    t.total_price,
    st.name;

CREATE OR REPLACE VIEW v_student_category_stats AS
SELECT
    s.student_id,
    s.name       AS student_name,
    s.surname    AS student_surname,

    c.category_id,
    c.name       AS category_name,

    COUNT(DISTINCT hr.task_id)               AS solved_tasks_count,
    COALESCE(SUM(hr.points), 0)              AS points_earned,
    COALESCE(SUM(t.points), 0)              AS points_possible,
    CASE
        WHEN COALESCE(SUM(t.points), 0) = 0
            THEN 0
        ELSE ROUND(100.0 * COALESCE(SUM(hr.points), 0)::NUMERIC
                        / SUM(t.points), 2)
        END AS success_percent
FROM student s
         JOIN homework_result hr
              ON hr.student_id = s.student_id
         JOIN status_homework sh
              ON sh.status_homework_id = hr.status_homework_id
         JOIN task t
              ON t.task_id = hr.task_id
         LEFT JOIN category c
                   ON c.category_id = t.category_id
WHERE sh.name = 'Зачтено'
GROUP BY
    s.student_id,
    s.name,
    s.surname,
    c.category_id,
    c.name;

-- ============================================
-- FUNCTIONS
-- ============================================

CREATE OR REPLACE FUNCTION fn_log_transaction_history()
RETURNS TRIGGER AS
$$
BEGIN
    IF TG_OP = 'INSERT' THEN
        INSERT INTO transaction_history (
            transaction_id,
            student_id,
            status_id,
            total_price,
            operation,
            changed_at,
            changed_by,
            old_row,
            new_row
        )
        VALUES (
            NEW.transaction_id,
            NEW.student_id,
            NEW.status_id,
            NEW.total_price,
            'INSERT',
            now(),
            current_user,
            NULL,
            to_jsonb(NEW)
        );
        RETURN NEW;

    ELSIF TG_OP = 'UPDATE' THEN
        INSERT INTO transaction_history (
            transaction_id,
            student_id,
            status_id,
            total_price,
            operation,
            changed_at,
            changed_by,
            old_row,
            new_row
        )
        VALUES (
            NEW.transaction_id,
            NEW.student_id,
            NEW.status_id,
            NEW.total_price,
            'UPDATE',
            now(),
            current_user,
            to_jsonb(OLD),
            to_jsonb(NEW)
        );
        RETURN NEW;

    ELSIF TG_OP = 'DELETE' THEN
        INSERT INTO transaction_history (
            transaction_id,
            student_id,
            status_id,
            total_price,
            operation,
            changed_at,
            changed_by,
            old_row,
            new_row
        )
        VALUES (
            OLD.transaction_id,
            OLD.student_id,
            OLD.status_id,
            OLD.total_price,
            'DELETE',
            now(),
            current_user,
            to_jsonb(OLD),
            NULL
        );
        RETURN OLD;
    END IF;

    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE TYPE transaction_report_row AS (
    transaction_id  INT,
    student_id      INT,
    student_name    TEXT,
    student_surname TEXT,
    status_name     TEXT,
    total_price     INT,
    courses_names   TEXT
);

CREATE OR REPLACE FUNCTION get_transactions_report_dynamic(
    p_status_name   TEXT DEFAULT NULL,
    p_min_total     INT  DEFAULT NULL,
    p_max_total     INT  DEFAULT NULL
)
RETURNS SETOF transaction_report_row
LANGUAGE plpgsql
AS
$$
DECLARE
    v_sql TEXT;
BEGIN
    v_sql := '
        SELECT
            t.transaction_id,
            t.student_id,
            s.name    AS student_name,
            s.surname AS student_surname,
            st.name   AS status_name,
            t.total_price,
            string_agg(DISTINCT c.name, '', '' ORDER BY c.name) AS courses_names
        FROM "transaction" t
        LEFT JOIN student s
            ON s.student_id = t.student_id
        LEFT JOIN status_transaction st
            ON st.status_transaction_id = t.status_id
        LEFT JOIN transactions_courses tc
            ON tc.transaction_id = t.transaction_id
        LEFT JOIN course c
            ON c.course_id = tc.course_id
        WHERE 1=1
    ';

    IF p_status_name IS NOT NULL AND p_status_name != '' THEN
        v_sql := v_sql || ' AND st.name = ' || quote_literal(p_status_name);
    END IF;

    IF p_min_total IS NOT NULL AND p_min_total != 0 THEN
        v_sql := v_sql || ' AND t.total_price >= ' || p_min_total;
    END IF;

    IF p_max_total IS NOT NULL AND p_max_total != 0 THEN
        v_sql := v_sql || ' AND t.total_price <= ' || p_max_total;
    END IF;

    v_sql := v_sql || '
        GROUP BY
            t.transaction_id,
            t.student_id,
            s.name,
            s.surname,
            st.name,
            t.total_price
        ORDER BY t.transaction_id
    ';

    RETURN QUERY EXECUTE v_sql;
END;
$$;

-- ============================================
-- PROCEDURES
-- ============================================

CREATE OR REPLACE PROCEDURE bulk_update_transaction_status(
    IN  p_old_status_id INT,
    IN  p_new_status_id INT,
    IN  p_min_total     INT DEFAULT NULL,
    IN  p_max_total     INT DEFAULT NULL
)
LANGUAGE plpgsql
AS
$$
DECLARE
    v_sql           TEXT;
    v_affected_rows INT;
BEGIN
    v_sql := '
        UPDATE "transaction" t
        SET status_id = ' || p_new_status_id || '
        WHERE t.status_id = ' || p_old_status_id;

    IF p_min_total IS NOT NULL THEN
        v_sql := v_sql || ' AND t.total_price >= ' || p_min_total;
    END IF;

    IF p_max_total IS NOT NULL THEN
        v_sql := v_sql || ' AND t.total_price <= ' || p_max_total;
    END IF;

    EXECUTE v_sql;
    GET DIAGNOSTICS v_affected_rows = ROW_COUNT;

    RAISE NOTICE 'Updated % rows in "transaction"', v_affected_rows;
END;
$$;

-- ============================================
-- TRIGGERS
-- ============================================

CREATE TRIGGER trg_transaction_history
    AFTER INSERT OR UPDATE OR DELETE
    ON "transaction"
    FOR EACH ROW
    EXECUTE FUNCTION fn_log_transaction_history();

-- ============================================
-- ROLES AND PERMISSIONS
-- ============================================

-- Create roles (without login - these are group roles)
CREATE ROLE app_base    NOLOGIN;
CREATE ROLE app_admin   NOLOGIN;
CREATE ROLE app_teacher NOLOGIN;
CREATE ROLE app_student NOLOGIN;

-- Create login roles (users that can connect to database)
CREATE ROLE app_base_user
    LOGIN PASSWORD '123';

CREATE ROLE app_admin_user
    LOGIN PASSWORD '123';

CREATE ROLE app_teacher_user
    LOGIN PASSWORD '123';

CREATE ROLE app_student_user
    LOGIN PASSWORD '123';

-- Grant group roles to login roles
GRANT app_base    TO app_base_user;
GRANT app_admin   TO app_admin_user;
GRANT app_teacher TO app_teacher_user;
GRANT app_student TO app_student_user;

-- Set schema owner
ALTER SCHEMA public OWNER TO app_admin;

-- Grant usage on schema to all roles
GRANT USAGE ON SCHEMA public
    TO app_base, app_admin, app_teacher, app_student;

-- ============================================
-- BASE ROLE (for registration: POST /students, POST /teachers)
-- ============================================
GRANT SELECT, INSERT
    ON TABLE student, teacher
    TO app_base;

GRANT SELECT
    ON TABLE admin
    TO app_base;
-- ============================================
-- ADMIN ROLE (full access to all tables)
-- ============================================
GRANT SELECT, INSERT, UPDATE, DELETE
    ON TABLE
        student, teacher, admin,
        category, currency, course,
        teachers_courses, course_lessons, lessons_materials,
        lesson, lesson_homeworks, homework, homeworks_tasks,
        level, subcategory, task,
        status_homework, status_answer, student_answer,
        homework_result, status_transaction, "transaction",
        transactions_courses, material, transaction_history
    TO app_admin;

-- ============================================
-- TEACHER ROLE
-- ============================================

-- Read access (all tables)
GRANT SELECT
    ON TABLE
        student, teacher, admin,
        category, currency, course,
        teachers_courses, course_lessons, lessons_materials,
        lesson, lesson_homeworks, homework, homeworks_tasks,
        level, subcategory, task,
        status_homework, status_answer, student_answer,
        homework_result, status_transaction, "transaction",
        transactions_courses, material, transaction_history
    TO app_teacher;

-- Write access (can create/edit educational content)
GRANT SELECT, INSERT, UPDATE, DELETE
    ON TABLE
        category, currency, course,
        teachers_courses, course_lessons, lessons_materials,
        lesson, lesson_homeworks, homework, homeworks_tasks,
        level, subcategory, task,
        status_homework, status_answer,
        homework_result, material
    TO app_teacher;

-- Can update own teacher record
GRANT SELECT, UPDATE
    ON TABLE teacher
    TO app_teacher;

-- ============================================
-- STUDENT ROLE
-- ============================================

-- Read access (all tables)
GRANT SELECT
    ON TABLE
        student, teacher, admin,
        category, currency, course,
        teachers_courses, course_lessons, lessons_materials,
        lesson, lesson_homeworks, homework, homeworks_tasks,
        level, subcategory, task,
        status_homework, status_answer, student_answer,
        homework_result, status_transaction, "transaction",
        transactions_courses, material, transaction_history
    TO app_student;

-- Write access (can create/edit own answers and transactions)
GRANT SELECT, INSERT, UPDATE, DELETE
    ON TABLE
        student_answer,
        "transaction", transactions_courses
    TO app_student;

-- Can update own student record
GRANT SELECT, UPDATE
    ON TABLE student
    TO app_student;

-- ============================================
-- SEQUENCES (for SERIAL columns)
-- ============================================
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public
    TO app_base, app_admin, app_teacher, app_student;

-- Grant on future sequences
ALTER DEFAULT PRIVILEGES IN SCHEMA public
    GRANT USAGE, SELECT ON SEQUENCES
    TO app_base, app_admin, app_teacher, app_student;

-- ============================================
-- VIEWS
-- ============================================
GRANT SELECT
    ON v_transactions_report, v_student_category_stats
    TO app_admin, app_teacher, app_student;

-- ============================================
-- FUNCTIONS AND PROCEDURES
-- ============================================
-- Grant execute on functions and procedures
GRANT EXECUTE ON FUNCTION
    get_transactions_report_dynamic,
    fn_log_transaction_history
    TO app_admin, app_teacher, app_student;

GRANT EXECUTE ON PROCEDURE
    bulk_update_transaction_status
    TO app_admin;

