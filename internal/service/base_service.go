package service

import (
	"context"

	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/database"
	"github.com/jackc/pgx/v5"
)

type RoleConfig struct {
	User string
}

type BaseService struct {
	role  map[string]RoleConfig
	cfgDB database.Config
}

func NewBaseService(roleMap map[string]RoleConfig, cfg database.Config) *BaseService {
	return &BaseService{
		role:  roleMap,
		cfgDB: cfg,
	}
}

func (b *BaseService) GetDBConn(role string) (*pgx.Conn, error) {
	ctx := context.Background()
	b.cfgDB.User = b.role[role].User
	dbConn, err := database.NewConn(ctx, b.cfgDB)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}
