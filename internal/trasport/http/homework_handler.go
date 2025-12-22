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

type HomeworkHandler struct {
	homeworkService service.HomeworkServiceInterface
}

func NewHomeworkHandler(homeworkService service.HomeworkServiceInterface) *HomeworkHandler {
	return &HomeworkHandler{
		homeworkService: homeworkService,
	}
}

// Create
// @Summary Create homework
// @Description Create a new homework
// @Tags Homeworks
// @Accept json
// @Produce json
// @Param input body dto.CreateHomeworkDTO true "Homework data"
// @Success 201 {object} dto.HomeworkDTO
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /homeworks [post]
func (h *HomeworkHandler) Create(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	body := new(dto.CreateHomeworkDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.homeworkService.Create(ctx.Context(), body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// GetByID
// @Summary Get homework by ID
// @Description Get homework by ID
// @Tags Homeworks
// @Accept json
// @Produce json
// @Param id path int true "Homework ID"
// @Success 200 {object} dto.HomeworkDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /homeworks/{id} [get]
func (h *HomeworkHandler) GetByID(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	homework, err := h.homeworkService.GetByID(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(homework)
}

// GetAll
// @Summary Get all homeworks
// @Description Get all homeworks
// @Tags Homeworks
// @Accept json
// @Produce json
// @Success 200 {array} dto.HomeworkDTO
// @Failure 500 {object} map[string]string
// @Router /homeworks [get]
func (h *HomeworkHandler) GetAll(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	homeworks, err := h.homeworkService.GetAll(ctx.Context())
	if err != nil {
		localLogger.Error(ctx.Context(), "get all status exception", zap.Error(err))
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(homeworks)
}

// Update
// @Summary Update homework
// @Description Update homework by ID
// @Tags Homeworks
// @Accept json
// @Produce json
// @Param id path int true "Homework ID"
// @Param input body dto.UpdateHomeworkDTO true "Homework data"
// @Success 200 {object} dto.HomeworkDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /homeworks/{id} [put]
func (h *HomeworkHandler) Update(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	body := new(dto.UpdateHomeworkDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.homeworkService.Update(ctx.Context(), id, body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// Delete
// @Summary Delete homework
// @Description Delete homework by ID
// @Tags Homeworks
// @Accept json
// @Produce json
// @Param id path int true "Homework ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /homeworks/{id} [delete]
func (h *HomeworkHandler) Delete(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	err = h.homeworkService.Delete(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
