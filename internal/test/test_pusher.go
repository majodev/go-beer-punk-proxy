package test

import (
	"database/sql"
	"testing"

	"github.com/majodev/go-beer-punk-proxy/internal/push"
	"github.com/majodev/go-beer-punk-proxy/internal/push/provider"
)

func WithTestPusher(t *testing.T, closure func(p *push.Service, db *sql.DB)) {
	t.Helper()

	WithTestDatabase(t, func(db *sql.DB) {
		t.Helper()
		closure(NewTestPusher(t, db), db)
	})
}

func NewTestPusher(t *testing.T, db *sql.DB) *push.Service {
	t.Helper()

	pushService := push.New(db)
	mockProvider := provider.NewMock(push.ProviderTypeFCM)
	pushService.RegisterProvider(mockProvider)

	return pushService
}
