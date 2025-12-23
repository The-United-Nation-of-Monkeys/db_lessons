package jwt

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	InvalidTokenStringError = "token is invalid"
)

var (
	UndefinedTokenError = fmt.Errorf("undefined token")
	InvalidTokenError   = fmt.Errorf(InvalidTokenStringError)
)

type TokenMode string

const (
	RefreshToken TokenMode = "refresh"
	AccessToken  TokenMode = "access"
)

const (
	RefreshTokenCookieName = "refresh-token"
	AccessTokenCookieName  = "access-token"
)

type Claims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type ServiceJWT struct {
	privateKey     *rsa.PrivateKey
	publicKey      *rsa.PublicKey
	RefreshTimeExp time.Duration
	AccessTimeExp  time.Duration
}

func NewServiceJWT(
	privateKey *rsa.PrivateKey,
	publicKey *rsa.PublicKey,
	refreshTimeExp time.Duration,
	accessTimeExp time.Duration,
) *ServiceJWT {
	return &ServiceJWT{
		privateKey:     privateKey,
		publicKey:      publicKey,
		RefreshTimeExp: refreshTimeExp,
		AccessTimeExp:  accessTimeExp,
	}
}

func (j *ServiceJWT) GetClaims(userID int, role string, tokenMode TokenMode) *Claims {
	var expiration *jwt.NumericDate

	if tokenMode == RefreshToken {
		expiration = jwt.NewNumericDate(time.Now().Add(j.RefreshTimeExp))
	} else if tokenMode == AccessToken {
		expiration = jwt.NewNumericDate(time.Now().Add(j.AccessTimeExp))
	} else {
		panic("invalid token mode")
	}

	return &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   fmt.Sprintf("%d", userID),
			ExpiresAt: expiration,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}

func (j *ServiceJWT) Encode(claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(j.privateKey)
	if err != nil {
		return "", fmt.Errorf("failed create token: %v", err)
	}

	return tokenString, nil
}

func (j *ServiceJWT) Decode(tokenString string) (*Claims, error) {
	if tokenString == "" {
		return nil, UndefinedTokenError
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return j.publicKey, nil
		},
	)

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, InvalidTokenError
		}
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, InvalidTokenError
	}

	return claims, nil
}

