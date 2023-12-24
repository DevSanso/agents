create or replace create_partition()
returns int as
declare
    parts record;
begin

    select to_char(generate_series(now()::date, now()::date + interval '15 days', interval '1 day')::date, 'YYYY-MM-DD') as partition_date into parts;

    for tablename, table_schema in select * from partition_tables
    loop
        for part in parts
        loop
            -- create table if not exists table_schema.tablename_part partition of table_schema.tablename
            -- for values from(part) to (part + interval '23 hour 59 days 59 seconds');

            execute ' create table if not exists ' || table_schema || '.' || tablename || '_' || part ||
            ' partition of ' || table_schema || '.' || tablename || ' for values from (to_date(' || part || ', ' || 'YYYY-MM-DD))' ||
            ' to (to_date(' || part || ', ' || 'YYYY-MM-DD) + interval ' || '''' || '23 hour 59 days 59 seconds' ||  '''' || ');';
        end loop;
    end loop;

    
    return 0;
end;

