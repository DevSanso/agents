import ODBC

odbc_dsn = ARGS[1]

ODBC.setunixODBC()
conn = ODBC.Connection(odbc_dsn)

query = "INSERT INTO osnet_tcp4_stat_etc VALUES(logtime, object_id, sock_ref_count, sock_memory_location, retrans_timeout)
SELECT logtime, object_id, SPLIT_PART(etc,' ','1'), SPLIT_PART(etc,' ','2'), SPLIT_PART(etc,' ','3') 
    FROM osnet_tcp4_stat 
WHERE logtime > now() - TO_DATE('1 minute'::interval)
"
is_error = false
try
    DBInterface.execute(conn, query)
    write(stdout, "done")
catch e
    output = SubString("$e",1,1024)
    write(stderr,  output)
    is_error = true
    exit(2)
end

conn.close!()

if is_error == true
    exit(1)
else
    exit(0)
end
