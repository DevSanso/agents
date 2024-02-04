namespace InfoGatherHub.HubSender.Worker;

using System.Net.Sockets;
using System.Collections.Concurrent;

using Google.Protobuf;

using InfoGatherHub.HubCommon.Format;
using InfoGatherHub.HubCommon.Compress;
using InfoGatherHub.HubProtos.Agent;
using InfoGatherHub.HubSender.Worker.Format;

public class SendDataWorker : IWorker
{
    TcpClient client;
    ConcurrentQueue<IFormat<WorkerFormatHeader>> recvice;
    public SendDataWorker(TcpClient tcpClient, ConcurrentQueue<IFormat<WorkerFormatHeader>> recvice)
    {
        client = tcpClient;
        this.recvice = recvice;
    }
    public void Work()
    {
        IFormat<WorkerFormatHeader>? output;
        if(recvice.TryDequeue(out output) == false)return;

        using(var stream = client.GetStream())
        {
            string objectName = output.Header().ObjectName;
            string id = output.Header().Id;

            if(objectName == "OS")
            {
                SnapData snapData = new SnapData()
                {
                    RawSnap = ByteString.CopyFrom(output.Data()),
                    Format = SnapFormat.Os
                };

                byte []zip;
                ICompress comp = new Lz4Compress();

                comp.Compress(snapData.ToByteArray(), out zip);
                stream.Write(zip);
            }
            else if(objectName == "REDIS")
            {
                SnapData snapData = new SnapData()
                {
                    RawSnap = ByteString.CopyFrom(output.Data()),
                    Id = id,
                    Format = SnapFormat.Redis
                };

                byte []zip;
                ICompress comp = new Lz4Compress();

                comp.Compress(snapData.ToByteArray(), out zip);
                stream.Write(zip);
            }
        }
    }
    public void Dispose()
    {
        client.Close();
    }
}