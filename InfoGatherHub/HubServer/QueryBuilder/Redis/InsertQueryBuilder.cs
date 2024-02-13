namespace InfoGatherHub.HubServer.QueryBuilder.Redis;

using SqlKata;
using SqlKata.Compilers;

using ps_redis = InfoGatherHub.HubProtos.Agent.Redis;

public class InsertQueryBuilder
{
    private readonly static string[] COLUMN_NAME = new string[]
    {
        "object_id",
        "collect_time",
        "addr",
        "fd",
        "age",
        "idle",
        "flags",
        "db",
        "sub",
        "p_sub",
        "multi",
        "q_buf",
        "q_buf_free",
        "obl",
        "oll",
        "o_mem",
        "events",
        "cmd",
        "rbp",
        "user",
        "tot_mem",
        "s_sub",
        "resp",
        "redir",
        "rbs",
        "name",
        "lib_name",
        "lib_ver",
        "local_addr",
        "multi_mem",
        "argv_mem"
    };
    public string ClientListQuery(string id, ps_redis::RedisClientList list)
    {
        Query CollectTimeQueryBuilder = new ("");
        CollectTimeQueryBuilder.SelectRaw("TIMEZONE('Etc/UTC', TO_TIMESTAMP(?))", 1);
        
        object[] insertData = new object[list.Clients.Count];
        foreach(var client in list.Clients)
        {
            var obj = new
            {
                ObjectId = id,
                CollectTime = CollectTimeQueryBuilder,
                client.Addr,
                client.FD,
                client.Age,
                client.Idle,
                client.Flags,
                client.DB,
                client.Sub,
                client.PSub,
                client.Multi,
                client.QBuf,
                client.QBufFree,
                client.OBL,
                client.OLL,
                client.OMem,
                client.Events,
                client.Cmd,
                client.RBP,
                client.User,
                client.TotMem,
                client.SSub,
                client.Resp,
                client.Redir,
                client.RBS,
                client.Name,
                client.LibName,
                client.LibVer,
                client.LocalAddr,
                client.MultiMem,
                client.ArgvMem
            };
            _ = insertData.Append(obj);
        }
        Query queryBuilder = new ("redis.client_list");
        queryBuilder = queryBuilder.AsInsert(COLUMN_NAME, insertData);
        var compiler = new PostgresCompiler();
        var ret = compiler.Compile(queryBuilder);

        return ret.Sql;        
    }
}