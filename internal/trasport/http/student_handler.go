package http

import (
	"strconv"

	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/service"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/exception"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/logger"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type StudentHandler struct {
	studentService service.StudentServiceInterface
}

func NewStudentHandler(studentService service.StudentServiceInterface) *StudentHandler {
	return &StudentHandler{
		studentService: studentService,
	}
}

// @Summary Create student
// @Description Create a new student
// @Tags Students
// @Accept json
// @Produce json
// @Param input body dto.CreateStudentDTO true "Student data"
// @Success 201 {object} dto.StudentDTO
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /students [post]
func (h *StudentHandler) Create(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	body := new(dto.CreateStudentDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.studentService.Create(ctx.Context(), body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// @Summary Get student by ID
// @Description Get student by ID
// @Tags Students
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Success 200 {object} dto.StudentDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Security BearerAuth
// @Router /students/{id} [get]
// @Security BearerAuth
func (h *StudentHandler) GetByID(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	student, err := h.studentService.GetByID(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(student)
}

// @Summary Get all students
// @Description Get all students
// @Tags Students
// @Accept json
// @Produce json
// @Success 200 {array} dto.StudentDTO
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /students [get]
// @Security BearerAuth
func (h *StudentHandler) GetAll(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	students, err := h.studentService.GetAll(ctx.Context())
	if err != nil {
		localLogger.Error(ctx.Context(), "get all status exception", zap.Error(err))
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(students)
}

// @Summary Update student
// @Description Update student by ID
// @Tags Students
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Param input body dto.UpdateStudentDTO true "Student data"
// @Success 200 {object} dto.StudentDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /students/{id} [put]
// @Security BearerAuth
func (h *StudentHandler) Update(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	body := new(dto.UpdateStudentDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.studentService.Update(ctx.Context(), id, body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// @Summary Delete student
// @Description Delete student by ID
// @Tags Students
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /students/{id} [delete]
// @Security BearerAuth
func (h *StudentHandler) Delete(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	err = h.studentService.Delete(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

