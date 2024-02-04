namespace InfoGatherHub.HubServer.Collection;

using InfoGatherHub.HubServer.QueryBuilder.Redis;
using InfoGatherHub.HubGlobal;
using InfoGatherHub.HubServer.Global.Extend;
using InfoGatherHub.HubServer.Global.Extend.DB;
using InfoGatherHub.HubServer.Config;
using ps_redis = InfoGatherHub.HubProtos.Agent.Redis;
using InfoGatherHub.HubGlobal.Extend;

public class RedisCollection 
{
    private readonly PgPool pool = GlobalProvider<Config, GlobalExtend>.Global().GetExtend().DbPool!;
    private readonly InsertQueryBuilder queryBuilder = new();
    public void ClientList(string id, ps_redis::RedisClientList list)
    {
        string query = queryBuilder.ClientListQuery(id, list);
        pool.Execute(query);
    }
}