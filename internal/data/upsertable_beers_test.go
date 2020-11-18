package data_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/majodev/go-beer-punk-proxy/internal/data"
	"github.com/majodev/go-beer-punk-proxy/internal/test"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestUpsertBeers(t *testing.T) {

	t.Parallel()

	test.WithTestDatabase(t, func(db *sql.DB) {

		beers, err := data.GetUpsertableBeerModels()
		require.NoError(t, err)

		tx, err := db.BeginTx(context.Background(), nil)
		require.NoError(t, err)

		for _, beer := range beers {
			err := beer.Upsert(context.Background(), tx, true, nil, boil.Infer(), boil.Infer())
			require.NoError(t, err)
		}

		err = tx.Commit()
		require.NoError(t, err)

	})

}
