package admin_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/majodev/go-beer-punk-proxy/internal/api"
	"github.com/majodev/go-beer-punk-proxy/internal/api/httperrors"
	"github.com/majodev/go-beer-punk-proxy/internal/models"
	"github.com/majodev/go-beer-punk-proxy/internal/test"
	"github.com/majodev/go-beer-punk-proxy/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// also try out
// var basePayload = test.GenericPayload{ [...] }

func getBasePayload() test.GenericPayload {
	return test.GenericPayload{
		"name":              "AAA Test beer",
		"tagline":           "all about apps",
		"first_brewed":      "2021",
		"description":       "This is our workshop test beer",
		"image_url":         nil,
		"abv":               12,
		"ibu":               75,
		"target_fg":         1025,
		"target_og":         1116,
		"ebc":               390,
		"srm":               198,
		"ph":                4.4,
		"attenuation_level": 78,
		"volume": test.GenericPayload{
			"value": 20,
			"unit":  "litres",
		},
		"boil_volume": test.GenericPayload{
			"value": 25,
			"unit":  "litres",
		},
		"method": test.GenericPayload{
			"mash_temp": []test.GenericPayload{
				{
					"temp": test.GenericPayload{
						"value": 64,
						"unit":  "celsius",
					},
					"duration": 135,
				},
			},
			"fermentation": test.GenericPayload{
				"temp": test.GenericPayload{
					"value": 21,
					"unit":  "celsius",
				},
			},
			"twist": nil,
		},
		"ingredients": test.GenericPayload{
			"malt": []test.GenericPayload{
				{
					"name": "Chocolate",
					"amount": test.GenericPayload{
						"value": 5.88,
						"unit":  "kilograms",
					},
				},
			},
			"hops": []test.GenericPayload{
				{
					"name": "Ground Coffee",
					"amount": test.GenericPayload{
						"value": 10,
						"unit":  "grams",
					},
					"add":       "Dry Hop",
					"attribute": "Flavour",
				},
			},
			"yeast": "Wyeast 1056 - American Ale™",
		},
		"food_pairing": []string{
			"Affogato with ice cream",
			"Chocolate cake",
		},
		"brewers_tips":   "If your mash run off takes a long time, increase the temperature of your sparge water to keep the mash bed warm and permeable.",
		"contributed_by": "Mario Ranftl <majodev>",
	}
}

func TestPostBeerInvalidISODate(t *testing.T) {
	t.Parallel()

	test.WithTestServer(t, func(s *api.Server) {
		payload := getBasePayload()
		payload["first_brewed"] = "2021-02-23T17:50:19.633Z"

		res := test.PerformRequest(t, s, "POST", "/api/v1/admin/beers?mgmt-secret="+s.Config.Management.Secret, payload, nil)
		require.Equal(t, http.StatusBadRequest, res.Result().StatusCode)

		var response httperrors.HTTPError
		test.ParseResponseAndValidate(t, res, &response)

		assert.Equal(t, "VALIDATE_FIRST_BREWED", *response.Type)
	})
}

func TestPostBeer(t *testing.T) {
	t.Parallel()

	test.WithTestServer(t, func(s *api.Server) {

		// payload := basePayload.Copy()
		payload := getBasePayload()

		// lets overwrite some fields...
		payload["name"] = "Another name"
		payload["ingredients"] = test.GenericPayload{
			"malt": []test.GenericPayload{},
			"hops": []test.GenericPayload{
				{
					"name": "AAA Coffee",
					"amount": test.GenericPayload{
						"value": 10,
						"unit":  "grams",
					},
					"add":       "Dry Hop",
					"attribute": "Flavour",
				},
			},
			"yeast": "Wyeast 1056 - Test",
		}

		res := test.PerformRequest(t, s, "POST", "/api/v1/admin/beers?mgmt-secret="+s.Config.Management.Secret, payload, nil)
		require.Equal(t, http.StatusOK, res.Result().StatusCode)

		var response types.Beer
		test.ParseResponseAndValidate(t, res, &response)

		assert.Equal(t, "Another name", *response.Name)
		assert.Len(t, response.Ingredients.Malt, 0)
		assert.Len(t, response.Ingredients.Hops, 1)
		assert.Equal(t, "AAA Coffee", *response.Ingredients.Hops[0].Name)

		count, err := models.Beers().Count(context.TODO(), s.DB)

		if err != nil {
			require.NoError(t, err)
		}

		assert.Equal(t, int64(326), count)
	})

	test.WithTestServer(t, func(s *api.Server) {

		// payload := basePayload.Copy()
		payload := getBasePayload()

		res := test.PerformRequest(t, s, "POST", "/api/v1/admin/beers?mgmt-secret="+s.Config.Management.Secret, payload, nil)
		require.Equal(t, http.StatusOK, res.Result().StatusCode)

		var response types.Beer
		test.ParseResponseAndValidate(t, res, &response)

		test.Snapshoter.Save(t, response)

		count, err := models.Beers().Count(context.TODO(), s.DB)

		if err != nil {
			require.NoError(t, err)
		}

		assert.Equal(t, int64(326), count)
	})
}
