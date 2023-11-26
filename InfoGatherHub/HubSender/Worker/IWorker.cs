namespace InfoGatherHub.HubSender.Worker;


interface IWorker : IDisposable
{
    void Work();
}