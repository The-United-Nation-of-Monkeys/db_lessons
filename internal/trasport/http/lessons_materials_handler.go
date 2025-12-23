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

type LessonsMaterialsHandler struct {
	lessonsMaterialsService service.LessonsMaterialsServiceInterface
}

func NewLessonsMaterialsHandler(lessonsMaterialsService service.LessonsMaterialsServiceInterface) *LessonsMaterialsHandler {
	return &LessonsMaterialsHandler{
		lessonsMaterialsService: lessonsMaterialsService,
	}
}

// Create
// @Summary Create lesson-material relation
// @Description Create a new relation between lesson and material
// @Tags Lesson Materials
// @Accept json
// @Produce json
// @Param input body dto.CreateLessonsMaterialsDTO true "Relation data"
// @Success 201 {object} dto.LessonsMaterialsDTO
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /lessons-materials [post]
// @Security BearerAuth
func (h *LessonsMaterialsHandler) Create(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	body := new(dto.CreateLessonsMaterialsDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.lessonsMaterialsService.Create(ctx.Context(), body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// GetByID
// @Summary Get lesson-material relation by IDs
// @Description Get relation by lesson ID and material ID
// @Tags Lesson Materials
// @Accept json
// @Produce json
// @Param lesson_id path int true "Lesson ID"
// @Param material_id path int true "Material ID"
// @Success 200 {object} dto.LessonsMaterialsDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /lessons-materials/{lesson_id}/{material_id} [get]
// @Security BearerAuth
func (h *LessonsMaterialsHandler) GetByID(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	lessonIDStr := ctx.Params("lesson_id")
	lessonID, err := strconv.Atoi(lessonIDStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param lesson_id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid lesson_id")
	}

	materialIDStr := ctx.Params("material_id")
	materialID, err := strconv.Atoi(materialIDStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param material_id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid material_id")
	}

	relation, err := h.lessonsMaterialsService.GetByID(ctx.Context(), lessonID, materialID)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(relation)
}

// GetAll
// @Summary Get all lesson-material relations
// @Description Get all relations between lessons and materials
// @Tags Lesson Materials
// @Accept json
// @Produce json
// @Success 200 {array} dto.LessonsMaterialsDTO
// @Failure 500 {object} map[string]string
// @Router /lessons-materials [get]
// @Security BearerAuth
func (h *LessonsMaterialsHandler) GetAll(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	relations, err := h.lessonsMaterialsService.GetAll(ctx.Context())
	if err != nil {
		localLogger.Error(ctx.Context(), "get all relations exception", zap.Error(err))
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(relations)
}

// Delete
// @Summary Delete lesson-material relation
// @Description Delete relation by lesson ID and material ID
// @Tags Lesson Materials
// @Accept json
// @Produce json
// @Param lesson_id path int true "Lesson ID"
// @Param material_id path int true "Material ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /lessons-materials/{lesson_id}/{material_id} [delete]
// @Security BearerAuth
func (h *LessonsMaterialsHandler) Delete(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	lessonIDStr := ctx.Params("lesson_id")
	lessonID, err := strconv.Atoi(lessonIDStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param lesson_id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid lesson_id")
	}

	materialIDStr := ctx.Params("material_id")
	materialID, err := strconv.Atoi(materialIDStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param material_id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid material_id")
	}

	err = h.lessonsMaterialsService.Delete(ctx.Context(), lessonID, materialID)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

