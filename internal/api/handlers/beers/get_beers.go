package beers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/majodev/go-beer-punk-proxy/internal/api"
	"github.com/majodev/go-beer-punk-proxy/internal/data"
	"github.com/majodev/go-beer-punk-proxy/internal/models"
	"github.com/majodev/go-beer-punk-proxy/internal/types/beers"
	"github.com/majodev/go-beer-punk-proxy/internal/util"
	"github.com/majodev/go-beer-punk-proxy/internal/util/db"
	"github.com/volatiletech/null/v8"
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

		err := util.BindAndValidatePathAndQueryParams(c, &params)
		if err != nil {
			return err
		}

		queryMods := []qm.QueryMod{
			qm.Limit(int(*params.PerPage)),
			qm.Offset(int(*params.PerPage) * (int(*params.Page) - 1)),
			qm.OrderBy("id"),
		}

		// Where filters received via params...
		if params.AbvGt != nil {
			queryMods = append(queryMods, models.BeerWhere.Abv.GT(float64(*params.AbvGt)))
		}
		if params.AbvLt != nil {
			queryMods = append(queryMods, models.BeerWhere.Abv.LT(float64(*params.AbvLt)))
		}

		if params.IbuGt != nil {
			queryMods = append(queryMods, models.BeerWhere.Ibu.GT(null.Float64From(float64(*params.IbuGt))))
		}
		if params.IbuLt != nil {
			queryMods = append(queryMods, models.BeerWhere.Ibu.LT(null.Float64From(float64(*params.IbuLt))))
		}

		if params.EbcGt != nil {
			queryMods = append(queryMods, models.BeerWhere.Ebc.GT(null.Float64From(float64(*params.EbcGt))))
		}
		if params.EbcLt != nil {
			queryMods = append(queryMods, models.BeerWhere.Ebc.LT(null.Float64From(float64(*params.EbcLt))))
		}

		if params.BeerName != nil {
			val := fmt.Sprintf("%%%s%%", *params.BeerName)
			queryMods = append(queryMods, qm.Expr(db.ILike(val, models.TableNames.Beers, models.BeerColumns.Name)))
		}

		beers, err := models.Beers(queryMods...).All(ctx, s.DB)

		if err != nil {
			log.Debug().Err(err).Msg("Failed to load beers")
			return err
		}

		response, err := data.MarschalBeers(beers)

		if err != nil {
			return err
		}

		return util.ValidateAndReturn(c, http.StatusOK, response)

	}
}
