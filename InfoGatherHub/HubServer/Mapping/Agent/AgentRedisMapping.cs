namespace InfoGatherHub.HubServer.Mapping.Agent;

using InfoGatherHub.HubServer.Mapping;
using InfoGatherHub.HubProtos.Agent.Redis;
using InfoGatherHub.HubServer.Collection;

public class AgentRedisMapping : IMapping<AgentRedisSnap>
{
    readonly RedisCollection collection = new();
    public void Run(string id, AgentRedisSnap snaps)
    {
        foreach(var snap in snaps.Datas)
        {
            switch(snap.Format)
            {
                case DataFormat.ClientLists:
                collection.ClientList(id, RedisClientList.Parser.ParseFrom(snap.RawData));
                break;
                case DataFormat.Dbsize:
                DbSize.Parser.ParseFrom(snap.RawData);
                break;
                case DataFormat.InfoCpu:
                RedisCpuInfo.Parser.ParseFrom(snap.RawData);
                break;
                case DataFormat.InfoMemory:
                RedisMemoryInfo.Parser.ParseFrom(snap.RawData);
                break;
                case DataFormat.InfoStat:
                RedisStatsInfo.Parser.ParseFrom(snap.RawData);
                break;
            }
        }
    }
}
