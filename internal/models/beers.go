// Code generated by SQLBoiler 4.3.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// Beer is an object representing the database table.
type Beer struct {
	ID               int               `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name             string            `boil:"name" json:"name" toml:"name" yaml:"name"`
	Tagline          string            `boil:"tagline" json:"tagline" toml:"tagline" yaml:"tagline"`
	FirstBrewed      time.Time         `boil:"first_brewed" json:"first_brewed" toml:"first_brewed" yaml:"first_brewed"`
	Description      string            `boil:"description" json:"description" toml:"description" yaml:"description"`
	ImageURL         null.String       `boil:"image_url" json:"image_url,omitempty" toml:"image_url" yaml:"image_url,omitempty"`
	Abv              float64           `boil:"abv" json:"abv" toml:"abv" yaml:"abv"`
	Ibu              null.Float64      `boil:"ibu" json:"ibu,omitempty" toml:"ibu" yaml:"ibu,omitempty"`
	TargetFG         null.Int64        `boil:"target_fg" json:"target_fg,omitempty" toml:"target_fg" yaml:"target_fg,omitempty"`
	TargetOg         null.Float64      `boil:"target_og" json:"target_og,omitempty" toml:"target_og" yaml:"target_og,omitempty"`
	Ebc              null.Float64      `boil:"ebc" json:"ebc,omitempty" toml:"ebc" yaml:"ebc,omitempty"`
	SRM              null.Float64      `boil:"srm" json:"srm,omitempty" toml:"srm" yaml:"srm,omitempty"`
	PH               null.Float64      `boil:"ph" json:"ph,omitempty" toml:"ph" yaml:"ph,omitempty"`
	AttenuationLevel null.Float64      `boil:"attenuation_level" json:"attenuation_level,omitempty" toml:"attenuation_level" yaml:"attenuation_level,omitempty"`
	Volume           types.JSON        `boil:"volume" json:"volume" toml:"volume" yaml:"volume"`
	BoilVolume       types.JSON        `boil:"boil_volume" json:"boil_volume" toml:"boil_volume" yaml:"boil_volume"`
	Method           types.JSON        `boil:"method" json:"method" toml:"method" yaml:"method"`
	Ingredients      types.JSON        `boil:"ingredients" json:"ingredients" toml:"ingredients" yaml:"ingredients"`
	FoodPairing      types.StringArray `boil:"food_pairing" json:"food_pairing" toml:"food_pairing" yaml:"food_pairing"`
	BrewersTips      string            `boil:"brewers_tips" json:"brewers_tips" toml:"brewers_tips" yaml:"brewers_tips"`
	ContributedBy    string            `boil:"contributed_by" json:"contributed_by" toml:"contributed_by" yaml:"contributed_by"`
	CreatedAt        time.Time         `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt        time.Time         `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *beerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L beerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BeerColumns = struct {
	ID               string
	Name             string
	Tagline          string
	FirstBrewed      string
	Description      string
	ImageURL         string
	Abv              string
	Ibu              string
	TargetFG         string
	TargetOg         string
	Ebc              string
	SRM              string
	PH               string
	AttenuationLevel string
	Volume           string
	BoilVolume       string
	Method           string
	Ingredients      string
	FoodPairing      string
	BrewersTips      string
	ContributedBy    string
	CreatedAt        string
	UpdatedAt        string
}{
	ID:               "id",
	Name:             "name",
	Tagline:          "tagline",
	FirstBrewed:      "first_brewed",
	Description:      "description",
	ImageURL:         "image_url",
	Abv:              "abv",
	Ibu:              "ibu",
	TargetFG:         "target_fg",
	TargetOg:         "target_og",
	Ebc:              "ebc",
	SRM:              "srm",
	PH:               "ph",
	AttenuationLevel: "attenuation_level",
	Volume:           "volume",
	BoilVolume:       "boil_volume",
	Method:           "method",
	Ingredients:      "ingredients",
	FoodPairing:      "food_pairing",
	BrewersTips:      "brewers_tips",
	ContributedBy:    "contributed_by",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperfloat64 struct{ field string }

func (w whereHelperfloat64) EQ(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperfloat64) NEQ(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperfloat64) LT(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperfloat64) LTE(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperfloat64) GT(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperfloat64) GTE(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperfloat64) IN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperfloat64) NIN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_Float64 struct{ field string }

func (w whereHelpernull_Float64) EQ(x null.Float64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Float64) NEQ(x null.Float64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Float64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Float64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Float64) LT(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Float64) LTE(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Float64) GT(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Float64) GTE(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Int64 struct{ field string }

func (w whereHelpernull_Int64) EQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int64) NEQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Int64) LT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int64) LTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int64) GT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int64) GTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpertypes_JSON struct{ field string }

func (w whereHelpertypes_JSON) EQ(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertypes_JSON) NEQ(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertypes_JSON) LT(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_JSON) LTE(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_JSON) GT(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_JSON) GTE(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpertypes_StringArray struct{ field string }

func (w whereHelpertypes_StringArray) EQ(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertypes_StringArray) NEQ(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertypes_StringArray) LT(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_StringArray) LTE(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_StringArray) GT(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_StringArray) GTE(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var BeerWhere = struct {
	ID               whereHelperint
	Name             whereHelperstring
	Tagline          whereHelperstring
	FirstBrewed      whereHelpertime_Time
	Description      whereHelperstring
	ImageURL         whereHelpernull_String
	Abv              whereHelperfloat64
	Ibu              whereHelpernull_Float64
	TargetFG         whereHelpernull_Int64
	TargetOg         whereHelpernull_Float64
	Ebc              whereHelpernull_Float64
	SRM              whereHelpernull_Float64
	PH               whereHelpernull_Float64
	AttenuationLevel whereHelpernull_Float64
	Volume           whereHelpertypes_JSON
	BoilVolume       whereHelpertypes_JSON
	Method           whereHelpertypes_JSON
	Ingredients      whereHelpertypes_JSON
	FoodPairing      whereHelpertypes_StringArray
	BrewersTips      whereHelperstring
	ContributedBy    whereHelperstring
	CreatedAt        whereHelpertime_Time
	UpdatedAt        whereHelpertime_Time
}{
	ID:               whereHelperint{field: "\"beers\".\"id\""},
	Name:             whereHelperstring{field: "\"beers\".\"name\""},
	Tagline:          whereHelperstring{field: "\"beers\".\"tagline\""},
	FirstBrewed:      whereHelpertime_Time{field: "\"beers\".\"first_brewed\""},
	Description:      whereHelperstring{field: "\"beers\".\"description\""},
	ImageURL:         whereHelpernull_String{field: "\"beers\".\"image_url\""},
	Abv:              whereHelperfloat64{field: "\"beers\".\"abv\""},
	Ibu:              whereHelpernull_Float64{field: "\"beers\".\"ibu\""},
	TargetFG:         whereHelpernull_Int64{field: "\"beers\".\"target_fg\""},
	TargetOg:         whereHelpernull_Float64{field: "\"beers\".\"target_og\""},
	Ebc:              whereHelpernull_Float64{field: "\"beers\".\"ebc\""},
	SRM:              whereHelpernull_Float64{field: "\"beers\".\"srm\""},
	PH:               whereHelpernull_Float64{field: "\"beers\".\"ph\""},
	AttenuationLevel: whereHelpernull_Float64{field: "\"beers\".\"attenuation_level\""},
	Volume:           whereHelpertypes_JSON{field: "\"beers\".\"volume\""},
	BoilVolume:       whereHelpertypes_JSON{field: "\"beers\".\"boil_volume\""},
	Method:           whereHelpertypes_JSON{field: "\"beers\".\"method\""},
	Ingredients:      whereHelpertypes_JSON{field: "\"beers\".\"ingredients\""},
	FoodPairing:      whereHelpertypes_StringArray{field: "\"beers\".\"food_pairing\""},
	BrewersTips:      whereHelperstring{field: "\"beers\".\"brewers_tips\""},
	ContributedBy:    whereHelperstring{field: "\"beers\".\"contributed_by\""},
	CreatedAt:        whereHelpertime_Time{field: "\"beers\".\"created_at\""},
	UpdatedAt:        whereHelpertime_Time{field: "\"beers\".\"updated_at\""},
}

// BeerRels is where relationship names are stored.
var BeerRels = struct {
}{}

// beerR is where relationships are stored.
type beerR struct {
}

// NewStruct creates a new relationship struct
func (*beerR) NewStruct() *beerR {
	return &beerR{}
}

// beerL is where Load methods for each relationship are stored.
type beerL struct{}

var (
	beerAllColumns            = []string{"id", "name", "tagline", "first_brewed", "description", "image_url", "abv", "ibu", "target_fg", "target_og", "ebc", "srm", "ph", "attenuation_level", "volume", "boil_volume", "method", "ingredients", "food_pairing", "brewers_tips", "contributed_by", "created_at", "updated_at"}
	beerColumnsWithoutDefault = []string{"name", "tagline", "first_brewed", "description", "image_url", "abv", "ibu", "target_fg", "target_og", "ebc", "srm", "ph", "attenuation_level", "volume", "boil_volume", "method", "ingredients", "food_pairing", "brewers_tips", "contributed_by", "created_at", "updated_at"}
	beerColumnsWithDefault    = []string{"id"}
	beerPrimaryKeyColumns     = []string{"id"}
)

type (
	// BeerSlice is an alias for a slice of pointers to Beer.
	// This should generally be used opposed to []Beer.
	BeerSlice []*Beer

	beerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	beerType                 = reflect.TypeOf(&Beer{})
	beerMapping              = queries.MakeStructMapping(beerType)
	beerPrimaryKeyMapping, _ = queries.BindMapping(beerType, beerMapping, beerPrimaryKeyColumns)
	beerInsertCacheMut       sync.RWMutex
	beerInsertCache          = make(map[string]insertCache)
	beerUpdateCacheMut       sync.RWMutex
	beerUpdateCache          = make(map[string]updateCache)
	beerUpsertCacheMut       sync.RWMutex
	beerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single beer record from the query.
func (q beerQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Beer, error) {
	o := &Beer{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for beers")
	}

	return o, nil
}

// All returns all Beer records from the query.
func (q beerQuery) All(ctx context.Context, exec boil.ContextExecutor) (BeerSlice, error) {
	var o []*Beer

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Beer slice")
	}

	return o, nil
}

// Count returns the count of all Beer records in the query.
func (q beerQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count beers rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q beerQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if beers exists")
	}

	return count > 0, nil
}

// Beers retrieves all the records using an executor.
func Beers(mods ...qm.QueryMod) beerQuery {
	mods = append(mods, qm.From("\"beers\""))
	return beerQuery{NewQuery(mods...)}
}

// FindBeer retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBeer(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Beer, error) {
	beerObj := &Beer{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"beers\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, beerObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from beers")
	}

	return beerObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Beer) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no beers provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(beerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	beerInsertCacheMut.RLock()
	cache, cached := beerInsertCache[key]
	beerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			beerAllColumns,
			beerColumnsWithDefault,
			beerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(beerType, beerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(beerType, beerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"beers\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"beers\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into beers")
	}

	if !cached {
		beerInsertCacheMut.Lock()
		beerInsertCache[key] = cache
		beerInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Beer.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Beer) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	beerUpdateCacheMut.RLock()
	cache, cached := beerUpdateCache[key]
	beerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			beerAllColumns,
			beerPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update beers, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"beers\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, beerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(beerType, beerMapping, append(wl, beerPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update beers row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for beers")
	}

	if !cached {
		beerUpdateCacheMut.Lock()
		beerUpdateCache[key] = cache
		beerUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q beerQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for beers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for beers")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BeerSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), beerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"beers\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, beerPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in beer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all beer")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Beer) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no beers provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(beerColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	beerUpsertCacheMut.RLock()
	cache, cached := beerUpsertCache[key]
	beerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			beerAllColumns,
			beerColumnsWithDefault,
			beerColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			beerAllColumns,
			beerPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert beers, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(beerPrimaryKeyColumns))
			copy(conflict, beerPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"beers\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(beerType, beerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(beerType, beerMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert beers")
	}

	if !cached {
		beerUpsertCacheMut.Lock()
		beerUpsertCache[key] = cache
		beerUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Beer record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Beer) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Beer provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), beerPrimaryKeyMapping)
	sql := "DELETE FROM \"beers\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from beers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for beers")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q beerQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no beerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from beers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for beers")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BeerSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), beerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"beers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, beerPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from beer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for beers")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Beer) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBeer(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BeerSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BeerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), beerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"beers\".* FROM \"beers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, beerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BeerSlice")
	}

	*o = slice

	return nil
}

// BeerExists checks if the Beer row exists.
func BeerExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"beers\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if beers exists")
	}

	return exists, nil
}
