package beers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/majodev/go-beer-punk-proxy/internal/api"
	"github.com/majodev/go-beer-punk-proxy/internal/models"
	"github.com/majodev/go-beer-punk-proxy/internal/types/beers"
	"github.com/majodev/go-beer-punk-proxy/internal/util"
)

func GetBeerRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1Beers.GET("/:id", getBeerHandler(s))
}

func getBeerHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		params := beers.NewGetBeerRouteParams()
		log := util.LogFromContext(ctx)

		err := util.BindAndValidate(c, &params)
		if err != nil {
			return err
		}

		beers, err := models.Beers(
			models.BeerWhere.ID.EQ(int(params.ID)),
		).All(ctx, s.DB)

		if err != nil {
			log.Debug().Err(err).Msg("Failed to load beer")
			return err
		}

		response, err := MarschalBeers(beers)

		if err != nil {
			return err
		}

		return util.ValidateAndReturn(c, http.StatusOK, response)

	}
}