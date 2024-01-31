namespace InfoGatherHub.HubServer.Mapping;

using InfoGatherHub.HubProtos.Agent;

using ps_redis = InfoGatherHub.HubProtos.Agent.Redis;

using InfoGatherHub.HubServer.Mapping.Agent;

public class AgentMapping : IMapping<SnapData>
{
    private readonly AgentRedisMapping redisMapping = new();
    public void Run(SnapData snap)
    {
        switch(snap.Format)
        {
            case SnapFormat.Os:
            break;
            case SnapFormat.Redis:
                redisMapping.Run(ps_redis::AgentRedisSnap.Parser.ParseFrom(snap.RawSnap));
            break;
        }
    }
}