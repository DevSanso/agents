namespace InfoGatherHub.HubSender.Snap;

public interface ISnapClient : IDisposable
{
    void fetchSnapData();
    byte[] getSnapData();
}