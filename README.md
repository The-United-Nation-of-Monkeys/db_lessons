# Online Classes Database Lab

Проект для работы с базой данных онлайн-курсов. Реализованы CRUD операции для всех таблиц базы данных.

## Структура проекта

- `cmd/app/main.go` - точка входа приложения
- `internal/` - внутренняя логика приложения
  - `config/` - конфигурация
  - `dto/` - Data Transfer Objects
  - `repository/` - репозитории для работы с БД
  - `service/` - бизнес-логика
  - `trasport/http/` - HTTP handlers и роутинг
  - `container/` - dependency injection
- `migrations/` - миграции базы данных
- `pkg/` - общие пакеты (database, logger, exception, validator, redis)

## Установка и запуск

1. Создайте базу данных:
```sql
CREATE DATABASE online_classes;
```

2. Настройте `config.env`:
```
SERVER_VERSION=1
SERVER_PORT_HTTP=8080
LOGGER_MOD=production
POSTGRES_DB=online_classes
POSTGRES_PASSWORD=postgres
POSTGRES_USER=postgres
POSTGRES_PORT=5432
POSTGRES_HOST=localhost
POSTGRES_MIN_CONN=5
POSTGRES_MAX_CONN=15
REDIS_PORT=6379
REDIS_HOST=localhost
```

3. Установите зависимости:
```bash
go mod tidy
```

4. Запустите приложение:

**Локально:**
```bash
go run cmd/app/main.go
```

**Или с помощью Docker Compose:**
```bash
# Соберите и запустите все сервисы (app, postgres, redis)
docker-compose up --build

# Или в фоновом режиме
docker-compose up -d --build

# Остановить все сервисы
docker-compose down

# Остановить и удалить volumes
docker-compose down -v
```

## Swagger документация

После запуска приложения Swagger UI доступен по адресу:
```
http://localhost:8080/api/v1/swagger/index.html
```

Для обновления Swagger документации после изменений в handlers:
```bash
# С помощью Makefile (рекомендуется)
make swagger

# Или напрямую
swag init -g cmd/app/main.go -o docs --parseDependency --parseInternal --exclude test-task-wallet
```

**Примечание:** Swagger аннотации добавлены для всех handlers (CRUD операции для всех сущностей).

## API Endpoints

Все endpoints доступны по префиксу `/api/v1`

### Студенты
- `POST /api/v1/students` - создать студента
- `GET /api/v1/students` - получить всех студентов
- `GET /api/v1/students/:id` - получить студента по ID
- `PUT /api/v1/students/:id` - обновить студента
- `DELETE /api/v1/students/:id` - удалить студента

### Преподаватели
- `POST /api/v1/teachers` - создать преподавателя
- `GET /api/v1/teachers` - получить всех преподавателей
- `GET /api/v1/teachers/:id` - получить преподавателя по ID
- `PUT /api/v1/teachers/:id` - обновить преподавателя
- `DELETE /api/v1/teachers/:id` - удалить преподавателя

Аналогичные endpoints доступны для:
- categories (категории)
- currencies (валюты)
- levels (уровни)
- subcategories (подкатегории)
- materials (материалы)
- homeworks (домашние задания)
- lessons (уроки)
- courses (курсы)
- tasks (задачи)
- status-homeworks (статусы домашних заданий)
- status-answers (статусы ответов)
- status-transactions (статусы транзакций)
- student-answers (ответы студентов)
- homework-results (результаты домашних заданий)
- transactions (транзакции)

## База данных

База данных содержит следующие таблицы:
- student (студенты)
- teacher (преподаватели)
- category (категории)
- currency (валюты)
- course (курсы)
- lesson (уроки)
- homework (домашние задания)
- task (задачи)
- material (материалы)
- level (уровни)
- subcategory (подкатегории)
- status_homework (статусы домашних заданий)
- status_answer (статусы ответов)
- status_transaction (статусы транзакций)
- student_answer (ответы студентов)
- homework_result (результаты домашних заданий)
- transaction (транзакции)
- teachers_courses (связь преподавателей и курсов)
- course_lessons (связь курсов и уроков)
- lessons_materials (связь уроков и материалов)
- lesson_homeworks (связь уроков и домашних заданий)
- homeworks_tasks (связь домашних заданий и задач)
- transactions_courses (связь транзакций и курсов)

Миграции автоматически применяются при запуске приложения.

## Docker

Проект включает Dockerfile и docker-compose.yml для удобного развертывания.

### Docker Compose

Docker Compose включает три сервиса:
- **app** - Go приложение
- **db** - PostgreSQL база данных
- **redis** - Redis кэш

### Использование Docker Compose

1. Убедитесь, что Docker и Docker Compose установлены
2. Запустите все сервисы:
   ```bash
   docker-compose up --build
   ```
3. Приложение будет доступно на `http://localhost:8080`
4. Swagger UI: `http://localhost:8080/api/v1/swagger/index.html`

### Переменные окружения для Docker

При использовании Docker Compose, хосты БД и Redis автоматически настраиваются на имена сервисов (`db` и `redis`). Для локальной разработки используйте `config.env`, для Docker - настройки в `docker-compose.yml`.

