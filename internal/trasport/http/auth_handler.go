package http

import (
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/service"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/exception"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/logger"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type AuthHandler struct {
	service service.AuthServiceInterface
}

func NewAuthHandler(service service.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{service: service}
}

// Login
// @Summary Login user
// @Description Login with email and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body dto.LoginDTO true "Login data"
// @Success 200 {object} dto.AuthResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/login [post]
func (h *AuthHandler) Login(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	body := new(dto.LoginDTO)
	if err := ctx.Bind().JSON(body); err != nil {
		localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
		return exception.UnprocessableEntity("invalid request body")
	}
	localLogger.Info(ctx.Context(), "parse body")

	user, err := h.service.Login(ctx.Context(), body)
	if err != nil {
		if err.Error() == "invalid email or password" {
			return exception.BadRequest(err.Error())
		}
		localLogger.Error(ctx.Context(), "login exception", zap.Error(err))
		return exception.InternalServerError()
	}

	// Set cookies
	ctx.Cookie(&fiber.Cookie{
		Name:     "access-token",
		Value:    user.AccessToken,
		HTTPOnly: false,
		Secure:   false,
		SameSite: "Lax",
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh-token",
		Value:    user.RefreshToken,
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
	})

	return ctx.JSON(user)
}

// RefreshToken
// @Summary Refresh access token
// @Description Refresh access token using refresh token
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body dto.RefreshTokenDTO true "Refresh token"
// @Success 200 {object} dto.AuthResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/refresh [post]
func (h *AuthHandler) RefreshToken(ctx fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(ctx.Context())

	// Try to get from cookie first
	refreshToken := ctx.Cookies("refresh-token")
	if refreshToken == "" {
		// Try to get from body
		body := new(dto.RefreshTokenDTO)
		if err := ctx.Bind().JSON(body); err != nil {
			localLogger.Info(ctx.Context(), "parse body exception", zap.Error(err))
			return exception.BadRequest("refresh token required")
		}
		refreshToken = body.RefreshToken
	}

	if refreshToken == "" {
		return exception.BadRequest("refresh token required")
	}

	user, err := h.service.RefreshToken(ctx.Context(), refreshToken)
	if err != nil {
		localLogger.Error(ctx.Context(), "refresh token exception", zap.Error(err))
		return exception.BadRequest("invalid refresh token")
	}

	// Set cookies
	ctx.Cookie(&fiber.Cookie{
		Name:     "access-token",
		Value:    user.AccessToken,
		HTTPOnly: false,
		Secure:   false,
		SameSite: "Lax",
	})

	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh-token",
		Value:    user.RefreshToken,
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
	})

	return ctx.JSON(user)
}
