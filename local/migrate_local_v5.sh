migrations_dir=/tmp/migrations
echo "Executing  $migrations_dir/12-upgrade.sql..."
psql -U 'stratos' -d stratos_db -f "$migrations_dir/12-upgrade.sql"