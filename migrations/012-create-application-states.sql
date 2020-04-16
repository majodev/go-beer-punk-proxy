-- +migrate Up
CREATE TABLE application_states (
    id uuid NOT NULL DEFAULT uuid_generate_v4 (),
    state text NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    CONSTRAINT application_states_pkey PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE application_states;
