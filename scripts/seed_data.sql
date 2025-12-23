-- Script to seed database with test data
-- Order matters due to foreign key constraints
--
-- Usage:
--   1. Via Makefile: make seed
--   2. Via script: ./scripts/seed.sh [database_name] [user] [host] [port]
--   3. Directly: psql -h localhost -U postgres -d online_classes -f scripts/seed_data.sql
--
-- Requirements:
--   - Database must be created and migrations must be run first
--   - All foreign key constraints must be satisfied
--   - Dates are set relative to CURRENT_DATE to satisfy constraints
--
-- Clear existing data (optional - uncomment if needed)
-- TRUNCATE TABLE transaction_history, transactions_courses, "transaction", homework_result, student_answer, 
--   homeworks_tasks, lesson_homeworks, lessons_materials, course_lessons, teachers_courses,
--   homework, task, material, lesson, course, subcategory, student, teacher, 
--   category, currency, level, status_homework, status_answer, status_transaction CASCADE;

-- 1. Insert reference data (no foreign keys)

-- Levels
INSERT INTO level (name) VALUES
    ('Начальный'),
    ('Средний'),
    ('Продвинутый'),
    ('Эксперт')
ON CONFLICT DO NOTHING;

-- Status Homework
INSERT INTO status_homework (name) VALUES
    ('Не проверено'),
    ('Зачтено'),
    ('Не зачтено'),
    ('На проверке')
ON CONFLICT DO NOTHING;

-- Status Answer
INSERT INTO status_answer (name) VALUES
    ('Правильно'),
    ('Неправильно'),
    ('Частично правильно'),
    ('Не проверено')
ON CONFLICT DO NOTHING;

-- Status Transaction
INSERT INTO status_transaction (name) VALUES
    ('В ожидании'),
    ('Завершено'),
    ('Отменено'),
    ('Ожидает оплаты')
ON CONFLICT DO NOTHING;

-- Categories
INSERT INTO category (name) VALUES
    ('Программирование'),
    ('Дизайн'),
    ('Аналитика'),
    ('Маркетинг'),
    ('Языки')
ON CONFLICT DO NOTHING;

-- Currency
INSERT INTO currency (name) VALUES
    ('RUB'),
    ('USD'),
    ('EUR'),
    ('Бонусы')
ON CONFLICT DO NOTHING;

-- 2. Insert data with dependencies

-- Subcategories
INSERT INTO subcategory (name, category_id) VALUES
    ('Фронтенд', 1),
    ('Бэкенд', 1),
    ('Мобильная разработка', 1),
    ('UI/UX', 2),
    ('Графический дизайн', 2),
    ('SQL', 3),
    ('Python для аналитики', 3),
    ('Excel', 3),
    ('SMM', 4),
    ('Контент-маркетинг', 4),
    ('Английский', 5),
    ('Немецкий', 5)
ON CONFLICT DO NOTHING;

-- Students (without bonus_amount) - 16 students
INSERT INTO student (name, surname, email, password) VALUES
    ('Анна', 'Ковалёва', 'anna@edu.ru', 'passanna'),
    ('Дмитрий', 'Лебедев', 'dmitry@edu.ru', 'passdmitry'),
    ('Елена', 'Соколова', 'elena@edu.ru', 'passelena'),
    ('Игорь', 'Козлов', 'igor@edu.ru', 'passigor'),
    ('Татьяна', 'Новикова', 'tanya@edu.ru', 'passtanya'),
    ('Александр', 'Морозов', 'alex@edu.ru', 'passalex'),
    ('Мария', 'Волкова', 'maria@edu.ru', 'passmaria'),
    ('Сергей', 'Семёнов', 'sergey@edu.ru', 'passsergey'),
    ('Ольга', 'Петрова', 'olga2@edu.ru', 'passolga'),
    ('Николай', 'Сидоров', 'nikolay@edu.ru', 'passnikolay'),
    ('Виктория', 'Фёдорова', 'viktoria@edu.ru', 'passviktoria'),
    ('Павел', 'Михайлов', 'pavel@edu.ru', 'passpavel'),
    ('Юлия', 'Александрова', 'yulia@edu.ru', 'passyulia'),
    ('Андрей', 'Дмитриев', 'andrey@edu.ru', 'passandrey'),
    ('Екатерина', 'Васильева', 'ekaterina@edu.ru', 'passekaterina'),
    ('Максим', 'Николаев', 'maxim@edu.ru', 'passmaxim')
ON CONFLICT (email) DO NOTHING;

-- Teachers - 12 teachers
INSERT INTO teacher (name, surname, email, password) VALUES
    ('Ирина', 'Миронова', 'irina@edu.ru', 'pass1'),
    ('Сергей', 'Петров', 'sergey@edu.ru', 'pass2'),
    ('Ольга', 'Смирнова', 'olga@edu.ru', 'pass3'),
    ('Алексей', 'Иванов', 'alex@edu.ru', 'pass4'),
    ('Марина', 'Кузнецова', 'marina@edu.ru', 'pass5'),
    ('Дмитрий', 'Попов', 'dmitry@edu.ru', 'pass6'),
    ('Елена', 'Соколова', 'elena.teacher@edu.ru', 'pass7'),
    ('Владимир', 'Новиков', 'vladimir@edu.ru', 'pass8'),
    ('Наталья', 'Козлова', 'natalya@edu.ru', 'pass9'),
    ('Роман', 'Лебедев', 'roman@edu.ru', 'pass10'),
    ('Анна', 'Волкова', 'anna.teacher@edu.ru', 'pass11'),
    ('Иван', 'Морозов', 'ivan@edu.ru', 'pass12')
ON CONFLICT (email) DO NOTHING;

-- Courses (start_date >= CURRENT_DATE) - 16 courses
INSERT INTO course (name, description, category_id, start_date, end_date, price, currency_id) VALUES
    ('Основы HTML и CSS', 'Изучение базовых технологий веб-разработки', 1, CURRENT_DATE + INTERVAL '7 days', CURRENT_DATE + INTERVAL '37 days', 9000, 1),
    ('Python для начинающих', 'Программирование с нуля', 1, CURRENT_DATE + INTERVAL '10 days', CURRENT_DATE + INTERVAL '50 days', 12000, 1),
    ('UX-дизайн интерфейсов', 'Создание удобных интерфейсов', 2, CURRENT_DATE + INTERVAL '5 days', CURRENT_DATE + INTERVAL '35 days', 15000, 1),
    ('Аналитика данных в SQL', 'Работа с базами данных', 3, CURRENT_DATE + INTERVAL '14 days', CURRENT_DATE + INTERVAL '44 days', 10000, 1),
    ('Английский для IT', 'Технический английский язык', 5, CURRENT_DATE + INTERVAL '3 days', CURRENT_DATE + INTERVAL '33 days', 8000, 1),
    ('React и современный фронтенд', 'Продвинутый курс по React', 1, CURRENT_DATE + INTERVAL '20 days', CURRENT_DATE + INTERVAL '60 days', 18000, 1),
    ('Go для бэкенда', 'Разработка серверных приложений', 1, CURRENT_DATE + INTERVAL '15 days', CURRENT_DATE + INTERVAL '55 days', 16000, 1),
    ('SMM стратегии', 'Продвижение в социальных сетях', 4, CURRENT_DATE + INTERVAL '12 days', CURRENT_DATE + INTERVAL '42 days', 11000, 1),
    ('JavaScript продвинутый', 'Продвинутые техники JavaScript', 1, CURRENT_DATE + INTERVAL '18 days', CURRENT_DATE + INTERVAL '58 days', 17000, 1),
    ('Node.js разработка', 'Серверная разработка на Node.js', 1, CURRENT_DATE + INTERVAL '22 days', CURRENT_DATE + INTERVAL '62 days', 19000, 1),
    ('Figma для дизайнеров', 'Профессиональная работа в Figma', 2, CURRENT_DATE + INTERVAL '8 days', CURRENT_DATE + INTERVAL '38 days', 14000, 1),
    ('Data Science с Python', 'Анализ данных и машинное обучение', 3, CURRENT_DATE + INTERVAL '16 days', CURRENT_DATE + INTERVAL '46 days', 22000, 1),
    ('Excel для аналитиков', 'Продвинутая работа с Excel', 3, CURRENT_DATE + INTERVAL '11 days', CURRENT_DATE + INTERVAL '41 days', 9500, 1),
    ('Контент-маркетинг', 'Создание контента для бизнеса', 4, CURRENT_DATE + INTERVAL '13 days', CURRENT_DATE + INTERVAL '43 days', 10500, 1),
    ('Немецкий для IT', 'Технический немецкий язык', 5, CURRENT_DATE + INTERVAL '6 days', CURRENT_DATE + INTERVAL '36 days', 8500, 1),
    ('Docker и контейнеризация', 'Работа с Docker и Kubernetes', 1, CURRENT_DATE + INTERVAL '25 days', CURRENT_DATE + INTERVAL '65 days', 20000, 1)
ON CONFLICT DO NOTHING;

-- Teachers-Courses relations
-- Using subqueries to get actual IDs from database by email and course name
INSERT INTO teachers_courses (course_id, teacher_id)
SELECT 
    c.course_id,
    t.teacher_id
FROM (VALUES
    ('Основы HTML и CSS', 'irina@edu.ru'),
    ('Python для начинающих', 'sergey@edu.ru'),
    ('UX-дизайн интерфейсов', 'olga@edu.ru'),
    ('Аналитика данных в SQL', 'alex@edu.ru'),
    ('Английский для IT', 'marina@edu.ru'),
    ('React и современный фронтенд', 'irina@edu.ru'),
    ('Go для бэкенда', 'sergey@edu.ru'),
    ('SMM стратегии', 'dmitry@edu.ru'),
    ('JavaScript продвинутый', 'elena.teacher@edu.ru'),
    ('Node.js разработка', 'vladimir@edu.ru'),
    ('Figma для дизайнеров', 'natalya@edu.ru'),
    ('Data Science с Python', 'roman@edu.ru'),
    ('Excel для аналитиков', 'anna.teacher@edu.ru'),
    ('Контент-маркетинг', 'ivan@edu.ru'),
    ('Немецкий для IT', 'marina@edu.ru'),
    ('Docker и контейнеризация', 'sergey@edu.ru')
) AS v(course_name, teacher_email)
JOIN course c ON c.name = v.course_name
JOIN teacher t ON t.email = v.teacher_email
ON CONFLICT DO NOTHING;

-- Lessons - 26 lessons
INSERT INTO lesson (name, description, open_time, video_link, lesson_text) VALUES
    ('Введение в HTML', 'Основы HTML', CURRENT_DATE + INTERVAL '7 days' + TIME '10:00:00', 'http://example.com/html_intro', 'Текст урока HTML'),
    ('CSS-базовый', 'Оформление страниц', CURRENT_DATE + INTERVAL '10 days' + TIME '10:00:00', 'http://example.com/css', 'Текст урока CSS'),
    ('Основы Python', 'Знакомство с Python', CURRENT_DATE + INTERVAL '10 days' + TIME '14:00:00', 'http://example.com/python_intro', 'Текст урока Python'),
    ('Переменные и типы данных', 'Работа с данными в Python', CURRENT_DATE + INTERVAL '15 days' + TIME '14:00:00', 'http://example.com/python_vars', 'Текст урока о переменных'),
    ('Введение в UX', 'Основы пользовательского опыта', CURRENT_DATE + INTERVAL '5 days' + TIME '11:00:00', 'http://example.com/ux_intro', 'Текст урока UX'),
    ('Прототипирование', 'Создание прототипов', CURRENT_DATE + INTERVAL '12 days' + TIME '11:00:00', 'http://example.com/prototyping', 'Текст урока прототипирования'),
    ('SELECT запросы', 'Основы SQL', CURRENT_DATE + INTERVAL '14 days' + TIME '13:00:00', 'http://example.com/sql_select', 'Текст урока SQL'),
    ('JOIN операции', 'Соединение таблиц', CURRENT_DATE + INTERVAL '21 days' + TIME '13:00:00', 'http://example.com/sql_join', 'Текст урока JOIN'),
    ('IT Vocabulary', 'Словарь IT терминов', CURRENT_DATE + INTERVAL '3 days' + TIME '09:00:00', 'http://example.com/english_vocab', 'Текст урока английского'),
    ('Technical Writing', 'Техническое письмо', CURRENT_DATE + INTERVAL '10 days' + TIME '09:00:00', 'http://example.com/english_writing', 'Текст урока письма'),
    ('React Components', 'Компоненты React', CURRENT_DATE + INTERVAL '20 days' + TIME '15:00:00', 'http://example.com/react_components', 'Текст урока React'),
    ('Go Basics', 'Основы Go', CURRENT_DATE + INTERVAL '15 days' + TIME '16:00:00', 'http://example.com/go_basics', 'Текст урока Go'),
    ('SMM Tools', 'Инструменты SMM', CURRENT_DATE + INTERVAL '12 days' + TIME '12:00:00', 'http://example.com/smm_tools', 'Текст урока SMM'),
    ('JavaScript ES6+', 'Современный JavaScript', CURRENT_DATE + INTERVAL '18 days' + TIME '14:00:00', 'http://example.com/js_es6', 'Текст урока ES6'),
    ('Асинхронность в JS', 'Promises и async/await', CURRENT_DATE + INTERVAL '20 days' + TIME '14:00:00', 'http://example.com/js_async', 'Текст урока асинхронности'),
    ('Node.js основы', 'Введение в Node.js', CURRENT_DATE + INTERVAL '22 days' + TIME '16:00:00', 'http://example.com/node_basics', 'Текст урока Node.js'),
    ('Express.js', 'Фреймворк Express', CURRENT_DATE + INTERVAL '25 days' + TIME '16:00:00', 'http://example.com/express', 'Текст урока Express'),
    ('Figma основы', 'Работа в Figma', CURRENT_DATE + INTERVAL '8 days' + TIME '10:00:00', 'http://example.com/figma_basics', 'Текст урока Figma'),
    ('Компоненты в Figma', 'Создание компонентов', CURRENT_DATE + INTERVAL '12 days' + TIME '10:00:00', 'http://example.com/figma_components', 'Текст урока компонентов'),
    ('Pandas для анализа', 'Библиотека Pandas', CURRENT_DATE + INTERVAL '16 days' + TIME '15:00:00', 'http://example.com/pandas', 'Текст урока Pandas'),
    ('Визуализация данных', 'Графики и диаграммы', CURRENT_DATE + INTERVAL '19 days' + TIME '15:00:00', 'http://example.com/visualization', 'Текст урока визуализации'),
    ('Excel формулы', 'Продвинутые формулы', CURRENT_DATE + INTERVAL '11 days' + TIME '13:00:00', 'http://example.com/excel_formulas', 'Текст урока формул'),
    ('Создание контента', 'Основы контент-маркетинга', CURRENT_DATE + INTERVAL '13 days' + TIME '11:00:00', 'http://example.com/content', 'Текст урока контента'),
    ('Немецкий базовый', 'Основы немецкого языка', CURRENT_DATE + INTERVAL '6 days' + TIME '09:00:00', 'http://example.com/german_basics', 'Текст урока немецкого'),
    ('Docker основы', 'Введение в Docker', CURRENT_DATE + INTERVAL '25 days' + TIME '17:00:00', 'http://example.com/docker_basics', 'Текст урока Docker'),
    ('Kubernetes', 'Оркестрация контейнеров', CURRENT_DATE + INTERVAL '30 days' + TIME '17:00:00', 'http://example.com/kubernetes', 'Текст урока Kubernetes')
ON CONFLICT DO NOTHING;

-- Course-Lessons relations
INSERT INTO course_lessons (lesson_id, course_id)
SELECT 
    l.lesson_id,
    c.course_id
FROM (VALUES
    ('Введение в HTML', 'Основы HTML и CSS'),
    ('CSS-базовый', 'Основы HTML и CSS'),
    ('Основы Python', 'Python для начинающих'),
    ('Переменные и типы данных', 'Python для начинающих'),
    ('Введение в UX', 'UX-дизайн интерфейсов'),
    ('Прототипирование', 'UX-дизайн интерфейсов'),
    ('SELECT запросы', 'Аналитика данных в SQL'),
    ('JOIN операции', 'Аналитика данных в SQL'),
    ('IT Vocabulary', 'Английский для IT'),
    ('Technical Writing', 'Английский для IT'),
    ('React Components', 'React и современный фронтенд'),
    ('Go Basics', 'Go для бэкенда'),
    ('SMM Tools', 'SMM стратегии'),
    ('JavaScript ES6+', 'JavaScript продвинутый'),
    ('Асинхронность в JS', 'JavaScript продвинутый'),
    ('Node.js основы', 'Node.js разработка'),
    ('Express.js', 'Node.js разработка'),
    ('Figma основы', 'Figma для дизайнеров'),
    ('Компоненты в Figma', 'Figma для дизайнеров'),
    ('Pandas для анализа', 'Data Science с Python'),
    ('Визуализация данных', 'Data Science с Python'),
    ('Excel формулы', 'Excel для аналитиков'),
    ('Создание контента', 'Контент-маркетинг'),
    ('Немецкий базовый', 'Немецкий для IT'),
    ('Docker основы', 'Docker и контейнеризация'),
    ('Kubernetes', 'Docker и контейнеризация')
) AS v(lesson_name, course_name)
JOIN lesson l ON l.name = v.lesson_name
JOIN course c ON c.name = v.course_name
ON CONFLICT DO NOTHING;

-- Materials - 20 materials
INSERT INTO material (source, extension, size) VALUES
    ('http://example.com/html_cheatsheet.pdf', 'pdf', 500),
    ('http://example.com/css_guide.pdf', 'pdf', 700),
    ('http://example.com/python_intro.ppt', 'ppt', 1024),
    ('http://example.com/python_exercises.zip', 'zip', 2048),
    ('http://example.com/ux_templates.sketch', 'sketch', 1500),
    ('http://example.com/sql_reference.pdf', 'pdf', 800),
    ('http://example.com/english_glossary.pdf', 'pdf', 600),
    ('http://example.com/react_tutorial.pdf', 'pdf', 1200),
    ('http://example.com/go_examples.zip', 'zip', 1800),
    ('http://example.com/smm_calendar.xlsx', 'xlsx', 400),
    ('http://example.com/js_handbook.pdf', 'pdf', 900),
    ('http://example.com/node_guide.pdf', 'pdf', 1100),
    ('http://example.com/express_api.zip', 'zip', 1600),
    ('http://example.com/figma_templates.sketch', 'sketch', 1300),
    ('http://example.com/pandas_tutorial.pdf', 'pdf', 1000),
    ('http://example.com/matplotlib_examples.zip', 'zip', 1900),
    ('http://example.com/excel_functions.xlsx', 'xlsx', 450),
    ('http://example.com/content_strategy.pdf', 'pdf', 750),
    ('http://example.com/german_vocab.pdf', 'pdf', 550),
    ('http://example.com/docker_guide.pdf', 'pdf', 1400)
ON CONFLICT DO NOTHING;

-- Lessons-Materials relations
INSERT INTO lessons_materials (lesson_id, material_id)
SELECT 
    l.lesson_id,
    m.material_id
FROM (VALUES
    ('Введение в HTML', 'http://example.com/html_cheatsheet.pdf'),
    ('CSS-базовый', 'http://example.com/css_guide.pdf'),
    ('Основы Python', 'http://example.com/python_intro.ppt'),
    ('Переменные и типы данных', 'http://example.com/python_exercises.zip'),
    ('Введение в UX', 'http://example.com/ux_templates.sketch'),
    ('SELECT запросы', 'http://example.com/sql_reference.pdf'),
    ('IT Vocabulary', 'http://example.com/english_glossary.pdf'),
    ('React Components', 'http://example.com/react_tutorial.pdf'),
    ('Go Basics', 'http://example.com/go_examples.zip'),
    ('SMM Tools', 'http://example.com/smm_calendar.xlsx'),
    ('JavaScript ES6+', 'http://example.com/js_handbook.pdf'),
    ('Node.js основы', 'http://example.com/node_guide.pdf'),
    ('Express.js', 'http://example.com/express_api.zip'),
    ('Figma основы', 'http://example.com/figma_templates.sketch'),
    ('Компоненты в Figma', 'http://example.com/figma_templates.sketch'),
    ('Pandas для анализа', 'http://example.com/pandas_tutorial.pdf'),
    ('Визуализация данных', 'http://example.com/matplotlib_examples.zip'),
    ('Excel формулы', 'http://example.com/excel_functions.xlsx'),
    ('Создание контента', 'http://example.com/content_strategy.pdf'),
    ('Немецкий базовый', 'http://example.com/german_vocab.pdf'),
    ('Docker основы', 'http://example.com/docker_guide.pdf'),
    ('Kubernetes', 'http://example.com/docker_guide.pdf')
) AS v(lesson_name, material_source)
JOIN lesson l ON l.name = v.lesson_name
JOIN material m ON m.source = v.material_source
ON CONFLICT DO NOTHING;

-- Homeworks (deadline_time as TIMESTAMPTZ, must be <= course end_date) - 26 homeworks
-- Course 1 ends: CURRENT_DATE + 37 days, Course 2 ends: CURRENT_DATE + 50 days
-- Course 3 ends: CURRENT_DATE + 35 days, Course 4 ends: CURRENT_DATE + 44 days
-- Course 5 ends: CURRENT_DATE + 33 days, Course 6 ends: CURRENT_DATE + 60 days
-- Course 7 ends: CURRENT_DATE + 55 days, Course 8 ends: CURRENT_DATE + 42 days
-- Course 9 ends: CURRENT_DATE + 58 days, Course 10 ends: CURRENT_DATE + 62 days
-- Course 11 ends: CURRENT_DATE + 38 days, Course 12 ends: CURRENT_DATE + 46 days
-- Course 13 ends: CURRENT_DATE + 41 days, Course 14 ends: CURRENT_DATE + 43 days
-- Course 15 ends: CURRENT_DATE + 36 days, Course 16 ends: CURRENT_DATE + 65 days
INSERT INTO homework (name, description, deadline_time) VALUES
    ('Задание по HTML', 'Сделать свою первую страницу', (CURRENT_DATE + INTERVAL '12 days')::TIMESTAMPTZ + TIME '18:00:00'),
    ('Задание по CSS', 'Стилизовать страницу', (CURRENT_DATE + INTERVAL '15 days')::TIMESTAMPTZ + TIME '20:00:00'),
    ('Задание по Python', 'Написать простой скрипт', (CURRENT_DATE + INTERVAL '18 days')::TIMESTAMPTZ + TIME '22:00:00'),
    ('Задание по переменным', 'Работа с типами данных', (CURRENT_DATE + INTERVAL '22 days')::TIMESTAMPTZ + TIME '20:00:00'),
    ('UX анализ', 'Проанализировать интерфейс', (CURRENT_DATE + INTERVAL '10 days')::TIMESTAMPTZ + TIME '19:00:00'),
    ('Прототип приложения', 'Создать прототип', (CURRENT_DATE + INTERVAL '18 days')::TIMESTAMPTZ + TIME '21:00:00'),
    ('SQL запросы', 'Написать SELECT запросы', (CURRENT_DATE + INTERVAL '20 days')::TIMESTAMPTZ + TIME '18:00:00'),
    ('JOIN практика', 'Соединить таблицы', (CURRENT_DATE + INTERVAL '28 days')::TIMESTAMPTZ + TIME '19:00:00'),
    ('Перевод текста', 'Перевести технический текст', (CURRENT_DATE + INTERVAL '8 days')::TIMESTAMPTZ + TIME '17:00:00'),
    ('Написать email', 'Составить деловое письмо', (CURRENT_DATE + INTERVAL '15 days')::TIMESTAMPTZ + TIME '18:00:00'),
    ('React компонент', 'Создать компонент', (CURRENT_DATE + INTERVAL '25 days')::TIMESTAMPTZ + TIME '20:00:00'),
    ('Go программа', 'Написать простую программу', (CURRENT_DATE + INTERVAL '20 days')::TIMESTAMPTZ + TIME '21:00:00'),
    ('SMM план', 'Составить план публикаций', (CURRENT_DATE + INTERVAL '18 days')::TIMESTAMPTZ + TIME '19:00:00'),
    ('ES6 функции', 'Использовать стрелочные функции', (CURRENT_DATE + INTERVAL '22 days')::TIMESTAMPTZ + TIME '19:00:00'),
    ('Async/await практика', 'Работа с асинхронным кодом', (CURRENT_DATE + INTERVAL '24 days')::TIMESTAMPTZ + TIME '20:00:00'),
    ('Node.js сервер', 'Создать простой сервер', (CURRENT_DATE + INTERVAL '26 days')::TIMESTAMPTZ + TIME '21:00:00'),
    ('Express API', 'Создать REST API', (CURRENT_DATE + INTERVAL '30 days')::TIMESTAMPTZ + TIME '22:00:00'),
    ('Figma макет', 'Создать макет в Figma', (CURRENT_DATE + INTERVAL '13 days')::TIMESTAMPTZ + TIME '18:00:00'),
    ('Компонентная система', 'Создать систему компонентов', (CURRENT_DATE + INTERVAL '17 days')::TIMESTAMPTZ + TIME '19:00:00'),
    ('Анализ данных', 'Проанализировать датасет', (CURRENT_DATE + INTERVAL '21 days')::TIMESTAMPTZ + TIME '20:00:00'),
    ('Визуализация', 'Создать графики', (CURRENT_DATE + INTERVAL '24 days')::TIMESTAMPTZ + TIME '21:00:00'),
    ('Excel отчет', 'Создать отчет в Excel', (CURRENT_DATE + INTERVAL '16 days')::TIMESTAMPTZ + TIME '18:00:00'),
    ('Контент-план', 'Составить контент-план', (CURRENT_DATE + INTERVAL '18 days')::TIMESTAMPTZ + TIME '19:00:00'),
    ('Немецкий диалог', 'Составить диалог', (CURRENT_DATE + INTERVAL '11 days')::TIMESTAMPTZ + TIME '17:00:00'),
    ('Docker контейнер', 'Создать Dockerfile', (CURRENT_DATE + INTERVAL '30 days')::TIMESTAMPTZ + TIME '22:00:00'),
    ('Kubernetes деплой', 'Настроить деплоймент', (CURRENT_DATE + INTERVAL '35 days')::TIMESTAMPTZ + TIME '23:00:00')
ON CONFLICT DO NOTHING;

-- Lesson-Homeworks relations
INSERT INTO lesson_homeworks (lesson_id, homework_id)
SELECT 
    l.lesson_id,
    h.homework_id
FROM (VALUES
    ('Введение в HTML', 'Задание по HTML'),
    ('CSS-базовый', 'Задание по CSS'),
    ('Основы Python', 'Задание по Python'),
    ('Переменные и типы данных', 'Задание по переменным'),
    ('Введение в UX', 'UX анализ'),
    ('Прототипирование', 'Прототип приложения'),
    ('SELECT запросы', 'SQL запросы'),
    ('JOIN операции', 'JOIN практика'),
    ('IT Vocabulary', 'Перевод текста'),
    ('Technical Writing', 'Написать email'),
    ('React Components', 'React компонент'),
    ('Go Basics', 'Go программа'),
    ('SMM Tools', 'SMM план'),
    ('JavaScript ES6+', 'ES6 функции'),
    ('Асинхронность в JS', 'Async/await практика'),
    ('Node.js основы', 'Node.js сервер'),
    ('Express.js', 'Express API'),
    ('Figma основы', 'Figma макет'),
    ('Компоненты в Figma', 'Компонентная система'),
    ('Pandas для анализа', 'Анализ данных'),
    ('Визуализация данных', 'Визуализация'),
    ('Excel формулы', 'Excel отчет'),
    ('Создание контента', 'Контент-план'),
    ('Немецкий базовый', 'Немецкий диалог'),
    ('Docker основы', 'Docker контейнер'),
    ('Kubernetes', 'Kubernetes деплой')
) AS v(lesson_name, homework_name)
JOIN lesson l ON l.name = v.lesson_name
JOIN homework h ON h.name = v.homework_name
ON CONFLICT DO NOTHING;

-- Tasks - 26 tasks
INSERT INTO task (type, description, right_answer, points, level_id, category_id, subcategory_id) VALUES
    ('тест', 'Какой тег создаёт абзац?', '<p>', 5, 1, 1, 1),
    ('тест', 'Какой язык используется для стилизации?', 'CSS', 5, 1, 1, 1),
    ('код', 'Напиши функцию, которая выводит "Hello"', 'def hello(): print("Hello")', 10, 2, 1, 2),
    ('тест', 'Какой тип данных используется для целых чисел?', 'int', 5, 1, 1, 2),
    ('тест', 'Что такое UX?', 'User Experience', 5, 1, 2, 4),
    ('практика', 'Создай прототип кнопки', 'prototype', 15, 2, 2, 4),
    ('тест', 'Какой оператор используется для выборки данных?', 'SELECT', 5, 1, 3, 6),
    ('код', 'Напиши запрос с JOIN', 'SELECT * FROM a JOIN b', 15, 2, 3, 6),
    ('тест', 'Переведи "database"', 'база данных', 5, 1, 5, 11),
    ('практика', 'Напиши деловое письмо', 'email', 10, 2, 5, 11),
    ('код', 'Создай React компонент', 'function Component() {}', 20, 3, 1, 1),
    ('код', 'Напиши Go программу', 'package main', 15, 2, 1, 2),
    ('практика', 'Составь контент-план', 'plan', 10, 2, 4, 9),
    ('тест', 'Что такое деструктуризация?', 'destructuring', 5, 2, 1, 1),
    ('код', 'Используй Promise для асинхронности', 'new Promise()', 15, 3, 1, 1),
    ('код', 'Создай HTTP сервер на Node.js', 'http.createServer()', 20, 2, 1, 2),
    ('код', 'Создай роут в Express', 'app.get()', 15, 2, 1, 2),
    ('практика', 'Создай иконку в Figma', 'icon', 10, 1, 2, 4),
    ('практика', 'Создай стиль в Figma', 'style', 12, 2, 2, 4),
    ('код', 'Загрузи данные через Pandas', 'pd.read_csv()', 15, 2, 3, 7),
    ('код', 'Создай график через matplotlib', 'plt.plot()', 18, 3, 3, 7),
    ('практика', 'Создай сводную таблицу', 'pivot', 12, 2, 3, 8),
    ('практика', 'Напиши статью для блога', 'article', 15, 2, 4, 10),
    ('тест', 'Переведи "computer" на немецкий', 'Computer', 5, 1, 5, 12),
    ('код', 'Создай Dockerfile', 'FROM node', 20, 3, 1, 2),
    ('код', 'Создай Deployment в Kubernetes', 'kubectl apply', 25, 4, 1, 2)
ON CONFLICT DO NOTHING;

-- Homeworks-Tasks relations
INSERT INTO homeworks_tasks (task_id, homework_id)
SELECT 
    t.task_id,
    h.homework_id
FROM (VALUES
    ('Какой тег создаёт абзац?', 'Задание по HTML'),
    ('Какой язык используется для стилизации?', 'Задание по CSS'),
    ('Напиши функцию, которая выводит "Hello"', 'Задание по Python'),
    ('Какой тип данных используется для целых чисел?', 'Задание по переменным'),
    ('Что такое UX?', 'UX анализ'),
    ('Создай прототип кнопки', 'Прототип приложения'),
    ('Какой оператор используется для выборки данных?', 'SQL запросы'),
    ('Напиши запрос с JOIN', 'JOIN практика'),
    ('Переведи "database"', 'Перевод текста'),
    ('Напиши деловое письмо', 'Написать email'),
    ('Создай React компонент', 'React компонент'),
    ('Напиши Go программу', 'Go программа'),
    ('Составь контент-план', 'SMM план'),
    ('Что такое деструктуризация?', 'ES6 функции'),
    ('Используй Promise для асинхронности', 'Async/await практика'),
    ('Создай HTTP сервер на Node.js', 'Node.js сервер'),
    ('Создай роут в Express', 'Express API'),
    ('Создай иконку в Figma', 'Figma макет'),
    ('Создай стиль в Figma', 'Компонентная система'),
    ('Загрузи данные через Pandas', 'Анализ данных'),
    ('Создай график через matplotlib', 'Визуализация'),
    ('Создай сводную таблицу', 'Excel отчет'),
    ('Напиши статью для блога', 'Контент-план'),
    ('Переведи "computer" на немецкий', 'Немецкий диалог'),
    ('Создай Dockerfile', 'Docker контейнер'),
    ('Создай Deployment в Kubernetes', 'Kubernetes деплой')
) AS v(task_description, homework_name)
JOIN task t ON t.description = v.task_description
JOIN homework h ON h.name = v.homework_name
ON CONFLICT DO NOTHING;

-- Student Answers (points are stored only in task table) - 26 answers
INSERT INTO student_answer (student_id, task_id, answer, status_answer_id)
SELECT 
    s.student_id,
    t.task_id,
    v.answer,
    v.status_answer_id
FROM (VALUES
    ('anna@edu.ru', 'Какой тег создаёт абзац?', '<p>', 1),
    ('anna@edu.ru', 'Какой язык используется для стилизации?', 'CSS', 1),
    ('dmitry@edu.ru', 'Напиши функцию, которая выводит "Hello"', 'def hello(): print("Hello")', 1),
    ('dmitry@edu.ru', 'Какой тип данных используется для целых чисел?', 'int', 1),
    ('elena@edu.ru', 'Что такое UX?', 'User Experience', 1),
    ('elena@edu.ru', 'Создай прототип кнопки', 'prototype', 1),
    ('igor@edu.ru', 'Какой оператор используется для выборки данных?', 'SELECT', 1),
    ('igor@edu.ru', 'Напиши запрос с JOIN', 'SELECT * FROM a JOIN b', 1),
    ('tanya@edu.ru', 'Переведи "database"', 'база данных', 1),
    ('tanya@edu.ru', 'Напиши деловое письмо', 'email', 1),
    ('anna@edu.ru', 'Создай React компонент', 'function Component() {}', 4),
    ('dmitry@edu.ru', 'Напиши Go программу', 'package main', 1),
    ('elena@edu.ru', 'Составь контент-план', 'plan', 1),
    ('alex@edu.ru', 'Что такое деструктуризация?', 'destructuring', 1),
    ('maria@edu.ru', 'Используй Promise для асинхронности', 'new Promise()', 1),
    ('sergey@edu.ru', 'Создай HTTP сервер на Node.js', 'http.createServer()', 1),
    ('olga2@edu.ru', 'Создай роут в Express', 'app.get()', 1),
    ('nikolay@edu.ru', 'Создай иконку в Figma', 'icon', 1),
    ('viktoria@edu.ru', 'Создай стиль в Figma', 'style', 1),
    ('pavel@edu.ru', 'Загрузи данные через Pandas', 'pd.read_csv()', 1),
    ('yulia@edu.ru', 'Создай график через matplotlib', 'plt.plot()', 1),
    ('andrey@edu.ru', 'Создай сводную таблицу', 'pivot', 1),
    ('ekaterina@edu.ru', 'Напиши статью для блога', 'article', 1),
    ('maxim@edu.ru', 'Переведи "computer" на немецкий', 'Computer', 1),
    ('anna@edu.ru', 'Создай Dockerfile', 'FROM node', 4),
    ('dmitry@edu.ru', 'Создай Deployment в Kubernetes', 'kubectl apply', 1)
) AS v(student_email, task_description, answer, status_answer_id)
JOIN student s ON s.email = v.student_email
JOIN task t ON t.description = v.task_description
ON CONFLICT DO NOTHING;

-- Homework Results - 24 results
INSERT INTO homework_result (student_answer_id, student_id, task_id, status_homework_id, points)
SELECT 
    sa.student_answer_id,
    s.student_id,
    t.task_id,
    v.status_homework_id,
    v.points
FROM (VALUES
    ('anna@edu.ru', 'Какой тег создаёт абзац?', '<p>', 2, 5),
    ('anna@edu.ru', 'Какой язык используется для стилизации?', 'CSS', 2, 5),
    ('dmitry@edu.ru', 'Напиши функцию, которая выводит "Hello"', 'def hello(): print("Hello")', 2, 10),
    ('dmitry@edu.ru', 'Какой тип данных используется для целых чисел?', 'int', 2, 5),
    ('elena@edu.ru', 'Что такое UX?', 'User Experience', 2, 5),
    ('elena@edu.ru', 'Создай прототип кнопки', 'prototype', 2, 15),
    ('igor@edu.ru', 'Какой оператор используется для выборки данных?', 'SELECT', 2, 5),
    ('igor@edu.ru', 'Напиши запрос с JOIN', 'SELECT * FROM a JOIN b', 2, 15),
    ('tanya@edu.ru', 'Переведи "database"', 'база данных', 2, 5),
    ('tanya@edu.ru', 'Напиши деловое письмо', 'email', 2, 10),
    ('dmitry@edu.ru', 'Напиши Go программу', 'package main', 2, 15),
    ('elena@edu.ru', 'Составь контент-план', 'plan', 2, 10),
    ('alex@edu.ru', 'Что такое деструктуризация?', 'destructuring', 2, 5),
    ('maria@edu.ru', 'Используй Promise для асинхронности', 'new Promise()', 2, 15),
    ('sergey@edu.ru', 'Создай HTTP сервер на Node.js', 'http.createServer()', 2, 20),
    ('olga2@edu.ru', 'Создай роут в Express', 'app.get()', 2, 15),
    ('nikolay@edu.ru', 'Создай иконку в Figma', 'icon', 2, 10),
    ('viktoria@edu.ru', 'Создай стиль в Figma', 'style', 2, 12),
    ('pavel@edu.ru', 'Загрузи данные через Pandas', 'pd.read_csv()', 2, 15),
    ('yulia@edu.ru', 'Создай график через matplotlib', 'plt.plot()', 2, 18),
    ('andrey@edu.ru', 'Создай сводную таблицу', 'pivot', 2, 12),
    ('ekaterina@edu.ru', 'Напиши статью для блога', 'article', 2, 15),
    ('maxim@edu.ru', 'Переведи "computer" на немецкий', 'Computer', 2, 5),
    ('dmitry@edu.ru', 'Создай Deployment в Kubernetes', 'kubectl apply', 2, 25)
) AS v(student_email, task_description, answer, status_homework_id, points)
JOIN student s ON s.email = v.student_email
JOIN task t ON t.description = v.task_description
JOIN student_answer sa ON sa.student_id = s.student_id AND sa.task_id = t.task_id AND sa.answer = v.answer
ON CONFLICT DO NOTHING;

-- Transactions - 16 transactions
INSERT INTO "transaction" (student_id, status_id, total_price)
SELECT 
    s.student_id,
    v.status_id,
    v.total_price
FROM (VALUES
    ('anna@edu.ru', 2, 9000),
    ('dmitry@edu.ru', 2, 12000),
    ('elena@edu.ru', 2, 15000),
    ('igor@edu.ru', 1, 10000),
    ('tanya@edu.ru', 2, 8000),
    ('alex@edu.ru', 1, 18000),
    ('maria@edu.ru', 2, 16000),
    ('sergey@edu.ru', 1, 11000),
    ('olga2@edu.ru', 2, 17000),
    ('nikolay@edu.ru', 1, 19000),
    ('viktoria@edu.ru', 2, 14000),
    ('pavel@edu.ru', 2, 22000),
    ('yulia@edu.ru', 1, 9500),
    ('andrey@edu.ru', 2, 10500),
    ('ekaterina@edu.ru', 1, 8500),
    ('maxim@edu.ru', 2, 20000)
) AS v(student_email, status_id, total_price)
JOIN student s ON s.email = v.student_email
ON CONFLICT DO NOTHING;

-- Transactions-Courses relations
INSERT INTO transactions_courses (transaction_id, course_id)
SELECT 
    t.transaction_id,
    c.course_id
FROM (VALUES
    ('anna@edu.ru', 'Основы HTML и CSS'),
    ('dmitry@edu.ru', 'Python для начинающих'),
    ('elena@edu.ru', 'UX-дизайн интерфейсов'),
    ('igor@edu.ru', 'Аналитика данных в SQL'),
    ('tanya@edu.ru', 'Английский для IT'),
    ('alex@edu.ru', 'React и современный фронтенд'),
    ('maria@edu.ru', 'Go для бэкенда'),
    ('sergey@edu.ru', 'SMM стратегии'),
    ('olga2@edu.ru', 'JavaScript продвинутый'),
    ('nikolay@edu.ru', 'Node.js разработка'),
    ('viktoria@edu.ru', 'Figma для дизайнеров'),
    ('pavel@edu.ru', 'Data Science с Python'),
    ('yulia@edu.ru', 'Excel для аналитиков'),
    ('andrey@edu.ru', 'Контент-маркетинг'),
    ('ekaterina@edu.ru', 'Немецкий для IT'),
    ('maxim@edu.ru', 'Docker и контейнеризация')
) AS v(student_email, course_name)
JOIN student s ON s.email = v.student_email
JOIN "transaction" t ON t.student_id = s.student_id
JOIN course c ON c.name = v.course_name
ON CONFLICT DO NOTHING;

-- Display summary
SELECT 'Data seeding completed!' AS status;
SELECT 
    (SELECT COUNT(*) FROM student) AS students,
    (SELECT COUNT(*) FROM teacher) AS teachers,
    (SELECT COUNT(*) FROM course) AS courses,
    (SELECT COUNT(*) FROM lesson) AS lessons,
    (SELECT COUNT(*) FROM homework) AS homeworks,
    (SELECT COUNT(*) FROM task) AS tasks,
    (SELECT COUNT(*) FROM "transaction") AS transactions;

