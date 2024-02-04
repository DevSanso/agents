namespace InfoGatherHub.HubServer.QueryBuilder.Redis;

using SqlKata;
using SqlKata.Compilers;

using ps_redis = InfoGatherHub.HubProtos.Agent.Redis;

public class InsertQueryBuilder
{
    public string ClientListQuery(string id, ps_redis::RedisClientList list)
    {
        Query queryBuilder = new ("client_info");
        foreach(var client in list.Clients)
        {
            queryBuilder.AsInsert(new 
            {
                object_id = id,
                collect_time = list.UnixEpoch,
                id = client.ID,
                addr = client.Addr,
                local_addr = client.LocalAddr,
                fd = client.FD,
                name = client.Name,
                age = client.Age,
                idle = client.Idle,


            });
        }
        var compiler = new PostgresCompiler();
        var ret = compiler.Compile(queryBuilder);

        return ret.Sql;        
    }
}