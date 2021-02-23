package beers_test

import (
	"net/http"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/majodev/go-beer-punk-proxy/internal/api"
	"github.com/majodev/go-beer-punk-proxy/internal/test"
	"github.com/majodev/go-beer-punk-proxy/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetBeers(t *testing.T) {
	t.Parallel()

	test.WithTestServer(t, func(s *api.Server) {

		fixtures := test.Fixtures()

		res := test.PerformRequest(t, s, "GET", "/api/v1/beers", nil, test.HeadersWithAuth(t, fixtures.User1AccessToken1.Token))
		require.Equal(t, http.StatusOK, res.Result().StatusCode)

		var response types.GetBeersResponse
		test.ParseResponseAndValidate(t, res, &response)

		require.Len(t, response, 25)

		// default check ordered by id
		for i, beer := range response {
			assert.Equal(t, i+1, int(*beer.ID))
		}

		test.Snapshoter.Save(t, response)
	})
}

func TestGetBeersFilterByBeerName(t *testing.T) {
	t.Parallel()

	test.WithTestServer(t, func(s *api.Server) {

		fixtures := test.Fixtures()

		res := test.PerformRequest(t, s, "GET", "/api/v1/beers?beer_name=pop", nil, test.HeadersWithAuth(t, fixtures.User1AccessToken1.Token))
		require.Equal(t, http.StatusOK, res.Result().StatusCode)

		var response types.GetBeersResponse
		test.ParseResponseAndValidate(t, res, &response)

		require.Len(t, response, 4)

		test.Snapshoter.Save(t, response)
	})
}

func TestGetBeersInvalidPerPageParam(t *testing.T) {
	t.Parallel()

	test.WithTestServer(t, func(s *api.Server) {

		fixtures := test.Fixtures()

		res := test.PerformRequest(t, s, "GET", "/api/v1/beers?per_page=100", nil, test.HeadersWithAuth(t, fixtures.User1AccessToken1.Token))
		require.Equal(t, http.StatusBadRequest, res.Result().StatusCode)

		var response types.PublicHTTPValidationError
		test.ParseResponseAndValidate(t, res, &response)

		test.Snapshoter.Save(t, response)
	})
}

// Try to implement tests for the additional params for yourself
// You should have a look at table-driven tests from here, as setting up each test will become quite tedious
// https://github.com/golang/go/wiki/TableDrivenTests

var paramtests = []struct {
	in         string
	outStatus  int
	outPayload runtime.Validatable
	beerCount  int // flag -1 if not a types.GetBeersResponse
}{
	{"/api/v1/beers?per_page=0", 400, new(types.PublicHTTPValidationError), -1},
	{"/api/v1/beers?beer_name=pop&page=1000", 200, new(types.GetBeersResponse), 0},
	{"/api/v1/beers?beer_name=pop&per_page=1&page=3", 200, new(types.GetBeersResponse), 1},
}

func TestGetBeersAllParams(t *testing.T) {
	t.Parallel()

	// we will use a single database / api.Server for each table driven tests (safe as our implementation for GET does not modify the db)...
	test.WithTestServer(t, func(s *api.Server) {

		fixtures := test.Fixtures()

		for _, tt := range paramtests {

			t.Run(tt.in, func(t *testing.T) {
				res := test.PerformRequest(t, s, "GET", tt.in, nil, test.HeadersWithAuth(t, fixtures.User1AccessToken1.Token))
				require.Equal(t, tt.outStatus, res.Result().StatusCode)

				test.ParseResponseAndValidate(t, res, tt.outPayload)

				if tt.beerCount >= 0 {

					// [...] to cast an interface into a concrete type, you do:
					// v = i.(T)
					// https://stackoverflow.com/questions/50939497/golang-cast-interface-to-struct
					beerRes := tt.outPayload.(*types.GetBeersResponse)
					require.Equal(t, tt.beerCount, len(*beerRes))
				}
			})
		}
	})

}
