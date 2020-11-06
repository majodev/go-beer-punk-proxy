package beers

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/majodev/go-beer-punk-proxy/internal/models"
	"github.com/majodev/go-beer-punk-proxy/internal/types"
)

func MarschalBeers(beers models.BeerSlice) (types.GetBeersResponse, error) {
	response := make(types.GetBeersResponse, len(beers))

	for i, beer := range beers {

		var volume types.BoilVolume
		if err := volume.UnmarshalBinary(beer.Volume); err != nil {
			return response, err
		}

		var boilVolume types.BoilVolume
		if err := boilVolume.UnmarshalBinary(beer.BoilVolume); err != nil {
			return response, err
		}

		var method types.Method
		if err := method.UnmarshalBinary(beer.Method); err != nil {
			return response, err
		}

		var ingredients types.Ingredients
		if err := ingredients.UnmarshalBinary(beer.Ingredients); err != nil {
			return response, err
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

	return response, nil
}
