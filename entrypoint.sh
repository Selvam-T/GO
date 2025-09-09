#!/bin/bash

# Start Postgres service
service postgresql start

# Wait for Postgres to start
sleep 3

# Set session variable for init.sql
psql -U postgres -d postgres -c "ALTER SYSTEM SET db_user TO '$DB_USER';"
psql -U postgres -d postgres -c "ALTER SYSTEM SET db_password TO '$DB_PASSWORD';"
psql -U postgres -d postgres -c "ALTER SYSTEM SET db_name TO '$DB_NAME';"

# Execute SQL init script
psql -U postgres -f /docker-entrypoint-initdb.d/init.sql

exec bash # If I remove this, the process PIS 1 exits, Docker then kills the container
