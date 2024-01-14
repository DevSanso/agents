namespace InfoGatherHub.HubSender.Ipc;

public interface ISnapClient : IDisposable
{
    void fetchSnapData();
    byte[]? getSnapData();
}