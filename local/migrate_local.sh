migrations_dir=/tmp/migrations
for migration_file in "$migrations_dir"/*
do
  echo "Executing  $migration_file..."
  psql -U 'stratos' -d stratos_db -f "$migration_file"
done