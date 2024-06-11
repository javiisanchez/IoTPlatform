#!/bin/bash

set -e
export PGPASSWORD=iotPlatform;
psql -v ON_ERROR_STOP=1 --username "iotPlatform" --dbname "iotPlatform" <<-EOSQL
  CREATE DATABASE iotPlatform;
  GRANT ALL PRIVILEGES ON DATABASE iotPlatform TO "postgres";
EOSQL
