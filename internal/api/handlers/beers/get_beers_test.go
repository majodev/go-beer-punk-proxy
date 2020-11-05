package beers_test

import (
	"net/http"
	"testing"

	"github.com/majodev/go-beer-punk-proxy/internal/api"
	"github.com/majodev/go-beer-punk-proxy/internal/test"
	"github.com/majodev/go-beer-punk-proxy/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetBeers(t *testing.T) {
	t.Parallel()

	test.WithTestServer(t, func(s *api.Server) {
		res := test.PerformRequest(t, s, "GET", "/api/v1/beers", nil, nil)
		require.Equal(t, http.StatusOK, res.Result().StatusCode)

		var response types.GetBeersResponse
		test.ParseResponseAndValidate(t, res, &response)

		assert.Len(t, response, 25)

		// default check ordered by id
		for i, beer := range response {
			assert.Equal(t, i+1, int(*beer.ID))
		}

		test.Snapshoter.Save(t, response)
	})
}
