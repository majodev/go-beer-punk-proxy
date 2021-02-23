package admin

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/majodev/go-beer-punk-proxy/internal/api"
	"github.com/majodev/go-beer-punk-proxy/internal/api/httperrors"
	"github.com/majodev/go-beer-punk-proxy/internal/data"
	"github.com/majodev/go-beer-punk-proxy/internal/models"
	"github.com/majodev/go-beer-punk-proxy/internal/types"
	"github.com/majodev/go-beer-punk-proxy/internal/types/admin"
	"github.com/majodev/go-beer-punk-proxy/internal/util"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func GetBeerRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1Admin.POST("/beers", postAdminBeerHandler(s))
}

func postAdminBeerHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		params := admin.NewPostBeerRouteParams()
		log := util.LogFromContext(ctx)

		err := util.BindAndValidatePathAndQueryParams(c, &params)
		if err != nil {
			return err
		}

		var body types.PostBeerPayload
		if err := util.BindAndValidateBody(c, &body); err != nil {
			return err
		}

		firstBrewed, err := time.Parse("01/2006", *body.FirstBrewed)

		if err != nil { // some dates are just YYYY
			firstBrewed, err = time.Parse("2006", *body.FirstBrewed)
		}

		if err != nil {
			return httperrors.NewHTTPErrorWithDetail(400, "VALIDATE_FIRST_BREWED", "Payload invalid", "Could not parse first_brewed, this field either requires a format ala '01/2006' or '2006'.")
		}

		volume, err := json.Marshal(body.Volume)
		if err != nil {
			return err
		}

		boilVolume, err := json.Marshal(body.BoilVolume)
		if err != nil {
			return err
		}

		method, err := json.Marshal(body.Method)
		if err != nil {
			return err
		}

		ingredients, err := json.Marshal(body.Ingredients)
		if err != nil {
			return err
		}

		newBeer := models.Beer{
			Name:             *body.Name,
			Tagline:          *body.Tagline,
			FirstBrewed:      firstBrewed,
			Description:      *body.Description,
			ImageURL:         null.StringFrom(body.ImageURL.String()),
			Abv:              *body.Abv,
			Ibu:              null.Float64From(body.Ibu),
			TargetFG:         null.Int64FromPtr(body.TargetFg),
			TargetOg:         null.Float64FromPtr(body.TargetOg),
			Ebc:              null.Float64From(body.Ebc),
			SRM:              null.Float64From(body.Srm),
			PH:               null.Float64From(body.Ph),
			AttenuationLevel: null.Float64FromPtr(body.AttenuationLevel),
			Volume:           volume,
			BoilVolume:       boilVolume,
			Method:           method,
			Ingredients:      ingredients,
			FoodPairing:      body.FoodPairing,
			BrewersTips:      *body.BrewersTips,
			ContributedBy:    *body.ContributedBy,
		}

		// ctx = boil.WithDebug(ctx, true)

		if err := newBeer.Insert(ctx, s.DB, boil.Infer()); err != nil {
			log.Debug().Err(err).Msg("Failed to load beer")
			return err
		}

		response, err := data.MarschalBeer(newBeer)

		if err != nil {
			return err
		}

		return util.ValidateAndReturn(c, http.StatusOK, response)

	}
}
