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

type StudentAnswerHandler struct {
	studentAnswerService service.StudentAnswerServiceInterface
}

func NewStudentAnswerHandler(studentAnswerService service.StudentAnswerServiceInterface) *StudentAnswerHandler {
	return &StudentAnswerHandler{
		studentAnswerService: studentAnswerService,
	}
}

// @Summary Create studentanswer
// @Description Create a new studentanswer
// @Tags StudentAnswers
// @Accept json
// @Produce json
// @Param input body dto.CreateStudentAnswerDTO true "StudentAnswer data"
// @Success 201 {object} dto.StudentAnswerDTO
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /student-answers [post]
// @Security BearerAuth
func (h *StudentAnswerHandler) Create(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	body := new(dto.CreateStudentAnswerDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.studentAnswerService.Create(ctx.Context(), body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// @Summary Get studentanswer by ID
// @Description Get studentanswer by ID
// @Tags StudentAnswers
// @Accept json
// @Produce json
// @Param id path int true "StudentAnswer ID"
// @Success 200 {object} dto.StudentAnswerDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /student-answers/{id} [get]
// @Security BearerAuth
func (h *StudentAnswerHandler) GetByID(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	studentanswer, err := h.studentAnswerService.GetByID(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(studentanswer)
}

// @Summary Get all studentanswers
// @Description Get all studentanswers
// @Tags StudentAnswers
// @Accept json
// @Produce json
// @Success 200 {array} dto.StudentAnswerDTO
// @Failure 500 {object} map[string]string
// @Router /student-answers [get]
// @Security BearerAuth
func (h *StudentAnswerHandler) GetAll(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	studentanswers, err := h.studentAnswerService.GetAll(ctx.Context())
	if err != nil {
		localLogger.Error(ctx.Context(), "get all status exception", zap.Error(err))
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(studentanswers)
}

// @Summary Update studentanswer
// @Description Update studentanswer by ID
// @Tags StudentAnswers
// @Accept json
// @Produce json
// @Param id path int true "StudentAnswer ID"
// @Param input body dto.UpdateStudentAnswerDTO true "StudentAnswer data"
// @Success 200 {object} dto.StudentAnswerDTO
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /student-answers/{id} [put]
// @Security BearerAuth
func (h *StudentAnswerHandler) Update(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	body := new(dto.UpdateStudentAnswerDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity(err.Error())
	}
	localLogger.Info(ctx.Context(), "parse body")

	response, err := h.studentAnswerService.Update(ctx.Context(), id, body)
	if err != nil {
		return err
	}
	localLogger.Info(ctx.Context(), "get response")

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// @Summary Delete studentanswer
// @Description Delete studentanswer by ID
// @Tags StudentAnswers
// @Accept json
// @Produce json
// @Param id path int true "StudentAnswer ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /student-answers/{id} [delete]
// @Security BearerAuth
func (h *StudentAnswerHandler) Delete(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		localLogger.Info(ctx.Context(), "parse path param id exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid id")
	}

	err = h.studentAnswerService.Delete(ctx.Context(), id)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
