-- +migrate Up
CREATE TABLE beers (
    -- Note that we do not prefer to use autoincremented integers, instead our primarys are UUIDs!
    -- However for the sake of compability with the existing Beer-Punk-API we'll use it here.
    -- Attention, bulk-inserting (predefined IDs) will require us the reset the sequence this primary uses later on
    -- This will automatically create the seq "beers_id_seq" -> SELECT nextval('beers_id_seq');
    id serial NOT NULL,
    name text NOT NULL,
    tagline text NOT NULL,
    first_brewed date NOT NULL,
    "description" text NOT NULL,
    image_url text,
    abv float NOT NULL,
    ibu float,
    target_fg bigint,
    target_og float,
    ebc float,
    srm float,
    ph float,
    attenuation_level float,
    volume jsonb NOT NULL,
    boil_volume jsonb NOT NULL,
    method jsonb NOT NULL,
    ingredients jsonb NOT NULL,
    food_pairing text[] NOT NULL,
    brewers_tips text NOT NULL,
    contributed_by text NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    CONSTRAINT beers_pkey PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE IF EXISTS beers;

