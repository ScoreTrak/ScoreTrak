package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/golang-jwt/jwt/v4"
)

type Manager struct {
	secretKey     string
	tokenDuration time.Duration
}

type claim string

const KeyClaim claim = "claim"

// UserClaims represents contents of JWT token.
type UserClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	TeamID   string `json:"team_id"`
	Role     string `json:"role"`
}

func NewJWTManager(secretKey string, tokenDuration time.Duration) *Manager {
	return &Manager{secretKey, tokenDuration}
}

type Config struct {
	Secret           string
	TimeoutInSeconds uint64
}

// Generate creates user claim based on passed user parameter, and encodes it to JWT token.
func (manager *Manager) Generate(user *user.User) (string, error) {
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(manager.tokenDuration)),
			ID:        user.ID.String(),
		},
		Username: user.Username,
		Role:     user.Role,
		TeamID:   user.TeamID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secretKey))
}

var ErrUnexpectedSigningToken = errors.New("unexpected token signing method")

// Verify ensures that the token provided by the client is valid, after which it extracts the claims and returns them.
func (manager *Manager) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, ErrUnexpectedSigningToken
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

var ErrInvalidToken = errors.New("invalid token claims")
