namespace InfoGatherHub.HubSender.Worker;

using System.Net.Sockets;
using System.Collections.Concurrent;

using Google.Protobuf;

using InfoGatherHub.HubCommon.Format;
using InfoGatherHub.HubCommon.Compress;
using InfoGatherHub.HubProtos.Agent;
using InfoGatherHub.HubSender.Worker.Format;
using InfoGatherHub.HubSender.Pusher;

public class SendDataWorker : IWorker
{
    private ConcurrentQueue<IFormat<WorkerFormatHeader>> recvice;
    private IPusher<byte[]> pusher;
    public SendDataWorker(IPusher<byte[]> pusher, ConcurrentQueue<IFormat<WorkerFormatHeader>> recvice)
    {
        this.recvice = recvice;
        this.pusher = pusher;
    }
    public void Work()
    {
        IFormat<WorkerFormatHeader>? output;
        if(recvice.TryDequeue(out output) == false)return;

        string objectName = output.Header().ObjectName;
        string id = output.Header().Id;
        SnapFormat format = SnapFormat.Os;

        if(objectName == "REDIS")
        {
            format = SnapFormat.Redis;
        }

        SnapData snapData = new ()
        {
            RawSnap = ByteString.CopyFrom(output.Data()),
            Format = format
        };

        byte []zip;
        ICompress comp = new Lz4Compress();
        comp.Compress(snapData.ToByteArray(), out zip);

        pusher.Push(objectName, zip);
        
    }

    public void Dispose()
    {
        pusher.Dispose();
    }
}