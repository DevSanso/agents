namespace InfoGatherHub.HubServer.Server;

using System;
using System.Collections.Concurrent;
using System.Net;
using System.Net.Sockets;

public class TcpSocketClient
{
    private readonly TcpClient client;
    internal TcpSocketClient(TcpClient client)
    {
        this.client = client;
    }

    public void Read(ref MemoryStream output)
    {
        using var stream = client.GetStream();
        output.CopyTo(stream);
    }

    public void Close()
    {
        client.Close();
    }

}