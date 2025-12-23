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

type CourseHandler struct {
	courseService service.CourseServiceInterface
}

func NewCourseHandler(courseService service.CourseServiceInterface) *CourseHandler {
	return &CourseHandler{
		courseService: courseService,
	}
}

// @Summary Create course
// @Description Create a new course
// @Tags Courses
// @Accept json
// @Produce json
// @Param input body dto.CreateCourseDTO true "Course data"
// @Success 201 {object} dto.CourseDTO
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /courses [post]
// @Security BearerAuth
func (h *CourseHandler) Create(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	body := new(dto.CreateCourseDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.courseService.Create(ctx.Context(), body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// @Summary Get course by ID
// @Description Get course by ID
// @Tags Courses
// @Accept json
// @Produce json
// @Param id path int true "Course ID"
// @Success 200 {object} dto.CourseDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /courses/{id} [get]
// @Security BearerAuth
func (h *CourseHandler) GetByID(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	course, err := h.courseService.GetByID(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(course)
}

// @Summary Get all courses
// @Description Get all courses
// @Tags Courses
// @Accept json
// @Produce json
// @Success 200 {array} dto.CourseDTO
// @Failure 500 {object} map[string]string
// @Router /courses [get]
// @Security BearerAuth
func (h *CourseHandler) GetAll(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	courses, err := h.courseService.GetAll(ctx.Context())
	if err != nil {
		localLogger.Error(ctx.Context(), "get all status exception", zap.Error(err))
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(courses)
}

// @Summary Update course
// @Description Update course by ID
// @Tags Courses
// @Accept json
// @Produce json
// @Param id path int true "Course ID"
// @Param input body dto.UpdateCourseDTO true "Course data"
// @Success 200 {object} dto.CourseDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /courses/{id} [put]
// @Security BearerAuth
func (h *CourseHandler) Update(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	body := new(dto.UpdateCourseDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.courseService.Update(ctx.Context(), id, body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// @Summary Delete course
// @Description Delete course by ID
// @Tags Courses
// @Accept json
// @Produce json
// @Param id path int true "Course ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /courses/{id} [delete]
// @Security BearerAuth
func (h *CourseHandler) Delete(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	err = h.courseService.Delete(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
