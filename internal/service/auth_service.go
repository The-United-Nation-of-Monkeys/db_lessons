package service

import (
	"context"
	"errors"

	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/dto"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/internal/repository"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/database"
	"github.com/The-United-Nation-of-Monkeys/db_lessons/pkg/jwt"
)

type AuthServiceInterface interface {
	Login(ctx context.Context, data *dto.LoginDTO) (*dto.AuthResponseDTO, error)
	RefreshToken(ctx context.Context, refreshToken string) (*dto.AuthResponseDTO, error)
}

type AuthService struct {
	repo     repository.AuthRepositoryInterface
	jwtService *jwt.ServiceJWT
	db       database.Config
}

func NewAuthService(repo repository.AuthRepositoryInterface, jwtService *jwt.ServiceJWT, db database.Config) *AuthService {
	return &AuthService{
		repo:      repo,
		jwtService: jwtService,
		db:        db,
	}
}

func (s *AuthService) Login(ctx context.Context, data *dto.LoginDTO) (*dto.AuthResponseDTO, error) {
	conn, err := database.NewConn(ctx, s.db)
	if err != nil {
		return nil, err
	}
	defer conn.Close(ctx)

	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	var user *dto.AuthResponseDTO

	switch data.Role {
	case "student":
		user, err = s.repo.LoginStudent(ctx, tx, data.Email, data.Password)
	case "teacher":
		user, err = s.repo.LoginTeacher(ctx, tx, data.Email, data.Password)
	case "admin":
		user, err = s.repo.LoginAdmin(ctx, tx, data.Email, data.Password)
	default:
		return nil, errors.New("invalid role")
	}

	if err != nil {
		return nil, err
	}

	// Generate tokens
	accessClaims := s.jwtService.GetClaims(user.UserID, user.Role, jwt.AccessToken)
	refreshClaims := s.jwtService.GetClaims(user.UserID, user.Role, jwt.RefreshToken)

	accessToken, err := s.jwtService.Encode(accessClaims)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.jwtService.Encode(refreshClaims)
	if err != nil {
		return nil, err
	}

	user.AccessToken = accessToken
	user.RefreshToken = refreshToken

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string) (*dto.AuthResponseDTO, error) {
	claims, err := s.jwtService.Decode(refreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	conn, err := database.NewConn(ctx, s.db)
	if err != nil {
		return nil, err
	}
	defer conn.Close(ctx)

	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	user, err := s.repo.GetUserByID(ctx, tx, claims.UserID, claims.Role)
	if err != nil {
		return nil, err
	}

	// Generate new tokens
	accessClaims := s.jwtService.GetClaims(user.UserID, user.Role, jwt.AccessToken)
	newRefreshClaims := s.jwtService.GetClaims(user.UserID, user.Role, jwt.RefreshToken)

	accessToken, err := s.jwtService.Encode(accessClaims)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := s.jwtService.Encode(newRefreshClaims)
	if err != nil {
		return nil, err
	}

	user.AccessToken = accessToken
	user.RefreshToken = newRefreshToken

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return user, nil
}

