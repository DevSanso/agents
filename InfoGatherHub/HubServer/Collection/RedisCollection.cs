namespace InfoGatherHub.HubServer.Collection;

using InfoGatherHub.HubServer.QueryBuilder.Redis;
using ps_redis = InfoGatherHub.HubProtos.Agent.Redis;

public class RedisCollection 
{
    private readonly InsertQueryBuilder queryBuilder = new();
    public void ClientList(ps_redis::RedisClientList list)
    {
        string query = queryBuilder.ClientListQuery(list);
    }
}