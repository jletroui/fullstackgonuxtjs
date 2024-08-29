#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER backend WITH PASSWORD 'backend';
    CREATE DATABASE backend;
    GRANT ALL PRIVILEGES ON DATABASE backend TO backend;
    CREATE DATABASE supertokens;
    CREATE USER supertokens WITH PASSWORD 'supertokens';
    GRANT ALL PRIVILEGES ON DATABASE supertokens TO supertokens;
EOSQL

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "supertokens" <<-EOSQL
    GRANT USAGE, CREATE ON SCHEMA public TO supertokens;
    ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL PRIVILEGES ON TABLES TO supertokens;
EOSQL
