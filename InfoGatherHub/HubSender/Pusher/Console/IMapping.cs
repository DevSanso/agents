namespace InfoGatherHub.HubSender.Pusher.Console;


public interface IMapping<T>
{
    void Run(string id, T data);
}