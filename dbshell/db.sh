#!/bin/bash
set -e
export PGPASSWORD=postgres123;
psql -v ON_ERROR_STOP=1 --username "postgres" --dbname "youtubestore" <<-EOSQL
  CREATE DATABASE youtubestore;
  GRANT ALL PRIVILEGES ON DATABASE youtubestore TO "postgres";
EOSQL