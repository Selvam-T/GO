-- init.sql
-- This file will be executed by entrypoint.sh
-- It assumes environment variables are passed by entrypoint.sh

-- set password for postgres role
ALTER USER postgres WITH PASSWORD '${DB_PASSWORD}';

-- Create user if it doesn't exist
DO
$do$
BEGIN
	IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname = '${DB_USER}') THEN
		EXECUTE format('CREATE USER %I WITH PASSWORD %L', '${DB_USER}', '${DB_PASSWORD}');
	END IF;
END
$do$;

-- Create database if it doesn't exist
DO
$do$
BEGIN
	IF NOT EXISTS (SELECT FROM pg_database WHERE datname = '${DB_NAME}') THEN
		EXECUTE format('CREATE DATABASE %I OWNER %I', '${DB_NAME}', '${DB_USER}');
END
$do$;

-- Grant all privileges
DO
$do$
BEGIN
	EXECUTE format('GRANT ALL PRIVILEGES ON DATABASE %I TO %I', '${DB_NAME}', '${DB_USER}');
END
$do$;
