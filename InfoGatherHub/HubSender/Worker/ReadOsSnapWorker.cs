namespace InfoGatherHub.HubSender.Worker;

using System.Collections.Concurrent;

using InfoGatherHub.HubSender.Ipc;
using InfoGatherHub.HubCommon.Format;
using InfoGatherHub.HubSender.Worker.Format;

public class ReadOsSnapWorker : ReadMmapSnapWorker
{
    public ReadOsSnapWorker(ISnapClient client, ConcurrentQueue<IFormat<WorkerFormatHeader>> sender) : base(client, sender, "OS") {}
}