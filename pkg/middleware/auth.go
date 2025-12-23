package middleware

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/exception"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/jwt"
	"github.com/gofiber/fiber/v3"
)

type RouteConfig struct {
	Path   string
	Method string
	Roles  []string
}

func AuthMiddleware(jwtService *jwt.ServiceJWT, routeConfigs map[string][]string, apiVersion int) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		path := ctx.Path()
		method := ctx.Method()

		apiBasePath := fmt.Sprintf("/api/v%d", apiVersion)

		if idx := strings.Index(path, "?"); idx != -1 {
			path = path[:idx]
		}

		publicRoutes := map[string]bool{
			fmt.Sprintf("POST %s/auth/login", apiBasePath):   true,
			fmt.Sprintf("POST %s/auth/refresh", apiBasePath): true,
			fmt.Sprintf("GET %s/ping", apiBasePath):          true,
			fmt.Sprintf("POST %s/students", apiBasePath):     true, 
			fmt.Sprintf("POST %s/teachers", apiBasePath):     true, 
		}

		routeKey := fmt.Sprintf("%s %s", method, path)
		if publicRoutes[routeKey] || publicRoutes[path] {
			return ctx.Next()
		}

		requiredRoles, exists := routeConfigs[routeKey]

		if !exists {
			requiredRoles, exists = routeConfigs[path]
		}

		if !exists {
			for routeKey := range routeConfigs {
				parts := strings.SplitN(routeKey, " ", 2)
				if len(parts) == 2 {
					routeMethod := parts[0]
					routePath := parts[1]
					if routeMethod == method && matchRoute(path, routePath) {
						requiredRoles = routeConfigs[routeKey]
						exists = true
						break
					}
				} else {
					if matchRoute(path, routeKey) {
						requiredRoles = routeConfigs[routeKey]
						exists = true
						break
					}
				}
			}
		}

		if !exists && strings.HasPrefix(path, apiBasePath) {
			requiredRoles = []string{"student", "teacher", "admin"}
			exists = true
		}

		if !exists {
			return ctx.Next()
		}

		token := ctx.Cookies("access-token")
		if token == "" {
			authHeader := ctx.Get("Authorization")
			if authHeader == "" {
				authHeader = ctx.Get("authorization")
			}

			if authHeader != "" {
				if strings.HasPrefix(authHeader, "Bearer ") {
					token = strings.TrimPrefix(authHeader, "Bearer ")
				} else if strings.HasPrefix(authHeader, "bearer ") {
					token = strings.TrimPrefix(authHeader, "bearer ")
				} else {
					token = strings.TrimSpace(authHeader)
				}
			}
		}

		if token == "" {
			return exception.BadRequest("authentication required")
		}

		claims, err := jwtService.Decode(token)
		if err != nil {
			if errors.Is(err, jwt.InvalidTokenError) || errors.Is(err, jwt.UndefinedTokenError) {
				return exception.BadRequest("invalid token")
			}
			return exception.InternalServerError()
		}

		if len(requiredRoles) > 0 && !slices.Contains(requiredRoles, claims.Role) {
			return exception.BadRequest("insufficient permissions")
		}

		ctx.Locals("user_id", claims.UserID)
		ctx.Locals("user_role", claims.Role)

		userCtx := context.WithValue(ctx.Context(), "user_id", claims.UserID)
		userCtx = context.WithValue(userCtx, "user_role", claims.Role)
		ctx.SetContext(userCtx)

		return ctx.Next()
	}
}

func matchRoute(actualPath, routePattern string) bool {
	actualParts := splitPath(actualPath)
	patternParts := splitPath(routePattern)

	if len(actualParts) != len(patternParts) {
		return false
	}

	for i := 0; i < len(actualParts); i++ {
		if len(patternParts[i]) > 0 && patternParts[i][0] == ':' {
			continue
		}
		if actualParts[i] != patternParts[i] {
			return false
		}
	}
	return true
}

func splitPath(path string) []string {
	parts := []string{}
	current := ""
	for _, char := range path {
		if char == '/' {
			if current != "" {
				parts = append(parts, current)
				current = ""
			}
		} else {
			current += string(char)
		}
	}
	if current != "" {
		parts = append(parts, current)
	}
	return parts
}
