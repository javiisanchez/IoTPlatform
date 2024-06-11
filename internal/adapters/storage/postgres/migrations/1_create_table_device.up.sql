CREATE TABLE IF NOT EXISTS device(
    "ID" BIGSERIAL NOT NULL PRIMARY KEY,
    "Name" VARCHAR NOT NULL,
    "Location" VARCHAR,
    "Status" BOOLEAN NOT NULL,
    "CreatedAt" timestamptz NOT NULL DEFAULT (now()),
    "UpdatedAt" timestamptz NOT NULL DEFAULT (now())
);

