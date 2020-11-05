package data

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/majodev/go-beer-punk-proxy/internal/models"
	"github.com/majodev/go-beer-punk-proxy/internal/util"
	"github.com/volatiletech/null/v8"
)

func GetUpsertableBeerModels() ([]models.Beer, error) {

	file, err := ioutil.ReadFile(filepath.Join(util.GetProjectRootDir(), "/docs/beers.json"))

	if err != nil {
		return nil, err
	}

	beers, err := UnmarshalBeers([]byte(file))

	if err != nil {
		return nil, err
	}

	beersModels := make([]models.Beer, 0, len(beers))

	for _, beer := range beers {

		firstBrewed, err := time.Parse("01/2006", beer.FirstBrewed)

		if err != nil { // some dates are just YYYY
			firstBrewed, err = time.Parse("2006", beer.FirstBrewed)
		}

		if err != nil {
			return nil, err
		}

		volume, err := json.Marshal(beer.Volume)
		if err != nil {
			return nil, err
		}

		boilVolume, err := json.Marshal(beer.BoilVolume)
		if err != nil {
			return nil, err
		}

		method, err := json.Marshal(beer.Method)
		if err != nil {
			return nil, err
		}

		ingredients, err := json.Marshal(beer.Ingredients)
		if err != nil {
			return nil, err
		}

		newBeer := models.Beer{
			ID:               int(beer.ID),
			Name:             beer.Name,
			Tagline:          beer.Tagline,
			FirstBrewed:      firstBrewed,
			Description:      beer.Description,
			ImageURL:         null.StringFromPtr(beer.ImageURL),
			Abv:              beer.Abv,
			Ibu:              null.Float64FromPtr(beer.Ibu),
			TargetFG:         null.Int64FromPtr(beer.TargetFg),
			TargetOg:         null.Float64FromPtr(beer.TargetOg),
			Ebc:              null.Float64FromPtr(beer.Ebc),
			SRM:              null.Float64FromPtr(beer.Srm),
			PH:               null.Float64FromPtr(beer.Ph),
			AttenuationLevel: null.Float64FromPtr(beer.AttenuationLevel),
			Volume:           volume,
			BoilVolume:       boilVolume,
			Method:           method,
			Ingredients:      ingredients,
			FoodPairing:      beer.FoodPairing,
			BrewersTips:      beer.BrewersTips,
			ContributedBy:    beer.ContributedBy,
		}

		beersModels = append(beersModels, newBeer)
	}

	return beersModels, nil

}
