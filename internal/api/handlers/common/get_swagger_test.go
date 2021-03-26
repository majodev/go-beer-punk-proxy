package common_test

import (
	"net/http"
	"testing"

	"github.com/majodev/go-beer-punk-proxy/internal/api"
	"github.com/majodev/go-beer-punk-proxy/internal/test"
	"github.com/stretchr/testify/require"
)

func TestSwaggerYAMLRetrieval(t *testing.T) {
	test.WithTestServer(t, func(s *api.Server) {
		res := test.PerformRequest(t, s, "GET", "/swagger.yml", nil, nil)
		require.Equal(t, http.StatusOK, res.Result().StatusCode)
	})
}
