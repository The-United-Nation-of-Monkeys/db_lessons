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

type SubcategoryHandler struct {
	subcategoryService service.SubcategoryServiceInterface
}

func NewSubcategoryHandler(subcategoryService service.SubcategoryServiceInterface) *SubcategoryHandler {
	return &SubcategoryHandler{
		subcategoryService: subcategoryService,
	}
}

// Create
// @Summary Create subcategory
// @Description Create a new subcategory
// @Tags Subcategories
// @Accept json
// @Produce json
// @Param input body dto.CreateSubcategoryDTO true "Subcategory data"
// @Success 201 {object} dto.SubcategoryDTO
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subcategories [post]
func (h *SubcategoryHandler) Create(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	body := new(dto.CreateSubcategoryDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.subcategoryService.Create(ctx.Context(), body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// GetByID
// @Summary Get subcategory by ID
// @Description Get subcategory by ID
// @Tags Subcategories
// @Accept json
// @Produce json
// @Param id path int true "Subcategory ID"
// @Success 200 {object} dto.SubcategoryDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /subcategories/{id} [get]
func (h *SubcategoryHandler) GetByID(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	subcategory, err := h.subcategoryService.GetByID(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(subcategory)
}

// GetAll
// @Summary Get all subcategorys
// @Description Get all subcategorys
// @Tags Subcategories
// @Accept json
// @Produce json
// @Success 200 {array} dto.SubcategoryDTO
// @Failure 500 {object} map[string]string
// @Router /subcategories [get]
func (h *SubcategoryHandler) GetAll(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	subcategorys, err := h.subcategoryService.GetAll(ctx.Context())
	if err != nil {
		localLogger.Error(ctx.Context(), "get all status exception", zap.Error(err))
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(subcategorys)
}

// Update
// @Summary Update subcategory
// @Description Update subcategory by ID
// @Tags Subcategories
// @Accept json
// @Produce json
// @Param id path int true "Subcategory ID"
// @Param input body dto.UpdateSubcategoryDTO true "Subcategory data"
// @Success 200 {object} dto.SubcategoryDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subcategories/{id} [put]
func (h *SubcategoryHandler) Update(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	body := new(dto.UpdateSubcategoryDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.subcategoryService.Update(ctx.Context(), id, body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// Delete
// @Summary Delete subcategory
// @Description Delete subcategory by ID
// @Tags Subcategories
// @Accept json
// @Produce json
// @Param id path int true "Subcategory ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subcategories/{id} [delete]
func (h *SubcategoryHandler) Delete(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	err = h.subcategoryService.Delete(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
