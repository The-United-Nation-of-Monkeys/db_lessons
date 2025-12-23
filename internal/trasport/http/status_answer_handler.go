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

type StatusAnswerHandler struct {
	statusAnswerService service.StatusAnswerServiceInterface
}

func NewStatusAnswerHandler(statusAnswerService service.StatusAnswerServiceInterface) *StatusAnswerHandler {
	return &StatusAnswerHandler{
		statusAnswerService: statusAnswerService,
	}
}

// @Summary Create statusanswer
// @Description Create a new statusanswer
// @Tags StatusAnswers
// @Accept json
// @Produce json
// @Param input body dto.CreateStatusAnswerDTO true "StatusAnswer data"
// @Success 201 {object} dto.StatusAnswerDTO
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /status-answers [post]
// @Security BearerAuth
func (h *StatusAnswerHandler) Create(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	body := new(dto.CreateStatusAnswerDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.statusAnswerService.Create(ctx.Context(), body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// @Summary Get statusanswer by ID
// @Description Get statusanswer by ID
// @Tags StatusAnswers
// @Accept json
// @Produce json
// @Param id path int true "StatusAnswer ID"
// @Success 200 {object} dto.StatusAnswerDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /status-answers/{id} [get]
// @Security BearerAuth
func (h *StatusAnswerHandler) GetByID(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	statusanswer, err := h.statusAnswerService.GetByID(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(statusanswer)
}

// @Summary Get all statusanswers
// @Description Get all statusanswers
// @Tags StatusAnswers
// @Accept json
// @Produce json
// @Success 200 {array} dto.StatusAnswerDTO
// @Failure 500 {object} map[string]string
// @Router /status-answers [get]
// @Security BearerAuth
func (h *StatusAnswerHandler) GetAll(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	statusanswers, err := h.statusAnswerService.GetAll(ctx.Context())
	if err != nil {
		localLogger.Error(ctx.Context(), "get all status exception", zap.Error(err))
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(statusanswers)
}

// @Summary Update statusanswer
// @Description Update statusanswer by ID
// @Tags StatusAnswers
// @Accept json
// @Produce json
// @Param id path int true "StatusAnswer ID"
// @Param input body dto.UpdateStatusAnswerDTO true "StatusAnswer data"
// @Success 200 {object} dto.StatusAnswerDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /status-answers/{id} [put]
// @Security BearerAuth
func (h *StatusAnswerHandler) Update(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	body := new(dto.UpdateStatusAnswerDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.statusAnswerService.Update(ctx.Context(), id, body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// @Summary Delete statusanswer
// @Description Delete statusanswer by ID
// @Tags StatusAnswers
// @Accept json
// @Produce json
// @Param id path int true "StatusAnswer ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /status-answers/{id} [delete]
// @Security BearerAuth
func (h *StatusAnswerHandler) Delete(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	err = h.statusAnswerService.Delete(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
