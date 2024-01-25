namespace InfoGatherHub.HubSender.Worker;

using System.Collections.Concurrent;

using InfoGatherHub.HubSender.Ipc;
using InfoGatherHub.HubCommon.Format;
public class ReadOsSnapWorker : ReadMmapSnapWorker
{
    public ReadOsSnapWorker(ISnapClient client, ConcurrentQueue<IFormat<Void>> sender) : base(client, sender, "OS") {}
}