namespace InfoGatherHub.HubServer.Mapping;


public interface IMapping<T>
{
    void Run(T data);
}