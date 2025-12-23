package http

import (
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/service"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/exception"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/logger"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type SQLExecuteHandler struct {
	service service.SQLExecuteServiceInterface
}

func NewSQLExecuteHandler(service service.SQLExecuteServiceInterface) *SQLExecuteHandler {
	return &SQLExecuteHandler{service: service}
}

// @Summary Execute SQL query
// @Description Execute arbitrary SQL query in the database. This endpoint allows administrators to execute any SQL query (SELECT, INSERT, UPDATE, DELETE, etc.). For SELECT queries, it returns the results with columns and rows. For modification queries (INSERT, UPDATE, DELETE), it returns the number of affected rows. **WARNING: This endpoint has full database access and should only be used by trusted administrators.**
// @Tags SQL Execute
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body dto.ExecuteSQLDTO true "SQL query to execute"
// @Success 200 {object} dto.ExecuteSQLResponseDTO "Query executed successfully"
// @Failure 400 {object} map[string]string "Bad request - invalid query or query is required"
// @Failure 401 {object} map[string]string "Unauthorized - missing or invalid token"
// @Failure 403 {object} map[string]string "Forbidden - admin role required"
// @Failure 422 {object} map[string]string "Unprocessable entity - invalid request body"
// @Failure 500 {object} map[string]string "Internal server error - SQL execution failed"
// @Router /sql/execute [post]
func (h *SQLExecuteHandler) ExecuteSQL(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	body := new(dto.ExecuteSQLDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid request body")
	}

	if body.Query == "" {
		return exception.BadRequest("query is required")
	}

	localLogger.Info(ctx.Context(), "parse body", zap.String("query_preview", truncateString(body.Query, 100)))

	result, err := h.service.ExecuteSQL(ctx.Context(), body)
	if err != nil {
		localLogger.Error(ctx.Context(), "execute sql exception", zap.Error(err))
		return err
	}

	return ctx.JSON(result)
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
