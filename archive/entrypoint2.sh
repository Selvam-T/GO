#!/bin/bash
set -e

# Export env
export DB_NAME=${DB_NAME}
export DB_USER=${DB_USER}
export DB_PASSWORD=${DB_PASSWORD}
export DB_HOST=${DB_HOST}
export DB_PORT=${DB_PORT}

# Start Postgres service
service postgresql start

# Wait for Postgres to start
until pg_isready -U postgres > /dev/null 2>&1; do
	echo "waiting for Postgres to start..."
	sleep 1
done

# Set session variable for init.sql
#psql -U postgres -d postgres -c "ALTER SYSTEM SET db_user TO '$DB_USER';"
#psql -U postgres -d postgres -c "ALTER SYSTEM SET db_password TO '$DB_PASSWORD';"
#psql -U postgres -d postgres -c "ALTER SYSTEM SET db_name TO '$DB_NAME';"

# Backup original config 
PG_HBA="/etc/postgresql/17/main/pg_hba.conf"

# Ensure root can write
chmod 644 "$PG_HBA"
cp "$PG_HBA" "$PG_HBA.bak" 

# Replace local and host rules to md5 in .conf
sed -i "s/peer/md5/g" "$PG_HBA"
sed -i "s/scram-sha-256/md5/g" "$PG_HBA"

# restore ownership (optional)
chown postgres:postgres "$PG_HBA"

# Restart Postgres to apply
service postgresql restart

# Execute SQL init script with environment variable substitution
envsubst < /docker-entrypoint-initdb.d/init.sql | psql -U postgres -h "$DB_HOST" -p "$DB_PORT"

exec bash # If I remove this, the process PIS 1 exits, Docker then kills the container
