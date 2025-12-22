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

type TransactionsCoursesHandler struct {
	transactionsCoursesService service.TransactionsCoursesServiceInterface
}

func NewTransactionsCoursesHandler(transactionsCoursesService service.TransactionsCoursesServiceInterface) *TransactionsCoursesHandler {
	return &TransactionsCoursesHandler{
		transactionsCoursesService: transactionsCoursesService,
	}
}

// Create
// @Summary Create transaction-course relation
// @Description Create a new relation between transaction and course
// @Tags Transaction Courses
// @Accept json
// @Produce json
// @Param input body dto.CreateTransactionsCoursesDTO true "Relation data"
// @Success 201 {object} dto.TransactionsCoursesDTO
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions-courses [post]
func (h *TransactionsCoursesHandler) Create(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	body := new(dto.CreateTransactionsCoursesDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.transactionsCoursesService.Create(ctx.Context(), body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// GetByID
// @Summary Get transaction-course relation by IDs
// @Description Get relation by transaction ID and course ID
// @Tags Transaction Courses
// @Accept json
// @Produce json
// @Param transaction_id path int true "Transaction ID"
// @Param course_id path int true "Course ID"
// @Success 200 {object} dto.TransactionsCoursesDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions-courses/{transaction_id}/{course_id} [get]
func (h *TransactionsCoursesHandler) GetByID(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	transactionIDStr := ctx.Params("transaction_id")
	transactionID, err := strconv.Atoi(transactionIDStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param transaction_id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid transaction_id")
	}

	courseIDStr := ctx.Params("course_id")
	courseID, err := strconv.Atoi(courseIDStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param course_id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid course_id")
	}

	relation, err := h.transactionsCoursesService.GetByID(ctx.Context(), transactionID, courseID)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(relation)
}

// GetAll
// @Summary Get all transaction-course relations
// @Description Get all relations between transactions and courses
// @Tags Transaction Courses
// @Accept json
// @Produce json
// @Success 200 {array} dto.TransactionsCoursesDTO
// @Failure 500 {object} map[string]string
// @Router /transactions-courses [get]
func (h *TransactionsCoursesHandler) GetAll(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	relations, err := h.transactionsCoursesService.GetAll(ctx.Context())
	if err != nil {
		localLogger.Error(ctx.Context(), "get all relations exception", zap.Error(err))
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(relations)
}

// Delete
// @Summary Delete transaction-course relation
// @Description Delete relation by transaction ID and course ID
// @Tags Transaction Courses
// @Accept json
// @Produce json
// @Param transaction_id path int true "Transaction ID"
// @Param course_id path int true "Course ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions-courses/{transaction_id}/{course_id} [delete]
func (h *TransactionsCoursesHandler) Delete(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	transactionIDStr := ctx.Params("transaction_id")
	transactionID, err := strconv.Atoi(transactionIDStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param transaction_id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid transaction_id")
	}

	courseIDStr := ctx.Params("course_id")
	courseID, err := strconv.Atoi(courseIDStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param course_id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid course_id")
	}

	err = h.transactionsCoursesService.Delete(ctx.Context(), transactionID, courseID)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
