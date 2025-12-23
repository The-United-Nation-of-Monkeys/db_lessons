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

type TeachersCoursesHandler struct {
	teachersCoursesService service.TeachersCoursesServiceInterface
}

func NewTeachersCoursesHandler(teachersCoursesService service.TeachersCoursesServiceInterface) *TeachersCoursesHandler {
	return &TeachersCoursesHandler{
		teachersCoursesService: teachersCoursesService,
	}
}

// Create
// @Summary Create teacher-course relation
// @Description Create a new relation between teacher and course
// @Tags Teacher Courses
// @Accept json
// @Produce json
// @Param input body dto.CreateTeachersCoursesDTO true "Relation data"
// @Success 201 {object} dto.TeachersCoursesDTO
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /teachers-courses [post]
// @Security BearerAuth
func (h *TeachersCoursesHandler) Create(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	body := new(dto.CreateTeachersCoursesDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.teachersCoursesService.Create(ctx.Context(), body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// GetByID
// @Summary Get teacher-course relation by IDs
// @Description Get relation by course ID and teacher ID
// @Tags Teacher Courses
// @Accept json
// @Produce json
// @Param course_id path int true "Course ID"
// @Param teacher_id path int true "Teacher ID"
// @Success 200 {object} dto.TeachersCoursesDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /teachers-courses/{course_id}/{teacher_id} [get]
// @Security BearerAuth
func (h *TeachersCoursesHandler) GetByID(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	courseIDStr := ctx.Params("course_id")
	courseID, err := strconv.Atoi(courseIDStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param course_id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid course_id")
	}

	teacherIDStr := ctx.Params("teacher_id")
	teacherID, err := strconv.Atoi(teacherIDStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param teacher_id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid teacher_id")
	}

	relation, err := h.teachersCoursesService.GetByID(ctx.Context(), courseID, teacherID)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(relation)
}

// GetAll
// @Summary Get all teacher-course relations
// @Description Get all relations between teachers and courses
// @Tags Teacher Courses
// @Accept json
// @Produce json
// @Success 200 {array} dto.TeachersCoursesDTO
// @Failure 500 {object} map[string]string
// @Router /teachers-courses [get]
// @Security BearerAuth
func (h *TeachersCoursesHandler) GetAll(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	relations, err := h.teachersCoursesService.GetAll(ctx.Context())
	if err != nil {
		localLogger.Error(ctx.Context(), "get all relations exception", zap.Error(err))
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(relations)
}

// Delete
// @Summary Delete teacher-course relation
// @Description Delete relation by course ID and teacher ID
// @Tags Teacher Courses
// @Accept json
// @Produce json
// @Param course_id path int true "Course ID"
// @Param teacher_id path int true "Teacher ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /teachers-courses/{course_id}/{teacher_id} [delete]
// @Security BearerAuth
func (h *TeachersCoursesHandler) Delete(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	courseIDStr := ctx.Params("course_id")
	courseID, err := strconv.Atoi(courseIDStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param course_id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid course_id")
	}

	teacherIDStr := ctx.Params("teacher_id")
	teacherID, err := strconv.Atoi(teacherIDStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param teacher_id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid teacher_id")
	}

	err = h.teachersCoursesService.Delete(ctx.Context(), courseID, teacherID)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

