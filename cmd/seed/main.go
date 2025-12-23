package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/config"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/logger"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func main() {
	var (
		seedFile = flag.String("file", "scripts/seed_data.sql", "Path to seed SQL file")
	)
	flag.Parse()

	cfg, err := config.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load config: %v\n", err)
		os.Exit(1)
	}

	ctx := context.Background()
	ctx, err = logger.New(ctx, "seed")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	log := logger.GetLoggerFromCtx(ctx)
	if log == nil {
		fmt.Fprintf(os.Stderr, "Failed to get logger from context\n")
		os.Exit(1)
	}

	dbURL := fmt.Sprintf(
		"postgres:
		cfg.DataBase.User,
		cfg.DataBase.Password,
		cfg.DataBase.Host,
		cfg.DataBase.Port,
		cfg.DataBase.Name,
	)

	log.Info(ctx, "Connecting to database", zap.String("host", cfg.DataBase.Host), zap.String("database", cfg.DataBase.Name))

	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Error(ctx, "Failed to connect to database", zap.Error(err))
		os.Exit(1)
	}
	defer conn.Close(ctx)

	if err := conn.Ping(ctx); err != nil {
		log.Error(ctx, "Failed to ping database", zap.Error(err))
		os.Exit(1)
	}

	log.Info(ctx, "Connected to database successfully")

	seedPath, err := filepath.Abs(*seedFile)
	if err != nil {
		log.Error(ctx, "Failed to get absolute path", zap.Error(err))
		os.Exit(1)
	}

	log.Info(ctx, "Reading seed file", zap.String("path", seedPath))

	file, err := os.Open(seedPath)
	if err != nil {
		log.Error(ctx, "Failed to open seed file", zap.Error(err), zap.String("path", seedPath))
		os.Exit(1)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Error(ctx, "Failed to read seed file", zap.Error(err))
		os.Exit(1)
	}

	sqlContent := string(content)
	statements := parseSQL(sqlContent)

	log.Info(ctx, "Executing seed statements", zap.Int("count", len(statements)))

	tx, err := conn.Begin(ctx)
	if err != nil {
		log.Error(ctx, "Failed to begin transaction", zap.Error(err))
		os.Exit(1)
	}
	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
				log.Error(ctx, "Failed to rollback transaction", zap.Error(rollbackErr))
			}
		}
	}()

	successCount := 0
	for i, stmt := range statements {
		if strings.TrimSpace(stmt) == "" {
			continue
		}

		stmtUpper := strings.ToUpper(strings.TrimSpace(stmt))
		if strings.HasPrefix(stmtUpper, "TRUNCATE") && strings.Contains(stmtUpper, "CASCADE") {
			log.Info(ctx, "Skipping TRUNCATE statement", zap.Int("statement", i+1))
			continue
		}

		_, err = tx.Exec(ctx, stmt)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") || 
			   strings.Contains(err.Error(), "already exists") ||
			   strings.Contains(err.Error(), "ON CONFLICT") {
				log.Info(ctx, "Skipping duplicate entry", zap.Int("statement", i+1), zap.Error(err))
				successCount++
				continue
			}
			log.Error(ctx, "Failed to execute statement", 
				zap.Int("statement", i+1), 
				zap.String("preview", truncateString(stmt, 100)),
				zap.Error(err))
			return
		}
		successCount++
	}

	if err = tx.Commit(ctx); err != nil {
		log.Error(ctx, "Failed to commit transaction", zap.Error(err))
		os.Exit(1)
	}

	log.Info(ctx, "Database seeded successfully", 
		zap.Int("total_statements", len(statements)),
		zap.Int("executed_statements", successCount))

	displaySummary(ctx, conn, log)
}

func parseSQL(content string) []string {
	lines := strings.Split(content, "\n")
	var cleanedLines []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "--") {
			continue
		}
		if idx := strings.Index(line, "--"); idx != -1 {
			line = line[:idx]
		}
		cleanedLines = append(cleanedLines, line)
	}

	joined := strings.Join(cleanedLines, "\n")
	
	var statements []string
	var current strings.Builder
	var inString bool
	var stringChar byte
	
	for i := 0; i < len(joined); i++ {
		char := joined[i]
		
		if !inString {
			if char == '\'' || char == '"' {
				inString = true
				stringChar = char
				current.WriteByte(char)
			} else if char == ';' {
				stmt := strings.TrimSpace(current.String())
				if stmt != "" {
					statements = append(statements, stmt)
				}
				current.Reset()
			} else {
				current.WriteByte(char)
			}
		} else {
			current.WriteByte(char)
			if char == stringChar {
				if i+1 < len(joined) && joined[i+1] == stringChar {
					i++ 
					current.WriteByte(joined[i])
				} else {
					inString = false
				}
			}
		}
	}
	
	if current.Len() > 0 {
		stmt := strings.TrimSpace(current.String())
		if stmt != "" {
			statements = append(statements, stmt)
		}
	}
	
	return statements
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

func displaySummary(ctx context.Context, conn *pgx.Conn, log *logger.Logger) {
	queries := map[string]string{
		"students":   "SELECT COUNT(*) FROM student",
		"teachers":   "SELECT COUNT(*) FROM teacher",
		"courses":    "SELECT COUNT(*) FROM course",
		"lessons":    "SELECT COUNT(*) FROM lesson",
		"homeworks":  "SELECT COUNT(*) FROM homework",
		"tasks":      "SELECT COUNT(*) FROM task",
		"transactions": "SELECT COUNT(*) FROM \"transaction\"",
	}

	log.Info(ctx, "=== Seeding Summary ===")
	
	for name, query := range queries {
		var count int
		if err := conn.QueryRow(ctx, query).Scan(&count); err != nil {
			log.Error(ctx, "Failed to get count", zap.String("table", name), zap.Error(err))
			continue
		}
		log.Info(ctx, fmt.Sprintf("%s: %d", name, count))
	}
}

