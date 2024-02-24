namespace InfoGatherHub.HubSender.Pusher;

public interface IPusher<T> : IDisposable
{
    void Push(string key, T data);
}
