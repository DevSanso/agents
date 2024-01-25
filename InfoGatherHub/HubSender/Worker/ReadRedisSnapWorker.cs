namespace InfoGatherHub.HubSender.Worker;

using System.Collections.Concurrent;

using InfoGatherHub.HubSender.Ipc;
using InfoGatherHub.HubCommon.Format;
public class ReadRedisSnapWorker : ReadMmapSnapWorker
{
    public ReadRedisSnapWorker(ISnapClient client, ConcurrentQueue<IFormat<Void>> sender) : base(client, sender, "REDIS") {}
}