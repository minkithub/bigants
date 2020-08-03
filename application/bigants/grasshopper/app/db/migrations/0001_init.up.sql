CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE SCHEMA IF NOT EXISTS gh;

CREATE TABLE gh.user (
    id
        uuid PRIMARY KEY,
    created
        TIMESTAMPTZ NOT NULL,
    password
        VARCHAR(128) NOT NULL
);

CREATE TABLE gh.stock_prediction (
    id uuid
        PRIMARY KEY,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    stock_code VARCHAR(20) NOT NULL,
    created TIMESTAMPTZ NOT NULL,
    requester_id UUID
        REFERENCES gh.user
);

CREATE TABLE gh.stock_prediction_holiday (
    stock_prediction_id UUID NOT NULL
        REFERENCES gh.stock_prediction ON DELETE CASCADE,
    date DATE NOT NULL
);

CREATE TABLE gh.series_prediction (
    stock_prediction_id UUID NOT NULL
        REFERENCES gh.stock_prediction ON DELETE CASCADE,
    source VARCHAR(1) NOT NULL,
    created TIMESTAMPTZ NOT NULL,
    mae NUMERIC NOT NULL,
    mape NUMERIC NOT NULL
);

CREATE TABLE gh.series_daily_prediction (
    stock_prediction_id UUID NOT NULL
        REFERENCES gh.stock_prediction ON DELETE CASCADE,
    source VARCHAR(1) NOT NULL,
    date DATE NOT NULL,
    value NUMERIC NOT NULL
);
