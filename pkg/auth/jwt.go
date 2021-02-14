package auth

import (
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Manager struct {
	secretKey     string
	tokenDuration time.Duration
}

type claim string

const KeyClaim claim = "claim"

type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	TeamID   string `json:"team_id"`
	Role     string `json:"role"`
}

func NewJWTManager(secretKey string, tokenDuration time.Duration) *Manager {
	return &Manager{secretKey, tokenDuration}
}

type Config struct {
	Secret           string `default:"changeme"`
	TimeoutInSeconds uint64 `default:"86400"`
}

func (manager *Manager) Generate(user *user.User) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
			Id:        user.ID.String(),
		},
		Username: user.Username,
		Role:     user.Role,
		TeamID:   user.TeamID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secretKey))
}

func (manager *Manager) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
