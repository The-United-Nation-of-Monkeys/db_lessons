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

type LessonHandler struct {
	lessonService service.LessonServiceInterface
}

func NewLessonHandler(lessonService service.LessonServiceInterface) *LessonHandler {
	return &LessonHandler{
		lessonService: lessonService,
	}
}

// @Summary Create lesson
// @Description Create a new lesson
// @Tags Lessons
// @Accept json
// @Produce json
// @Param input body dto.CreateLessonDTO true "Lesson data"
// @Success 201 {object} dto.LessonDTO
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /lessons [post]
// @Security BearerAuth
func (h *LessonHandler) Create(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	body := new(dto.CreateLessonDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.lessonService.Create(ctx.Context(), body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// @Summary Get lesson by ID
// @Description Get lesson by ID
// @Tags Lessons
// @Accept json
// @Produce json
// @Param id path int true "Lesson ID"
// @Success 200 {object} dto.LessonDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /lessons/{id} [get]
// @Security BearerAuth
func (h *LessonHandler) GetByID(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	lesson, err := h.lessonService.GetByID(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(lesson)
}

// @Summary Get all lessons
// @Description Get all lessons
// @Tags Lessons
// @Accept json
// @Produce json
// @Success 200 {array} dto.LessonDTO
// @Failure 500 {object} map[string]string
// @Router /lessons [get]
// @Security BearerAuth
func (h *LessonHandler) GetAll(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	lessons, err := h.lessonService.GetAll(ctx.Context())
	if err != nil {
		localLogger.Error(ctx.Context(), "get all status exception", zap.Error(err))
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(lessons)
}

// @Summary Update lesson
// @Description Update lesson by ID
// @Tags Lessons
// @Accept json
// @Produce json
// @Param id path int true "Lesson ID"
// @Param input body dto.UpdateLessonDTO true "Lesson data"
// @Success 200 {object} dto.LessonDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /lessons/{id} [put]
// @Security BearerAuth
func (h *LessonHandler) Update(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	body := new(dto.UpdateLessonDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.lessonService.Update(ctx.Context(), id, body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// @Summary Delete lesson
// @Description Delete lesson by ID
// @Tags Lessons
// @Accept json
// @Produce json
// @Param id path int true "Lesson ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /lessons/{id} [delete]
// @Security BearerAuth
func (h *LessonHandler) Delete(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	err = h.lessonService.Delete(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
