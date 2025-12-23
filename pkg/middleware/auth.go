package middleware

import (
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
		// Get actual request path and method
		// In Fiber v3, ctx.Path() returns the full path including the group prefix
		path := ctx.Path()
		method := ctx.Method()

		// Build dynamic API base path
		apiBasePath := fmt.Sprintf("/api/v%d", apiVersion)

		// Normalize path - remove query string if present
		if idx := strings.Index(path, "?"); idx != -1 {
			path = path[:idx]
		}

		// Public routes that don't require authentication (dynamic based on version)
		// Format: "METHOD /path" or just "/path" (for all methods)
		publicRoutes := map[string]bool{
			fmt.Sprintf("POST %s/auth/login", apiBasePath):   true,
			fmt.Sprintf("POST %s/auth/refresh", apiBasePath): true,
			fmt.Sprintf("GET %s/ping", apiBasePath):          true,
			fmt.Sprintf("POST %s/students", apiBasePath):     true, // Registration
			fmt.Sprintf("POST %s/teachers", apiBasePath):     true, // Registration
		}

		// Check if it's a public route (try both with method and without)
		routeKey := fmt.Sprintf("%s %s", method, path)
		if publicRoutes[routeKey] || publicRoutes[path] {
			return ctx.Next()
		}

		// Check if route requires authentication by matching against route configs
		// Try exact match with method first: "METHOD /path"
		requiredRoles, exists := routeConfigs[routeKey]

		// If not found, try match without method: "/path"
		if !exists {
			requiredRoles, exists = routeConfigs[path]
		}

		// If not found, try to match parameterized routes with method
		if !exists {
			for routeKey := range routeConfigs {
				// Check if routeKey contains method (format: "METHOD /path")
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
					// No method specified, match path only
					if matchRoute(path, routeKey) {
						requiredRoles = routeConfigs[routeKey]
						exists = true
						break
					}
				}
			}
		}

		// If route is not in configs and starts with API base path, require auth
		// (all API routes except public ones require authentication)
		if !exists && strings.HasPrefix(path, apiBasePath) {
			// Default: require any authenticated user (any role)
			requiredRoles = []string{"student", "teacher", "admin"}
			exists = true
		}

		// If route still doesn't exist, it's not an API route, allow it
		if !exists {
			return ctx.Next()
		}

		// Get token from cookie or header
		token := ctx.Cookies("access-token")
		if token == "" {
			// Try different header name variations (case-insensitive)
			authHeader := ctx.Get("Authorization")
			if authHeader == "" {
				authHeader = ctx.Get("authorization")
			}

			if authHeader != "" {
				// Remove "Bearer " prefix if present
				if strings.HasPrefix(authHeader, "Bearer ") {
					token = strings.TrimPrefix(authHeader, "Bearer ")
				} else if strings.HasPrefix(authHeader, "bearer ") {
					token = strings.TrimPrefix(authHeader, "bearer ")
				} else {
					// If no Bearer prefix, assume the whole value is the token
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

		// Store user info in context
		ctx.Locals("user_id", claims.UserID)
		ctx.Locals("user_role", claims.Role)

		return ctx.Next()
	}
}

// matchRoute checks if actual path matches route pattern
// e.g., /api/v1/students/123 matches /api/v1/students/:id
func matchRoute(actualPath, routePattern string) bool {
	actualParts := splitPath(actualPath)
	patternParts := splitPath(routePattern)

	if len(actualParts) != len(patternParts) {
		return false
	}

	for i := 0; i < len(actualParts); i++ {
		if len(patternParts[i]) > 0 && patternParts[i][0] == ':' {
			// Parameter placeholder, skip
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
