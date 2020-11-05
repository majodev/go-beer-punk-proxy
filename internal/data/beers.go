// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    beers, err := UnmarshalBeers(bytes)
//    bytes, err = beers.Marshal()

// see /docs/beers.json

package data

import "encoding/json"

type Beers []Beer

func UnmarshalBeers(data []byte) (Beers, error) {
	var r Beers
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Beers) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Beer struct {
	ID               int64       `json:"id"`
	Name             string      `json:"name"`
	Tagline          string      `json:"tagline"`
	FirstBrewed      string      `json:"first_brewed"`
	Description      string      `json:"description"`
	ImageURL         *string     `json:"image_url"`
	Abv              float64     `json:"abv"`
	Ibu              *float64    `json:"ibu"`
	TargetFg         *int64      `json:"target_fg"`
	TargetOg         *float64    `json:"target_og"`
	Ebc              *float64    `json:"ebc"`
	Srm              *float64    `json:"srm"`
	Ph               *float64    `json:"ph"`
	AttenuationLevel *float64    `json:"attenuation_level"`
	Volume           BoilVolume  `json:"volume"`
	BoilVolume       BoilVolume  `json:"boil_volume"`
	Method           Method      `json:"method"`
	Ingredients      Ingredients `json:"ingredients"`
	FoodPairing      []string    `json:"food_pairing"`
	BrewersTips      string      `json:"brewers_tips"`
	ContributedBy    string      `json:"contributed_by"`
}

type BoilVolume struct {
	Value *float64 `json:"value"`
	Unit  string   `json:"unit"`
}

type Ingredients struct {
	Malt  []Malt  `json:"malt"`
	Hops  []Hop   `json:"hops"`
	Yeast *string `json:"yeast"`
}

type Hop struct {
	Name      string     `json:"name"`
	Amount    BoilVolume `json:"amount"`
	Add       string     `json:"add"`
	Attribute string     `json:"attribute"`
}

type Malt struct {
	Name   string     `json:"name"`
	Amount BoilVolume `json:"amount"`
}

type Method struct {
	MashTemp     []MashTemp   `json:"mash_temp"`
	Fermentation Fermentation `json:"fermentation"`
	Twist        *string      `json:"twist"`
}

type Fermentation struct {
	Temp BoilVolume `json:"temp"`
}

type MashTemp struct {
	Temp     BoilVolume `json:"temp"`
	Duration *int64     `json:"duration"`
}
