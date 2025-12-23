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

type LevelHandler struct {
	levelService service.LevelServiceInterface
}

func NewLevelHandler(levelService service.LevelServiceInterface) *LevelHandler {
	return &LevelHandler{
		levelService: levelService,
	}
}

// Create
// @Summary Create level
// @Description Create a new level
// @Tags Levels
// @Accept json
// @Produce json
// @Param input body dto.CreateLevelDTO true "Level data"
// @Success 201 {object} dto.LevelDTO
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /levels [post]
// @Security BearerAuth
func (h *LevelHandler) Create(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	body := new(dto.CreateLevelDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.levelService.Create(ctx.Context(), body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// GetByID
// @Summary Get level by ID
// @Description Get level by ID
// @Tags Levels
// @Accept json
// @Produce json
// @Param id path int true "Level ID"
// @Success 200 {object} dto.LevelDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /levels/{id} [get]
// @Security BearerAuth
func (h *LevelHandler) GetByID(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	level, err := h.levelService.GetByID(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(level)
}

// GetAll
// @Summary Get all levels
// @Description Get all levels
// @Tags Levels
// @Accept json
// @Produce json
// @Success 200 {array} dto.LevelDTO
// @Failure 500 {object} map[string]string
// @Router /levels [get]
// @Security BearerAuth
func (h *LevelHandler) GetAll(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	levels, err := h.levelService.GetAll(ctx.Context())
	if err != nil {
		localLogger.Error(ctx.Context(), "get all status exception", zap.Error(err))
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(levels)
}

// Update
// @Summary Update level
// @Description Update level by ID
// @Tags Levels
// @Accept json
// @Produce json
// @Param id path int true "Level ID"
// @Param input body dto.UpdateLevelDTO true "Level data"
// @Success 200 {object} dto.LevelDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /levels/{id} [put]
// @Security BearerAuth
func (h *LevelHandler) Update(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	body := new(dto.UpdateLevelDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.levelService.Update(ctx.Context(), id, body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// Delete
// @Summary Delete level
// @Description Delete level by ID
// @Tags Levels
// @Accept json
// @Produce json
// @Param id path int true "Level ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /levels/{id} [delete]
// @Security BearerAuth
func (h *LevelHandler) Delete(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	err = h.levelService.Delete(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
