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

type StatusHomeworkHandler struct {
	statusHomeworkService service.StatusHomeworkServiceInterface
}

func NewStatusHomeworkHandler(statusHomeworkService service.StatusHomeworkServiceInterface) *StatusHomeworkHandler {
	return &StatusHomeworkHandler{
		statusHomeworkService: statusHomeworkService,
	}
}

// Create
// @Summary Create statushomework
// @Description Create a new statushomework
// @Tags StatusHomeworks
// @Accept json
// @Produce json
// @Param input body dto.CreateStatusHomeworkDTO true "StatusHomework data"
// @Success 201 {object} dto.StatusHomeworkDTO
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /status-homeworks [post]
// @Security BearerAuth
func (h *StatusHomeworkHandler) Create(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	body := new(dto.CreateStatusHomeworkDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.statusHomeworkService.Create(ctx.Context(), body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// GetByID
// @Summary Get statushomework by ID
// @Description Get statushomework by ID
// @Tags StatusHomeworks
// @Accept json
// @Produce json
// @Param id path int true "StatusHomework ID"
// @Success 200 {object} dto.StatusHomeworkDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /status-homeworks/{id} [get]
// @Security BearerAuth
func (h *StatusHomeworkHandler) GetByID(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	statushomework, err := h.statusHomeworkService.GetByID(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(statushomework)
}

// GetAll
// @Summary Get all statushomeworks
// @Description Get all statushomeworks
// @Tags StatusHomeworks
// @Accept json
// @Produce json
// @Success 200 {array} dto.StatusHomeworkDTO
// @Failure 500 {object} map[string]string
// @Router /status-homeworks [get]
// @Security BearerAuth
func (h *StatusHomeworkHandler) GetAll(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	statushomeworks, err := h.statusHomeworkService.GetAll(ctx.Context())
	if err != nil {
		localLogger.Error(ctx.Context(), "get all status homeworks exception", zap.Error(err))
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(statushomeworks)
}

// Update
// @Summary Update statushomework
// @Description Update statushomework by ID
// @Tags StatusHomeworks
// @Accept json
// @Produce json
// @Param id path int true "StatusHomework ID"
// @Param input body dto.UpdateStatusHomeworkDTO true "StatusHomework data"
// @Success 200 {object} dto.StatusHomeworkDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /status-homeworks/{id} [put]
// @Security BearerAuth
func (h *StatusHomeworkHandler) Update(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	body := new(dto.UpdateStatusHomeworkDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.statusHomeworkService.Update(ctx.Context(), id, body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// Delete
// @Summary Delete statushomework
// @Description Delete statushomework by ID
// @Tags StatusHomeworks
// @Accept json
// @Produce json
// @Param id path int true "StatusHomework ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /status-homeworks/{id} [delete]
// @Security BearerAuth
func (h *StatusHomeworkHandler) Delete(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	err = h.statusHomeworkService.Delete(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
