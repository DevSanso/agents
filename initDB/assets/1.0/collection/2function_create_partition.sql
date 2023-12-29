create or replace function collection.create_partition()
returns int as $$
declare
 	part record;
   	tablename varchar(256);
	table_schema varchar(256);
begin

    for tablename, table_schema in select * from collection.partition_tables
    loop
        for part in select part_date from 
        	(
        	select 
        		 to_char(generate_series(now()::date, now()::date + interval '15 days', interval '1 day')::date, 'YYYY_MM_DD') 
        		as part_date
        	) as a
        loop
            -- create table if not exists table_schema.tablename_part partition of table_schema.tablename
            -- for values from(part) to (part + interval '23 hour 59 days 59 seconds');

            execute ' create table if not exists ' || table_schema || '.' || tablename || '_' || part.part_date ||
            ' partition of ' || table_schema || '.' || tablename || ' for values from (to_date(''' || part.part_date || ''', ' || '''YYYY_MM_DD' || '''' || '))' ||
            ' to (to_date(''' || part.part_date || ''', ' || '''YYYY_MM_DD'') + interval ' || '''' || '23 hour 59 minute 59 seconds' ||  '''' || ');';
        end loop;
    end loop;

    return 0;
end;
$$ LANGUAGE plpgsql;