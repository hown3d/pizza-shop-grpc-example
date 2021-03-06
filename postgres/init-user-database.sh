#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER bakery_user WITH PASSWORD $BAKERY_PASSWORD;
    CREATE DATABASE bakery;
    GRANT ALL PRIVILEGES ON DATABASE bakery TO bakery_user;
EOSQL