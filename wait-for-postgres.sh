#!/bin/sh
# wait-for-postgres.sh

set -e

host="$1"
shift
cmd="$@"

until PGPASSWORD="${DB_PASSWORD}" psql -h "$host" -U "postgres" -c '\q'; do
  >&2 echo "Postgres loading - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing migrations"
exec $cmd
