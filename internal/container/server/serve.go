package server

import (
	"time"

	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/config"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/container/initializer"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/trasport/http"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/exception"
	pkgvalidator "github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/validator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	redisstorage "github.com/gofiber/storage/redis/v3"
)

func NewServer(cfg *config.Config, serviceList *initializer.ServiceList, redisStg *redisstorage.Storage) *fiber.App {

	serverConfig := fiber.Config{
		AppName:         "Online Classes",
		StructValidator: pkgvalidator.Validator{Validator: validator.New()},
		ErrorHandler:    exception.Middleware,
		ReadTimeout:     10 * time.Second,
		WriteTimeout:    10 * time.Second,
		IdleTimeout:     30 * time.Second,
		ProxyHeader:     fiber.HeaderXForwardedFor,
	}
	server := fiber.New(serverConfig)

	http.NewController(server, cfg, serviceList, redisStg)
	return server

}
