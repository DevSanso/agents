namespace InfoGatherHub.HubSender.Pusher.Console;

using InfoGatherHub.HubSender;
using InfoGatherHub.HubProtos.Agent;
using InfoGatherHub.HubGlobal.Logger.Extension.Line;
using InfoGatherHub.HubGlobal;

using ps_redis = InfoGatherHub.HubProtos.Agent.Redis;

using InfoGatherHub.HubSender.Pusher.Console.Agent;
using InfoGatherHub.HubGlobal.Logger;

public class AgentMapping : IMapping<SnapData>
{
    private readonly AgentRedisMapping redisMapping = new();
    public void Run(string id, SnapData snap)
    {
        switch(snap.Format)
        {
            case SnapFormat.Os:
            break;
            case SnapFormat.Redis:
                redisMapping.Run(snap.Id, ps_redis::AgentRedisSnap.Parser.ParseFrom(snap.RawSnap));
            break;
        }
    }

        public void Run(string id, byte[] rawData)
        {
            try 
            {
                var snap = SnapData.Parser.ParseFrom(rawData);
                Run(id, snap);
            }
            catch (System.Exception e)
            {
                GlobalProvider<Config>.Global().Log(LogLevel.Error, LogCategory.Code, e.Message);
                return;
            }
        }
}