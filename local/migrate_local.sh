migrations_dir=/tmp/migrations
for migration_file in "$migrations_dir"/*
do
  if [ "$migration_file" != "/tmp/migrations/12-upgrade.sql" ]; then
    echo "Executing  $migration_file..."
    psql -U 'stratos' -d stratos_db -f "$migration_file"
  else
    echo "Skipping $migration_file"
  fi
done