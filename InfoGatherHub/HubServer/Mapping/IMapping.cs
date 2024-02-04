namespace InfoGatherHub.HubServer.Mapping;


public interface IMapping<T>
{
    void Run(string id, T data);
}