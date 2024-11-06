-- messages_by_address modified in order to work with latest 3 partitions
CREATE OR REPLACE FUNCTION messages_by_address(
  addresses TEXT[],
  types TEXT[],
  "limit" BIGINT = 100,
  "offset" BIGINT = 0
)
RETURNS SETOF message AS $$
DECLARE
    v_partition_name text;  -- Renamed the variable to avoid conflict
    v_sql text := '';  -- Variable to hold dynamic SQL query
BEGIN
    FOR v_partition_name IN
        WITH latest_partitions AS (
            SELECT inhrelid::regclass::text AS partition_name
            FROM pg_inherits
            WHERE inhparent = 'message'::regclass
            AND inhrelid::regclass::text ~ '^message_\d+$'
        )
        SELECT partition_name
        FROM latest_partitions
        ORDER BY CAST(SUBSTRING(partition_name FROM '\d+$') AS INTEGER) DESC
        LIMIT 3
    LOOP
        IF v_sql = '' THEN
            v_sql := format('SELECT * FROM %I WHERE $1 && involved_accounts_addresses', v_partition_name);
        ELSE
            v_sql := v_sql || format(' UNION ALL SELECT * FROM %I WHERE $1 && involved_accounts_addresses', v_partition_name);
        END IF;
    END LOOP;

    v_sql := v_sql || format(' ORDER BY height DESC LIMIT $2 OFFSET $3');
    RETURN QUERY EXECUTE v_sql USING addresses, "limit", "offset";
END
$$ LANGUAGE plpgsql STABLE;

-- update_latest_transactions_view is used for message join in according to latest N partitions
CREATE OR REPLACE FUNCTION update_latest_transactions_view(
    "N" INT = 3
)
RETURNS void AS $$
DECLARE
    partitions text[];
    partition_name text;
    query text;
BEGIN
    SELECT array_agg(latest_partition_name)
    INTO partitions
    FROM (
        WITH latest_partitions AS (
        SELECT inhrelid::regclass::text AS latest_partition_name
        FROM pg_inherits
        WHERE inhparent = 'transaction'::regclass
        AND inhrelid::regclass::text ~ '^transaction_\d+$'
        )
        SELECT latest_partition_name
        FROM latest_partitions
        ORDER BY CAST(SUBSTRING(latest_partition_name FROM '\d+$') AS INTEGER) DESC
        LIMIT "N"
    ) AS latest;

    FOREACH partition_name IN ARRAY partitions LOOP
        RAISE NOTICE 'Partition to add into view: %', partition_name;
    END LOOP;

    query := 'CREATE OR REPLACE VIEW public.latest_transactions AS ';

    FOR i IN 1..array_length(partitions, 1) LOOP
        IF i > 1 THEN
        query := query || ' UNION ALL ';
        END IF;
        query := query || 'SELECT * FROM public.' || partitions[i];
    END LOOP;

    EXECUTE query;

    RAISE NOTICE 'View "latest_transactions" updated successfully.';
END;
$$ LANGUAGE plpgsql;
