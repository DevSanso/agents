namespace InfoGatherHub.HubSender.Worker;

using System.Net.Sockets;
using System.Collections.Concurrent;

using Google.Protobuf;

using InfoGatherHub.HubCommon.Format;
using InfoGatherHub.HubCommon.Compress;
using InfoGatherHub.HubProtos.Agent;

public class SendDataWorker : IWorker
{
    TcpClient client;
    ConcurrentQueue<IFormat<Void>> recvice;
    public SendDataWorker(TcpClient tcpClient, ConcurrentQueue<IFormat<Void>> recvice)
    {
        client = tcpClient;
        this.recvice = recvice;
    }
    public void Work()
    {
        IFormat<Void>? output;
        if(recvice.TryDequeue(out output) == false)return;

        using(var stream = client.GetStream())
        {
            string objectName = output.ObjectName();

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