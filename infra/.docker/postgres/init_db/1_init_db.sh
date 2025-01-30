#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL

    CREATE USER myapp PASSWORD 'myapp';
    CREATE DATABASE myapp;
    GRANT ALL PRIVILEGES ON DATABASE myapp TO myapp;

    CREATE USER myapp_id PASSWORD 'myapp_id';
    CREATE DATABASE myapp_id;
    GRANT ALL PRIVILEGES ON DATABASE myapp_id TO myapp_id;


    \connect myapp;
    CREATE EXTENSION IF NOT EXISTS postgis;
    CREATE EXTENSION IF NOT EXISTS postgis_topology;
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
    GRANT ALL ON SCHEMA public TO myapp;
    GRANT USAGE ON SCHEMA topology to myapp;
    GRANT SELECT ON ALL SEQUENCES IN SCHEMA topology TO myapp;
    GRANT SELECT ON ALL TABLES IN SCHEMA topology TO myapp;

    \connect myapp_id;
    GRANT ALL ON SCHEMA public TO myapp_id;
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

EOSQL
