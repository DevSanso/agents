namespace InfoGatherHub.HubCommon.Observer;
interface IObserver<T>
{
    void Update(T data);
}