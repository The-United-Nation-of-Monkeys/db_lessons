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

type CurrencyHandler struct {
	currencyService service.CurrencyServiceInterface
}

func NewCurrencyHandler(currencyService service.CurrencyServiceInterface) *CurrencyHandler {
	return &CurrencyHandler{
		currencyService: currencyService,
	}
}

// @Summary Create currency
// @Description Create a new currency
// @Tags Currencies
// @Accept json
// @Produce json
// @Param input body dto.CreateCurrencyDTO true "Currency data"
// @Success 201 {object} dto.CurrencyDTO
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /currencies [post]
// @Security BearerAuth
func (h *CurrencyHandler) Create(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	body := new(dto.CreateCurrencyDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.currencyService.Create(ctx.Context(), body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// @Summary Get currency by ID
// @Description Get currency by ID
// @Tags Currencies
// @Accept json
// @Produce json
// @Param id path int true "Currency ID"
// @Success 200 {object} dto.CurrencyDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /currencies/{id} [get]
// @Security BearerAuth
func (h *CurrencyHandler) GetByID(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	currency, err := h.currencyService.GetByID(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(currency)
}

// @Summary Get all currencys
// @Description Get all currencys
// @Tags Currencies
// @Accept json
// @Produce json
// @Success 200 {array} dto.CurrencyDTO
// @Failure 500 {object} map[string]string
// @Router /currencies [get]
// @Security BearerAuth
func (h *CurrencyHandler) GetAll(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	currencys, err := h.currencyService.GetAll(ctx.Context())
	if err != nil {
		localLogger.Error(ctx.Context(), "get all status exception", zap.Error(err))
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(currencys)
}

// @Summary Update currency
// @Description Update currency by ID
// @Tags Currencies
// @Accept json
// @Produce json
// @Param id path int true "Currency ID"
// @Param input body dto.UpdateCurrencyDTO true "Currency data"
// @Success 200 {object} dto.CurrencyDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /currencies/{id} [put]
// @Security BearerAuth
func (h *CurrencyHandler) Update(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	body := new(dto.UpdateCurrencyDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.currencyService.Update(ctx.Context(), id, body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// @Summary Delete currency
// @Description Delete currency by ID
// @Tags Currencies
// @Accept json
// @Produce json
// @Param id path int true "Currency ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /currencies/{id} [delete]
// @Security BearerAuth
func (h *CurrencyHandler) Delete(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	err = h.currencyService.Delete(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
