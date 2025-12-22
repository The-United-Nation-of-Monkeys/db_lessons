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

type StatusTransactionHandler struct {
	statusTransactionService service.StatusTransactionServiceInterface
}

func NewStatusTransactionHandler(statusTransactionService service.StatusTransactionServiceInterface) *StatusTransactionHandler {
	return &StatusTransactionHandler{
		statusTransactionService: statusTransactionService,
	}
}

// Create
// @Summary Create statustransaction
// @Description Create a new statustransaction
// @Tags StatusTransactions
// @Accept json
// @Produce json
// @Param input body dto.CreateStatusTransactionDTO true "StatusTransaction data"
// @Success 201 {object} dto.StatusTransactionDTO
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /status-transactions [post]
func (h *StatusTransactionHandler) Create(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	body := new(dto.CreateStatusTransactionDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.statusTransactionService.Create(ctx.Context(), body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// GetByID
// @Summary Get statustransaction by ID
// @Description Get statustransaction by ID
// @Tags StatusTransactions
// @Accept json
// @Produce json
// @Param id path int true "StatusTransaction ID"
// @Success 200 {object} dto.StatusTransactionDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /status-transactions/{id} [get]
func (h *StatusTransactionHandler) GetByID(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	statustransaction, err := h.statusTransactionService.GetByID(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(statustransaction)
}

// GetAll
// @Summary Get all statustransactions
// @Description Get all statustransactions
// @Tags StatusTransactions
// @Accept json
// @Produce json
// @Success 200 {array} dto.StatusTransactionDTO
// @Failure 500 {object} map[string]string
// @Router /status-transactions [get]
func (h *StatusTransactionHandler) GetAll(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	statustransactions, err := h.statusTransactionService.GetAll(ctx.Context())
	if err != nil {
		localLogger.Error(ctx.Context(), "get all status exception", zap.Error(err))
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(statustransactions)
}

// Update
// @Summary Update statustransaction
// @Description Update statustransaction by ID
// @Tags StatusTransactions
// @Accept json
// @Produce json
// @Param id path int true "StatusTransaction ID"
// @Param input body dto.UpdateStatusTransactionDTO true "StatusTransaction data"
// @Success 200 {object} dto.StatusTransactionDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /status-transactions/{id} [put]
func (h *StatusTransactionHandler) Update(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	body := new(dto.UpdateStatusTransactionDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.statusTransactionService.Update(ctx.Context(), id, body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// Delete
// @Summary Delete statustransaction
// @Description Delete statustransaction by ID
// @Tags StatusTransactions
// @Accept json
// @Produce json
// @Param id path int true "StatusTransaction ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /status-transactions/{id} [delete]
func (h *StatusTransactionHandler) Delete(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	err = h.statusTransactionService.Delete(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
