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

type CourseLessonsHandler struct {
	courseLessonsService service.CourseLessonsServiceInterface
}

func NewCourseLessonsHandler(courseLessonsService service.CourseLessonsServiceInterface) *CourseLessonsHandler {
	return &CourseLessonsHandler{
		courseLessonsService: courseLessonsService,
	}
}

// Create
// @Summary Create course-lesson relation
// @Description Create a new relation between course and lesson
// @Tags Course Lessons
// @Accept json
// @Produce json
// @Param input body dto.CreateCourseLessonsDTO true "Relation data"
// @Success 201 {object} dto.CourseLessonsDTO
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /course-lessons [post]
// @Security BearerAuth
func (h *CourseLessonsHandler) Create(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	body := new(dto.CreateCourseLessonsDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.courseLessonsService.Create(ctx.Context(), body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// GetByID
// @Summary Get course-lesson relation by lesson ID
// @Description Get relation by lesson ID
// @Tags Course Lessons
// @Accept json
// @Produce json
// @Param id path int true "Lesson ID"
// @Success 200 {object} dto.CourseLessonsDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /course-lessons/{id} [get]
// @Security BearerAuth
func (h *CourseLessonsHandler) GetByID(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	relation, err := h.courseLessonsService.GetByID(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(relation)
}

// GetAll
// @Summary Get all course-lesson relations
// @Description Get all relations between courses and lessons
// @Tags Course Lessons
// @Accept json
// @Produce json
// @Success 200 {array} dto.CourseLessonsDTO
// @Failure 500 {object} map[string]string
// @Router /course-lessons [get]
// @Security BearerAuth
func (h *CourseLessonsHandler) GetAll(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	relations, err := h.courseLessonsService.GetAll(ctx.Context())
	if err != nil {
		localLogger.Error(ctx.Context(), "get all relations exception", zap.Error(err))
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(relations)
}

// Delete
// @Summary Delete course-lesson relation
// @Description Delete relation by lesson ID
// @Tags Course Lessons
// @Accept json
// @Produce json
// @Param id path int true "Lesson ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /course-lessons/{id} [delete]
// @Security BearerAuth
func (h *CourseLessonsHandler) Delete(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	err = h.courseLessonsService.Delete(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

