package data_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
	"time"

	"github.com/majodev/go-beer-punk-proxy/internal/data"
	"github.com/majodev/go-beer-punk-proxy/internal/models"
	"github.com/majodev/go-beer-punk-proxy/internal/test"
	"github.com/majodev/go-beer-punk-proxy/internal/util"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestInsertBeers(t *testing.T) {

	t.Parallel()

	test.WithTestDatabase(t, func(db *sql.DB) {

		file, err := ioutil.ReadFile(filepath.Join(util.GetProjectRootDir(), "/docs/beers.json"))

		require.NoError(t, err)

		beers, err := data.UnmarshalBeers([]byte(file))
		require.NoError(t, err)

		for _, beer := range beers {

			fmt.Println(beer.ID, beer.Name)
			firstBrewed, err := time.Parse("01/2006", beer.FirstBrewed)

			if err != nil { // some dates are just YYYY
				firstBrewed, err = time.Parse("2006", beer.FirstBrewed)
			}

			require.NoError(t, err)

			volume, err := json.Marshal(beer.Volume)
			require.NoError(t, err)

			boilVolume, err := json.Marshal(beer.BoilVolume)
			require.NoError(t, err)

			method, err := json.Marshal(beer.Method)
			require.NoError(t, err)

			ingredients, err := json.Marshal(beer.Ingredients)
			require.NoError(t, err)

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

			err = newBeer.Insert(context.Background(), db, boil.Infer())

			require.NoError(t, err)
		}

	})

}
