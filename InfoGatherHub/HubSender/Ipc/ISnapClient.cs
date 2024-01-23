namespace InfoGatherHub.HubSender.Ipc;

public interface ISnapClient : IDisposable
{
    void FetchSnapData();
    byte[]? GetSnapData();
}