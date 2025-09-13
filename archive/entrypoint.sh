#!/bin/bash
set -e

# Export env for envsubst and other commands
export DB_NAME=${DB_NAME}
export DB_USER=${DB_USER}
export DB_PASSWORD=${DB_PASSWORD}
export DB_HOST=${DB_HOST:-127.0.0.1}
export DB_PORT=${DB_PORT:-5432}

PG_VER="17"
PG_DATA="/var/lib/postgresql/${PG_VER}/main"
PG_CONF_DIR="/etc/postgresql/${PG_VER}/main"
PG_HBA="$PG_CONF_DIR/pg_hba.conf"
PG_CLUSTER_CONF="$PG_CONF_DIR/postgresql.conf"
PG_BIN_DIR="/usr/lib/postgresql/${PG_VER}/bin"

# Ensure data dir exists
mkdir -p "$PG_DATA"
chown -R postgres:postgres "$PG_DATA"

# Initialize cluster if needed
if [ ! -s "$PG_DATA/PG_VERSION" ]; then
    echo "Initializing database cluster..."
    gosu postgres "$PG_BIN_DIR/initdb" \
        --locale=C --encoding=UTF8 \
        --username=postgres \
        --pwfile=<(echo "$DB_PASSWORD") \
        -D "$PG_DATA"
fi

# Configure pg_hba.conf
cat > "$PG_HBA" <<EOL
local   all             all                                     md5
host    all             all             127.0.0.1/32            md5
host    all             all             ::1/128                 md5
EOL

# Configure postgresql.conf
sed -i "s/^#\?listen_addresses.*/listen_addresses = '127.0.0.1'/" "$PG_CLUSTER_CONF" || echo "listen_addresses = '127.0.0.1'" >> "$PG_CLUSTER_CONF"
sed -i "s/^#\?port.*/port = ${DB_PORT}/" "$PG_CLUSTER_CONF" || echo "port = ${DB_PORT}" >> "$PG_CLUSTER_CONF"

# Start postgres in background
gosu postgres "$PG_BIN_DIR/pg_ctl" -D "$PG_DATA" start

# Wait until ready
until PGPASSWORD="${DB_PASSWORD}" psql -U postgres -h "$DB_HOST" -p "$DB_PORT" -d postgres -c "SELECT 1;" >/dev/null 2>&1; do
    echo "Waiting for PostgreSQL to start..."
    sleep 1
done
echo "PostgreSQL is up!"

# Run init.sql with envsubst
if [ -f /docker-entrypoint-initdb.d/init.sql ]; then
    echo "Executing init.sql..."
    envsubst < /docker-entrypoint-initdb.d/init.sql | \
        PGPASSWORD="${DB_PASSWORD}" psql -U postgres -h "$DB_HOST" -p "$DB_PORT" -d postgres
fi

# Stop background process and restart in foreground
gosu postgres "$PG_BIN_DIR/pg_ctl" -D "$PG_DATA" stop
exec gosu postgres "$PG_BIN_DIR/postgres" -D "$PG_DATA" -c config_file="${PG_CLUSTER_CONF}"

