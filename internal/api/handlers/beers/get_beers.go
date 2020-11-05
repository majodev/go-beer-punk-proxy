package beers

import (
	"net/http"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/labstack/echo/v4"
	"github.com/majodev/go-beer-punk-proxy/internal/api"
	"github.com/majodev/go-beer-punk-proxy/internal/models"
	"github.com/majodev/go-beer-punk-proxy/internal/types"
	"github.com/majodev/go-beer-punk-proxy/internal/types/beers"
	"github.com/majodev/go-beer-punk-proxy/internal/util"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func GetBeersRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1Beers.GET("", getBeersHandler(s))
}

func getBeersHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		params := beers.NewGetBeersRouteParams()
		log := util.LogFromContext(ctx)

		err := util.BindAndValidate(c, &params)
		if err != nil {
			return err
		}

		beers, err := models.Beers(
			qm.Limit(int(*params.PerPage)),
			qm.Offset(int(*params.PerPage)*(int(*params.Page)-1)),
			qm.OrderBy("id"),
		).All(ctx, s.DB)

		if err != nil {
			log.Debug().Err(err).Msg("Failed to load beers")
			return err
		}

		response := make(types.GetBeersResponse, len(beers))

		for i, beer := range beers {

			var volume types.BoilVolume
			if err := volume.UnmarshalBinary(beer.Volume); err != nil {
				return err
			}

			var boilVolume types.BoilVolume
			if err := boilVolume.UnmarshalBinary(beer.BoilVolume); err != nil {
				return err
			}

			var method types.Method
			if err := method.UnmarshalBinary(beer.Method); err != nil {
				return err
			}

			var ingredients types.Ingredients
			if err := ingredients.UnmarshalBinary(beer.Ingredients); err != nil {
				return err
			}

			response[i] = &types.Beer{
				ID:               swag.Int64(int64(beer.ID)),
				Name:             &beer.Name,
				Tagline:          &beer.Tagline,
				FirstBrewed:      swag.String(beer.FirstBrewed.Format("01/2006")),
				Description:      &beer.Description,
				ImageURL:         strfmt.URI(beer.ImageURL.String),
				Abv:              &beer.Abv,
				Ibu:              beer.Ibu,
				TargetFg:         swag.Int64(beer.TargetFG.Int64),
				TargetOg:         swag.Float64(beer.TargetOg.Float64),
				Ebc:              swag.Float64(beer.Ebc.Float64),
				Srm:              swag.Float64(beer.SRM.Float64),
				Ph:               swag.Float64(beer.PH.Float64),
				AttenuationLevel: swag.Float64(beer.AttenuationLevel.Float64),
				Volume:           &volume,
				BoilVolume:       &boilVolume,
				Method:           &method,
				Ingredients:      &ingredients,
				FoodPairing:      beer.FoodPairing,
				BrewersTips:      &beer.BrewersTips,
				ContributedBy:    swag.String(beer.ContributedBy),
			}
		}

		return util.ValidateAndReturn(c, http.StatusOK, response)

	}
}
