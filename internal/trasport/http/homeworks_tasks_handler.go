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

type HomeworksTasksHandler struct {
	homeworksTasksService service.HomeworksTasksServiceInterface
}

func NewHomeworksTasksHandler(homeworksTasksService service.HomeworksTasksServiceInterface) *HomeworksTasksHandler {
	return &HomeworksTasksHandler{
		homeworksTasksService: homeworksTasksService,
	}
}

// @Summary Create homework-task relation
// @Description Create a new relation between homework and task
// @Tags Homework Tasks
// @Accept json
// @Produce json
// @Param input body dto.CreateHomeworksTasksDTO true "Relation data"
// @Success 201 {object} dto.HomeworksTasksDTO
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /homeworks-tasks [post]
// @Security BearerAuth
func (h *HomeworksTasksHandler) Create(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	body := new(dto.CreateHomeworksTasksDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.homeworksTasksService.Create(ctx.Context(), body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// @Summary Get homework-task relation by IDs
// @Description Get relation by task ID and homework ID
// @Tags Homework Tasks
// @Accept json
// @Produce json
// @Param task_id path int true "Task ID"
// @Param homework_id path int true "Homework ID"
// @Success 200 {object} dto.HomeworksTasksDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /homeworks-tasks/{task_id}/{homework_id} [get]
// @Security BearerAuth
func (h *HomeworksTasksHandler) GetByID(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	taskIDStr := ctx.Params("task_id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param task_id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid task_id")
	}

	homeworkIDStr := ctx.Params("homework_id")
	homeworkID, err := strconv.Atoi(homeworkIDStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param homework_id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid homework_id")
	}

	relation, err := h.homeworksTasksService.GetByID(ctx.Context(), taskID, homeworkID)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(relation)
}

// @Summary Get all homework-task relations
// @Description Get all relations between homeworks and tasks
// @Tags Homework Tasks
// @Accept json
// @Produce json
// @Success 200 {array} dto.HomeworksTasksDTO
// @Failure 500 {object} map[string]string
// @Router /homeworks-tasks [get]
// @Security BearerAuth
func (h *HomeworksTasksHandler) GetAll(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	relations, err := h.homeworksTasksService.GetAll(ctx.Context())
	if err != nil {
		localLogger.Error(ctx.Context(), "get all relations exception", zap.Error(err))
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(relations)
}

// @Summary Delete homework-task relation
// @Description Delete relation by task ID and homework ID
// @Tags Homework Tasks
// @Accept json
// @Produce json
// @Param task_id path int true "Task ID"
// @Param homework_id path int true "Homework ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /homeworks-tasks/{task_id}/{homework_id} [delete]
// @Security BearerAuth
func (h *HomeworksTasksHandler) Delete(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	taskIDStr := ctx.Params("task_id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param task_id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid task_id")
	}

	homeworkIDStr := ctx.Params("homework_id")
	homeworkID, err := strconv.Atoi(homeworkIDStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param homework_id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid homework_id")
	}

	err = h.homeworksTasksService.Delete(ctx.Context(), taskID, homeworkID)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

