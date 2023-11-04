namespace InfoGatherHub.HubSender.Worker;

using InfoGatherHub.HubSender.Snap;

static public class ReadOsSnapWorker
{
    static public void Work(ISnapClient client)
    {
        client.fetchSnapData();
        byte[] data = client.getSnapData();

        byte[] seqBin = new byte[8];
        Array.Copy(data,6, seqBin, 0, 8);


        
    }
}