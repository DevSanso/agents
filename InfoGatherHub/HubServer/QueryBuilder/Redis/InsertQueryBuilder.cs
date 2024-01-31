namespace InfoGatherHub.HubServer.QueryBuilder.Redis;

using SqlKata;
using SqlKata.Compilers;

using ps_redis = InfoGatherHub.HubProtos.Agent.Redis;

public class InsertQueryBuilder
{
    public string ClientListQuery(ps_redis::RedisClientList list)
    {
        Query queryBuilder = new Query("client_info");
        foreach(var client in list.Clients)
        {
            queryBuilder.AsInsert(new {
                collect_time = list.UnixEpoch
            });
        }
        var compiler = new SqlKata.Compilers.PostgresCompiler();
        var ret = compiler.Compile(queryBuilder);

        return ret.Sql;        
    }
}