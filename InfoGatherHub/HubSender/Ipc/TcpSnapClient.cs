namespace InfoGatherHub.HubSender.Ipc;

using System.Net.Sockets;
using System.Collections.Concurrent;
using System.Net;

using InfoGatherHub.HubCommon.Sequence;
using System.Net.Http.Headers;

class TcpSnapClient : ISnapClient
{
    private bool disposedValue;
    private TcpListener listener;
    private LongSequence seq = new LongSequence();
    private LongSequence readSeq = new LongSequence();
    private ConcurrentDictionary<long, TcpClient> clients = new ConcurrentDictionary<long, TcpClient>();
    public TcpSnapClient(int port)
    {
        listener = new TcpListener(IPAddress.Any, port);
    }
    public void AcceptClient()
    {
        CancellationTokenSource source = new CancellationTokenSource();
        var clientTask = listener.AcceptTcpClientAsync(source.Token);

        Thread.Sleep(100);
        
        if(clientTask.IsCompleted)
        {
            clients.TryAdd(seq.Next(), clientTask.Result);
        }
        else
        {
            source.Cancel();
        }

        source.Dispose();
    }
    public void ClearCloseClient()
    {
        List<long> keys = new List<long>();
        foreach(var client in clients)
        {
            if(!client.Value.Connected)
            {
                keys.Add(client.Key);
            }
        }
        TcpClient? output;

        foreach(var k in keys)
        {
            if(clients.TryGetValue(k, out output))
            {
                output.Dispose();
                clients.Remove(k, out output);
                output = null;
            }
        }
    }
    public void FetchSnapData()
    {
        var i = readSeq.Next();
        if(i >= clients.Count) readSeq.Reset();
    }

    public byte[]? GetSnapData()
    {
        TcpClient? client = null;
        if(!clients.TryGetValue(readSeq.Current(), out client))
        {
            return null;
        }

        var mem = new MemoryStream();
        client.ReceiveTimeout = 1000;
        
        using(var stream = client.GetStream())
        {
            mem.CopyTo(stream);
        }

        client.ReceiveTimeout = 0;
        if(mem.Length > 0)
        {
            var output = new byte[mem.Length];
            mem.Write(output, 0, (int)mem.Length);
            mem.Dispose();
            mem = null;
            return output;
        }
        mem.Dispose();
        mem = null;
        return null;
    }

    protected virtual void Dispose(bool disposing)
    {
        if (!disposedValue)
        {
            if (disposing)
            {
                listener.Stop();
            }

            // TODO: 비관리형 리소스(비관리형 개체)를 해제하고 종료자를 재정의합니다.
            // TODO: 큰 필드를 null로 설정합니다.
            disposedValue = true;
        }
    }
    public void Dispose()
    {
        // 이 코드를 변경하지 마세요. 'Dispose(bool disposing)' 메서드에 정리 코드를 입력합니다.
        Dispose(disposing: true);
        GC.SuppressFinalize(this);
    }
}