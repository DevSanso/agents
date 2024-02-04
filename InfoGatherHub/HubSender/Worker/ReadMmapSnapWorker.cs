namespace InfoGatherHub.HubSender.Worker;

using System.Collections.Concurrent;

using InfoGatherHub.HubSender.Ipc;
using InfoGatherHub.HubCommon.Format;
using InfoGatherHub.HubSender.Worker.Format;
public class ReadMmapSnapWorker : IWorker
{
    private UInt64 currentSeq = 0;
    private ISnapClient client;
    private ConcurrentQueue<IFormat<WorkerFormatHeader>> sender;
    readonly string category;
    internal ReadMmapSnapWorker(ISnapClient client, ConcurrentQueue<IFormat<WorkerFormatHeader>> sender, string category)
    {
        this.client = client;
        this.sender = sender;
        this.category = category;
    }
    private UInt32 ParsingSize(byte []data)
    {
        byte[] seqBin = new byte[8];
        Array.Copy(data,9, seqBin, 0, 8);

        return BitConverter.ToUInt32(seqBin);
    }
    private UInt64 ParsingSeq(byte []data)
    {
        byte[] bin = new byte[8];
        Array.Copy(data,0, bin, 0, 8);

        return BitConverter.ToUInt64(bin);
    }
    private string ParsingCategory(byte []data)
    {
        byte []dataBin = new byte[5];
        Array.Copy(data, 18, dataBin, 0, 5);
        return dataBin.ToString()!.Trim(' ');
    }
    private string ParsingID(byte []data)
    {
        byte []dataBin = new byte[8];
        Array.Copy(data, 24, dataBin, 0, 8);
        return dataBin.ToString()!;
    }
    private byte[] ParsingData(byte []data, int size)
    {
        byte []dataBin = new byte[size];
        Array.Copy(data, 33, dataBin, 0, size);
        return dataBin;
    }
    private void WorkImpl(string category, ISnapClient client, ConcurrentQueue<IFormat<WorkerFormatHeader>> sender)
    {
        client.FetchSnapData();
        byte[]? data = client.GetSnapData();

        UInt64 seq = ParsingSeq(data!);
        
        if(currentSeq == seq) return;

        currentSeq = seq;

        string snapCategory = ParsingCategory(data!);

        if(snapCategory != category) return;
        
        UInt32 size = ParsingSize(data!);

        String id = ParsingID(data!);

        byte[] sendData = ParsingData(data!, (int)size);

        sender.Enqueue(new WorkerFormat(new(id, snapCategory) , sendData));
    }
    public void Work()
    {
        WorkImpl(this.category, this.client, this.sender);
    }
    public void Dispose()
    {
        client.Dispose();
    }
}