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

type TeacherHandler struct {
	teacherService service.TeacherServiceInterface
}

func NewTeacherHandler(teacherService service.TeacherServiceInterface) *TeacherHandler {
	return &TeacherHandler{
		teacherService: teacherService,
	}
}

// Create
// @Summary Create teacher
// @Description Create a new teacher
// @Tags Teachers
// @Accept json
// @Produce json
// @Param input body dto.CreateTeacherDTO true "Teacher data"
// @Success 201 {object} dto.TeacherDTO
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /teachers [post]
func (h *TeacherHandler) Create(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	body := new(dto.CreateTeacherDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.teacherService.Create(ctx.Context(), body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// GetByID
// @Summary Get teacher by ID
// @Description Get teacher by ID
// @Tags Teachers
// @Accept json
// @Produce json
// @Param id path int true "Teacher ID"
// @Success 200 {object} dto.TeacherDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /teachers/{id} [get]
func (h *TeacherHandler) GetByID(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	teacher, err := h.teacherService.GetByID(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(teacher)
}

// GetAll
// @Summary Get all teachers
// @Description Get all teachers
// @Tags Teachers
// @Accept json
// @Produce json
// @Success 200 {array} dto.TeacherDTO
// @Failure 500 {object} map[string]string
// @Router /teachers [get]
func (h *TeacherHandler) GetAll(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	teachers, err := h.teacherService.GetAll(ctx.Context())
	if err != nil {
		localLogger.Error(ctx.Context(), "get all status exception", zap.Error(err))
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(teachers)
}

// Update
// @Summary Update teacher
// @Description Update teacher by ID
// @Tags Teachers
// @Accept json
// @Produce json
// @Param id path int true "Teacher ID"
// @Param input body dto.UpdateTeacherDTO true "Teacher data"
// @Success 200 {object} dto.TeacherDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /teachers/{id} [put]
func (h *TeacherHandler) Update(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	body := new(dto.UpdateTeacherDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.teacherService.Update(ctx.Context(), id, body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// Delete
// @Summary Delete teacher
// @Description Delete teacher by ID
// @Tags Teachers
// @Accept json
// @Produce json
// @Param id path int true "Teacher ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /teachers/{id} [delete]
func (h *TeacherHandler) Delete(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	err = h.teacherService.Delete(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
