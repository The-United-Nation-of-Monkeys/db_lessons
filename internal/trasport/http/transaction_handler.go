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

type TransactionHandler struct {
	transactionService service.TransactionServiceInterface
}

func NewTransactionHandler(transactionService service.TransactionServiceInterface) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
	}
}

// @Summary Create transaction
// @Description Create a new transaction
// @Tags Transactions
// @Accept json
// @Produce json
// @Param input body dto.CreateTransactionDTO true "Transaction data"
// @Success 201 {object} dto.TransactionDTO
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions [post]
// @Security BearerAuth
func (h *TransactionHandler) Create(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	body := new(dto.CreateTransactionDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.transactionService.Create(ctx.Context(), body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// @Summary Get transaction by ID
// @Description Get transaction by ID
// @Tags Transactions
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} dto.TransactionDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /transactions/{id} [get]
// @Security BearerAuth
func (h *TransactionHandler) GetByID(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	transaction, err := h.transactionService.GetByID(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(transaction)
}

// @Summary Get all transactions
// @Description Get all transactions
// @Tags Transactions
// @Accept json
// @Produce json
// @Success 200 {array} dto.TransactionDTO
// @Failure 500 {object} map[string]string
// @Router /transactions [get]
// @Security BearerAuth
func (h *TransactionHandler) GetAll(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	transactions, err := h.transactionService.GetAll(ctx.Context())
	if err != nil {
		localLogger.Error(ctx.Context(), "get all status exception", zap.Error(err))
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(transactions)
}

// @Summary Update transaction
// @Description Update transaction by ID
// @Tags Transactions
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Param input body dto.UpdateTransactionDTO true "Transaction data"
// @Success 200 {object} dto.TransactionDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions/{id} [put]
// @Security BearerAuth
func (h *TransactionHandler) Update(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	body := new(dto.UpdateTransactionDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.transactionService.Update(ctx.Context(), id, body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// @Summary Delete transaction
// @Description Delete transaction by ID
// @Tags Transactions
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions/{id} [delete]
// @Security BearerAuth
func (h *TransactionHandler) Delete(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	err = h.transactionService.Delete(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

// @Summary Get transaction report by parameters
// @Description Get transaction report filtered by status name, min total, max total
// @Tags Transactions
// @Accept json
// @Produce json
// @Param status_name query string false "Status name"
// @Param min_total query int false "Min total price"
// @Param max_total query int false "Max total price"
// @Success 200 {array} dto.TransactionReportDTO
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions/report [get]
// @Security BearerAuth
func (h *TransactionHandler) GetReportByParams(ctx fiber.Ctx) error {
	params := &dto.TransactionReportRequestDTO{
		StatusName: ctx.Query("status_name"),
	}

	if minTotalStr := ctx.Query("min_total"); minTotalStr != "" {
		if minTotal, err := strconv.Atoi(minTotalStr); err == nil {
			params.MinTotal = minTotal
		}
	}

	if maxTotalStr := ctx.Query("max_total"); maxTotalStr != "" {
		if maxTotal, err := strconv.Atoi(maxTotalStr); err == nil {
			params.MaxTotal = maxTotal
		}
	}

	reports, err := h.transactionService.GetReportByParams(ctx.Context(), params)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(reports)
}

// @Summary Bulk update transaction status
// @Description Bulk update transaction status by old status id, new status id and optional price range
// @Tags Transactions
// @Accept json
// @Produce json
// @Param input body dto.BulkUpdateTransactionStatusDTO true "Bulk update parameters"
// @Success 200 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions/bulk-update-status [post]
// @Security BearerAuth
func (h *TransactionHandler) BulkUpdateTransactionStatus(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	body := new(dto.BulkUpdateTransactionStatusDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	err := h.transactionService.BulkUpdateTransactionStatus(ctx.Context(), body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "bulk update completed")

	return ctx.Status(fiber.StatusOK).JSON(map[string]string{
		"message": "Transaction statuses updated successfully",
	})
}
