package beers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/majodev/go-beer-punk-proxy/internal/api"
	"github.com/majodev/go-beer-punk-proxy/internal/models"
	"github.com/majodev/go-beer-punk-proxy/internal/util"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func GetRandomBeerRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1Beers.GET("/random", getRandomBeerHandler(s))
}

func getRandomBeerHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		log := util.LogFromContext(ctx)

		beers, err := models.Beers(
			qm.OrderBy("RANDOM()"),
			qm.Limit(1),
		).All(ctx, s.DB)

		if err != nil {
			log.Debug().Err(err).Msg("Failed to load random beer")
			return err
		}

		response, err := MarschalBeers(beers)

		if err != nil {
			return err
		}

		return util.ValidateAndReturn(c, http.StatusOK, response)

	}
}
