namespace InfoGatherHub.HubSender.Worker;

using System.Collections.Concurrent;

using InfoGatherHub.HubSender.Snap;
using InfoGatherHub.HubCommon.Format;
public class ReadOsSnapWorker
{
    private UInt64 currentSeq = 0;
    private UInt64 ParsingSeq(byte []data)
    {
        byte[] seqBin = new byte[8];
        Array.Copy(data,6, seqBin, 0, 8);

        return BitConverter.ToUInt64(seqBin);
    }
    private UInt32 ParsingSize(byte []data)
    {
        byte[] bin = new byte[8];
        Array.Copy(data,0, bin, 0, 4);

        return BitConverter.ToUInt32(bin);
    }
    private byte[] ParsingData(byte []data, int size)
    {
        byte []dataBin = new byte[size];
        Array.Copy(data, 14, dataBin, 0, size);

        return dataBin;
    }
    public void Work(ISnapClient client, ConcurrentQueue<IFormat<Void>> sender)
    {
        client.fetchSnapData();
        byte[] data = client.getSnapData();

        UInt64 seq = ParsingSeq(data);
        
        if(currentSeq == seq) return;

        currentSeq = seq;

        UInt32 size = ParsingSize(data);

        byte[] sendData = ParsingData(data, (int)size);

        sender.Enqueue(Format.Now("OS", sendData));
    }
}