namespace InfoGatherHub.HubServer;
using System.Threading;

using InfoGatherHub.HubProtos.Agent;
using InfoGatherHub.HubServer.Mapping;
using InfoGatherHub.HubServer.Server;


public class ClientThread
{
    private readonly TcpSocketClient client;
    private Thread? thread = null;
    private static IMapping<SnapData> mapping = new AgentMapping();
    private static void ClientThreadStartUp(TcpSocketClient client)
    {
        var memStream = new MemoryStream();
        while(true)
        {
            client.Read(ref memStream, 3);
            if(memStream.Length <= 0) continue;

            var snap = SnapData.Parser.ParseFrom(memStream.ToArray());
            mapping.Run(snap.Id, snap);
        }
    }
    public ClientThread(TcpSocketClient client)
    {
        this.client = client;
    }
    public void StartAsync()
    {
        thread = new Thread(() => ClientThreadStartUp(client));
        thread.IsBackground = true;
        thread.Start();
    }
}