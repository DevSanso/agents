namespace InfoGatherHub.HubCommon.Observer;

using System.Threading;

class CacheObserver<T> : IObserver<T> where T :struct
{
    private string name;
    private string Name {get {return name;} }
    private T data;
    public T Data 
    {
        get 
        {
            lock(dataLock)
            {
                return data;
            }
            
        }
    }
    private long id = 1;
    private long Id {get {return Interlocked.Read(ref id);}}

    private readonly object dataLock = new object();

    public CacheObserver(string name)
    {
        this.name = name;
    }

    public void Update(T data)
    {
        lock(dataLock)
        {
            this.data = data;
        }
        Interlocked.Increment(ref id);
    }
}
