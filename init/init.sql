-- set password for postgres role
ALTER USER postgres WITH PASSWORD '${DB_PASSWORD}';

-- Create user if it doesn't exist
CREATE ROLE IF NOT EXISTS ${DB_USER} LOGIN PASSWORD '${DB_PASSWORD}';

-- Create database if it doesn't exist
CREATE DATABASE ${DB_NAME} OWNER ${DB_USER};

-- Grant all privileges
GRANT ALL PRIVILEGES ON DATABASE ${DB_NAME} TO ${DB_USER};
