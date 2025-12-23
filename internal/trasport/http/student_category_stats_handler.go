package http

import (
	"strconv"

	_ "github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto" // for swagger
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/service"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/exception"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/logger"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type StudentCategoryStatsHandler struct {
	studentCategoryStatsService service.StudentCategoryStatsServiceInterface
}

func NewStudentCategoryStatsHandler(studentCategoryStatsService service.StudentCategoryStatsServiceInterface) *StudentCategoryStatsHandler {
	return &StudentCategoryStatsHandler{
		studentCategoryStatsService: studentCategoryStatsService,
	}
}

// GetAll
// @Summary Get all student category statistics
// @Description Get statistics for all students by categories
// @Tags Reports
// @Accept json
// @Produce json
// @Success 200 {array} dto.StudentCategoryStatsDTO
// @Failure 500 {object} map[string]string
// @Router /reports/student-category-stats [get]
// @Security BearerAuth
func (h *StudentCategoryStatsHandler) GetAll(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	stats, err := h.studentCategoryStatsService.GetAll(ctx.Context())
	if err != nil {
		localLogger.Error(ctx.Context(), "get all stats error", zap.Error(err))
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(stats)
}

// GetByStudentID
// @Summary Get student category statistics by student ID
// @Description Get statistics for a specific student by categories
// @Tags Reports
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Success 200 {array} dto.StudentCategoryStatsDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /reports/student-category-stats/students/{id} [get]
// @Security BearerAuth
func (h *StudentCategoryStatsHandler) GetByStudentID(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	stats, err := h.studentCategoryStatsService.GetByStudentID(ctx.Context(), id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(stats)
}

// GetByCategoryID
// @Summary Get student category statistics by category ID
// @Description Get statistics for all students in a specific category
// @Tags Reports
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {array} dto.StudentCategoryStatsDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /reports/student-category-stats/categories/{id} [get]
// @Security BearerAuth
func (h *StudentCategoryStatsHandler) GetByCategoryID(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	stats, err := h.studentCategoryStatsService.GetByCategoryID(ctx.Context(), id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(stats)
}
