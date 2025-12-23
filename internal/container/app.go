package container

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/config"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/container/initializer"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/container/server"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/database"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/redis"
)

func NewApp() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("get config error: %v", err)
	}

	log.Printf("Applying database migrations as user: %s", cfg.DataBase.User)
	if err := database.RunMigrations(ctx, cfg.DataBase); err != nil {
		log.Fatalf("migration error: %v", err)
	}

	redisConn, err := redis.New(cfg.Redis)
	if err != nil {
		log.Fatalf("redis connection error: %v", err)
	}

	repositoryList := initializer.NewRepositoryList()
	serviceList := initializer.NewServiceList(repositoryList, cfg)
	app := server.NewServer(cfg, serviceList, redisConn)

	serverPortStr := strconv.Itoa(int(cfg.Server.PortHttp))
	go func() {
		pid := os.Getpid()
		log.Printf("[PID %d] starting server...", pid)

		if err := app.Listen(":" + serverPortStr); err != nil {
			log.Printf("[PID %d] server listen error: %v", pid, err)
		}
	}()
	select {
	case <-ctx.Done():
		if err := app.Shutdown(); err != nil {
			log.Fatalf("server shutdown error: %v", err)
		}
		log.Printf("server stopped")
	}
}
