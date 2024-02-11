namespace InfoGatherHub.HubServer.Server;

using System;
using System.Collections.Concurrent;
using System.Net;
using System.Net.Sockets;

public class TcpSocketServer
{
    private TcpListener listener;

    public TcpSocketServer(string address, int port)
    {
        listener = new TcpListener(IPAddress.Parse(address), port);
        listener.Start();
    }
    private bool IsTimeOutException(Exception e)
    {
        return e is TaskCanceledException || e is TimeoutException || e is OperationCanceledException;
    }
    public TcpSocketClient? AcceptTcpClient(int timeoutSecond)
    {
        Task<TcpClient> clientTask = listener.AcceptTcpClientAsync();
        clientTask.Wait(timeoutSecond * 1000);
        clientTask.RunSynchronously();
        
        if(!clientTask.IsCompleted)
        {
            var exception = clientTask.Exception;
            if(exception != null && !exception.InnerExceptions.Any(IsTimeOutException))
            {
                throw exception;
            }
            return null;
        }

        var client = clientTask.Result;
        clientTask.Dispose();

        return new TcpSocketClient(client);
    }
    public void Stop()
    {
        listener.Stop();
    }
}
