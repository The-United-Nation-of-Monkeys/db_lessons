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

type LessonHomeworksHandler struct {
	lessonHomeworksService service.LessonHomeworksServiceInterface
}

func NewLessonHomeworksHandler(lessonHomeworksService service.LessonHomeworksServiceInterface) *LessonHomeworksHandler {
	return &LessonHomeworksHandler{
		lessonHomeworksService: lessonHomeworksService,
	}
}

// Create
// @Summary Create lesson-homework relation
// @Description Create a new relation between lesson and homework
// @Tags Lesson Homeworks
// @Accept json
// @Produce json
// @Param input body dto.CreateLessonHomeworksDTO true "Relation data"
// @Success 201 {object} dto.LessonHomeworksDTO
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /lesson-homeworks [post]
func (h *LessonHomeworksHandler) Create(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	body := new(dto.CreateLessonHomeworksDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.lessonHomeworksService.Create(ctx.Context(), body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// GetByID
// @Summary Get lesson-homework relation by IDs
// @Description Get relation by lesson ID and homework ID
// @Tags Lesson Homeworks
// @Accept json
// @Produce json
// @Param lesson_id path int true "Lesson ID"
// @Param homework_id path int true "Homework ID"
// @Success 200 {object} dto.LessonHomeworksDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /lesson-homeworks/{lesson_id}/{homework_id} [get]
func (h *LessonHomeworksHandler) GetByID(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	lessonIDStr := ctx.Params("lesson_id")
	lessonID, err := strconv.Atoi(lessonIDStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param lesson_id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid lesson_id")
	}

	homeworkIDStr := ctx.Params("homework_id")
	homeworkID, err := strconv.Atoi(homeworkIDStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param homework_id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid homework_id")
	}

	relation, err := h.lessonHomeworksService.GetByID(ctx.Context(), lessonID, homeworkID)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(relation)
}

// GetAll
// @Summary Get all lesson-homework relations
// @Description Get all relations between lessons and homeworks
// @Tags Lesson Homeworks
// @Accept json
// @Produce json
// @Success 200 {array} dto.LessonHomeworksDTO
// @Failure 500 {object} map[string]string
// @Router /lesson-homeworks [get]
func (h *LessonHomeworksHandler) GetAll(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	relations, err := h.lessonHomeworksService.GetAll(ctx.Context())
	if err != nil {
		localLogger.Error(ctx.Context(), "get all relations exception", zap.Error(err))
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(relations)
}

// Delete
// @Summary Delete lesson-homework relation
// @Description Delete relation by lesson ID and homework ID
// @Tags Lesson Homeworks
// @Accept json
// @Produce json
// @Param lesson_id path int true "Lesson ID"
// @Param homework_id path int true "Homework ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /lesson-homeworks/{lesson_id}/{homework_id} [delete]
func (h *LessonHomeworksHandler) Delete(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	lessonIDStr := ctx.Params("lesson_id")
	lessonID, err := strconv.Atoi(lessonIDStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param lesson_id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid lesson_id")
	}

	homeworkIDStr := ctx.Params("homework_id")
	homeworkID, err := strconv.Atoi(homeworkIDStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param homework_id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid homework_id")
	}

	err = h.lessonHomeworksService.Delete(ctx.Context(), lessonID, homeworkID)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

