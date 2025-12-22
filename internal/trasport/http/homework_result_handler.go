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

type HomeworkResultHandler struct {
	homeworkResultService service.HomeworkResultServiceInterface
}

func NewHomeworkResultHandler(homeworkResultService service.HomeworkResultServiceInterface) *HomeworkResultHandler {
	return &HomeworkResultHandler{
		homeworkResultService: homeworkResultService,
	}
}

// Create
// @Summary Create homeworkresult
// @Description Create a new homeworkresult
// @Tags HomeworkResults
// @Accept json
// @Produce json
// @Param input body dto.CreateHomeworkResultDTO true "HomeworkResult data"
// @Success 201 {object} dto.HomeworkResultDTO
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /homework-results [post]
func (h *HomeworkResultHandler) Create(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	body := new(dto.CreateHomeworkResultDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.homeworkResultService.Create(ctx.Context(), body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// GetByID
// @Summary Get homeworkresult by ID
// @Description Get homeworkresult by ID
// @Tags HomeworkResults
// @Accept json
// @Produce json
// @Param id path int true "HomeworkResult ID"
// @Success 200 {object} dto.HomeworkResultDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /homework-results/{id} [get]
func (h *HomeworkResultHandler) GetByID(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	homeworkresult, err := h.homeworkResultService.GetByID(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(homeworkresult)
}

// GetAll
// @Summary Get all homeworkresults
// @Description Get all homeworkresults
// @Tags HomeworkResults
// @Accept json
// @Produce json
// @Success 200 {array} dto.HomeworkResultDTO
// @Failure 500 {object} map[string]string
// @Router /homework-results [get]
func (h *HomeworkResultHandler) GetAll(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	homeworkresults, err := h.homeworkResultService.GetAll(ctx.Context())
	if err != nil {
		localLogger.Error(ctx.Context(), "get all status exception", zap.Error(err))
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(homeworkresults)
}

// Update
// @Summary Update homeworkresult
// @Description Update homeworkresult by ID
// @Tags HomeworkResults
// @Accept json
// @Produce json
// @Param id path int true "HomeworkResult ID"
// @Param input body dto.UpdateHomeworkResultDTO true "HomeworkResult data"
// @Success 200 {object} dto.HomeworkResultDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /homework-results/{id} [put]
func (h *HomeworkResultHandler) Update(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	body := new(dto.UpdateHomeworkResultDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.homeworkResultService.Update(ctx.Context(), id, body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// Delete
// @Summary Delete homeworkresult
// @Description Delete homeworkresult by ID
// @Tags HomeworkResults
// @Accept json
// @Produce json
// @Param id path int true "HomeworkResult ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /homework-results/{id} [delete]
func (h *HomeworkResultHandler) Delete(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	err = h.homeworkResultService.Delete(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
