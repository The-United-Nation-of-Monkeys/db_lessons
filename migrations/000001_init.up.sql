CREATE TABLE student (
                         student_id   SERIAL PRIMARY KEY,
                         name         TEXT,
                         surname      TEXT,
                         email        TEXT UNIQUE,
                         password     TEXT,
                         bonus_amount INT CHECK (bonus_amount >= 0)
);

CREATE TABLE teacher (
                         teacher_id SERIAL PRIMARY KEY,
                         name       TEXT,
                         surname    TEXT,
                         email      TEXT UNIQUE,
                         password   TEXT
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
                        CHECK (start_date <= end_date)
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
                          deadline_time TIME
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
                                status_homework_id INT,
                                points             INT CHECK (points >= 0)
);

CREATE TABLE homework_result (
                                 student_answer_id  SERIAL PRIMARY KEY,
                                 student_id         INT,
                                 task_id            INT,
                                 answer             TEXT,
                                 status_homework_id INT,
                                 points             INT CHECK (points >= 0)
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

CREATE OR REPLACE VIEW v_student_course_activity AS
SELECT
    st.student_id,
    st.name       AS student_name,
    st.surname    AS student_surname,
    c.course_id,
    c.name        AS course_name,
    c.description AS course_description,
    cat.name      AS category_name,
    cur.name      AS currency_name,
    c.price       AS course_price,
    COUNT(DISTINCT l.lesson_id)    AS lessons_count,
    COUNT(DISTINCT hw.homework_id) AS homeworks_count,
    COUNT(DISTINCT sa.task_id)     AS solved_tasks_count,
    COALESCE(SUM(sa.points), 0)    AS total_points_earned
FROM student st
         LEFT JOIN "transaction" tr
                   ON tr.student_id = st.student_id
         LEFT JOIN transactions_courses tc
                   ON tc.transaction_id = tr.transaction_id
         LEFT JOIN course c
                   ON c.course_id = tc.course_id
         LEFT JOIN category cat
                   ON cat.category_id = c.category_id
         LEFT JOIN currency cur
                   ON cur.currency_id = c.currency_id
         LEFT JOIN course_lessons cl
                   ON cl.course_id = c.course_id
         LEFT JOIN lesson l
                   ON l.lesson_id = cl.lesson_id
         LEFT JOIN lesson_homeworks lh
                   ON lh.lesson_id = l.lesson_id
         LEFT JOIN homework hw
                   ON hw.homework_id = lh.homework_id
         LEFT JOIN homeworks_tasks ht
                   ON ht.homework_id = hw.homework_id
         LEFT JOIN student_answer sa
                   ON sa.student_id = st.student_id
                       AND sa.task_id = ht.task_id
GROUP BY
    st.student_id,
    st.name,
    st.surname,
    c.course_id,
    c.name,
    c.description,
    cat.name,
    cur.name,
    c.price;


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

CREATE TRIGGER trg_transaction_history
    AFTER INSERT OR UPDATE OR DELETE
                    ON "transaction"
                        FOR EACH ROW
                        EXECUTE FUNCTION fn_log_transaction_history();


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

    IF p_max_total IS NOT NULL AND p_max_total !=0 THEN
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