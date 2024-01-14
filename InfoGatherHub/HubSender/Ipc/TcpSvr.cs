namespace InfoGatherHub.HubSender.Ipc;

using System.Net.Sockets;
using System.Net;

class TcpSvr : IDisposable
{
    private class TcpSnapClient : ISnapClient
    {
        TcpClient client;
        
        public TcpSnapClient(TcpClient client)
        {
            this.client = client;
        }

        public void Dispose()
        {
            this.client.Close();
        }

        public void fetchSnapData()
        {
            return;
        }

        public byte[]? getSnapData()
        {
            byte[]? buf = new byte[1024 * 1024];
            int size = 0;
            using(var stream = client.GetStream())
            {
                size = stream.Read(buf);
            }
            if(size <= 0)buf = null;
            return buf;
        }
    }
    private TcpListener listener;
    public TcpSvr(int port)
    {
        listener = new TcpListener(IPAddress.Any, port);
    }
    public void Dispose()
    {
        listener.Stop();
    }
    public ISnapClient? GetClient(int timeoutSecond)
    {
        CancellationTokenSource source = new CancellationTokenSource();
        var clientTask = listener.AcceptTcpClientAsync(source.Token);

        Thread.Sleep(timeoutSecond * 1000);
        
        if(clientTask.IsCompleted)
        {
            source.Dispose();
            return new TcpSnapClient(clientTask.Result);
        }
        else
        {
            source.Cancel();
            source.Dispose();
            return null;
        }
    }
}