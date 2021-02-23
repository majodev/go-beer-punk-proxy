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

func TestGetBeer(t *testing.T) {
	t.Parallel()

	test.WithTestServer(t, func(s *api.Server) {

		fixtures := test.Fixtures()

		res := test.PerformRequest(t, s, "GET", "/api/v1/beers/77", nil, test.HeadersWithAuth(t, fixtures.User1AccessToken1.Token))
		require.Equal(t, http.StatusOK, res.Result().StatusCode)

		var response types.GetBeersResponse
		test.ParseResponseAndValidate(t, res, &response)

		assert.Len(t, response, 1)

		assert.Equal(t, 77, int(*response[0].ID))

		test.Snapshoter.Save(t, response)
	})
}

func TestGetNonFoundBeer(t *testing.T) {

	t.Parallel()

	test.WithTestServer(t, func(s *api.Server) {
		fixtures := test.Fixtures()

		res := test.PerformRequest(t, s, "GET", "/api/v1/beers/1000", nil, test.HeadersWithAuth(t, fixtures.User1AccessToken1.Token))
		require.Equal(t, http.StatusNotFound, res.Result().StatusCode)

		var response types.PublicHTTPError
		test.ParseResponseAndValidate(t, res, &response)

		test.Snapshoter.Save(t, response)
	})

}

// Only here to showcase that an inaccessable database will result in a 500
// We won't actually hit the get_beers controller, as the authorization middleware will kick in before (try debugging)
func TestGetBeerWhileDBInaccessable(t *testing.T) {
	t.Parallel()

	test.WithTestServer(t, func(s *api.Server) {
		fixtures := test.Fixtures()

		s.DB.Close() // forcefully close the database!

		res := test.PerformRequest(t, s, "GET", "/api/v1/beers/1000", nil, test.HeadersWithAuth(t, fixtures.User1AccessToken1.Token))
		require.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)

		var response types.PublicHTTPError
		test.ParseResponseAndValidate(t, res, &response)

		test.Snapshoter.Save(t, response)
	})
}
