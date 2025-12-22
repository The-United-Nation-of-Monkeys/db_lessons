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

type MaterialHandler struct {
	materialService service.MaterialServiceInterface
}

func NewMaterialHandler(materialService service.MaterialServiceInterface) *MaterialHandler {
	return &MaterialHandler{
		materialService: materialService,
	}
}

// Create
// @Summary Create material
// @Description Create a new material
// @Tags Materials
// @Accept json
// @Produce json
// @Param input body dto.CreateMaterialDTO true "Material data"
// @Success 201 {object} dto.MaterialDTO
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /materials [post]
func (h *MaterialHandler) Create(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	body := new(dto.CreateMaterialDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.materialService.Create(ctx.Context(), body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// GetByID
// @Summary Get material by ID
// @Description Get material by ID
// @Tags Materials
// @Accept json
// @Produce json
// @Param id path int true "Material ID"
// @Success 200 {object} dto.MaterialDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /materials/{id} [get]
func (h *MaterialHandler) GetByID(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	material, err := h.materialService.GetByID(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(material)
}

// GetAll
// @Summary Get all materials
// @Description Get all materials
// @Tags Materials
// @Accept json
// @Produce json
// @Success 200 {array} dto.MaterialDTO
// @Failure 500 {object} map[string]string
// @Router /materials [get]
func (h *MaterialHandler) GetAll(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	materials, err := h.materialService.GetAll(ctx.Context())
	if err != nil {
		localLogger.Error(ctx.Context(), "get all status exception", zap.Error(err))
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(materials)
}

// Update
// @Summary Update material
// @Description Update material by ID
// @Tags Materials
// @Accept json
// @Produce json
// @Param id path int true "Material ID"
// @Param input body dto.UpdateMaterialDTO true "Material data"
// @Success 200 {object} dto.MaterialDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /materials/{id} [put]
func (h *MaterialHandler) Update(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	body := new(dto.UpdateMaterialDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.materialService.Update(ctx.Context(), id, body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// Delete
// @Summary Delete material
// @Description Delete material by ID
// @Tags Materials
// @Accept json
// @Produce json
// @Param id path int true "Material ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /materials/{id} [delete]
func (h *MaterialHandler) Delete(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	err = h.materialService.Delete(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
