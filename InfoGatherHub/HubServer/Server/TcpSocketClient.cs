namespace InfoGatherHub.HubServer.Server;

using System.Net.Sockets;

using InfoGatherHub.HubGlobal;
using InfoGatherHub.HubGlobal.Logger;
using InfoGatherHub.HubServer.Config;
using InfoGatherHub.HubServer.Global.Extend;
public class TcpSocketClient
{
    private readonly TcpClient client;
    private readonly ILogger logger = GlobalProvider<Config, GlobalExtend>.Global();
    internal TcpSocketClient(TcpClient client)
    {
        this.client = client;
    }

    public void Read(ref MemoryStream output, int timeoutSecond)
    {
        client.ReceiveTimeout = timeoutSecond * 1000;
        try 
        {
            using var stream = client.GetStream();
            stream.CopyTo(output);
        }
        catch(IOException io)
        {
            logger.Log(LogLevel.Debug, LogCategory.Network, io.ToString());
            return;
        }
        catch(Exception)
        {
            throw;
        }
    }

    public void Close()
    {
        client.Close();
    }

}