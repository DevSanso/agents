namespace InfoGatherHub.HubServer;
using System.Threading;

using InfoGatherHub.HubServer.Server;

public class MainThread
{
    private readonly TcpSocketServer server;
    public MainThread(TcpSocketServer server)
    {
        this.server = server;
    }

    public void Start()
    {
        while (true)
        {
            var client = server.AcceptTcpClient(5);
            if (client != null)
            {
                var clientThread = new ClientThread(client);
                clientThread.StartAsync();
            }
            Thread.Sleep(1000);
        }
    }
}