-- init.sql
-- This file will be executed by entrypoint.sh
-- It assumes environment variables are passed by entrypoint.sh

-- Create user if it doesn't exist
DO
$do$
BEGIN
	IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname = current_setting('db_user')) THEN
		EXECUTE format('CREATE USER %I WITH PASSWORD %L', current_setting('db_user'), current_setting('db_password'));
	END IF;
END
$do$;

-- Create database if it doesn't exist
DO
$do$
BEGIN
	IF NOT EXISTS (SELECT FROM pg_database WHERE datname = current_setting('db_name')) THEN
		EXECUTE format('CREATE DATABASE %I OWNER %I', current_setting('db_name'), current_setting('db_user'));
END
$do$;

-- Grant all privileges
GRANT ALL PRIVILEGES ON DATABASE current_setting('db_name') TO current_setting('db_user');
