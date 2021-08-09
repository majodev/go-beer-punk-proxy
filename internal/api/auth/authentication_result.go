package auth

import (
	"time"

	"github.com/majodev/go-beer-punk-proxy/internal/models"
)

type AuthenticationResult struct {
	Token      string
	User       *models.User
	ValidUntil time.Time
	Scopes     []string
}
