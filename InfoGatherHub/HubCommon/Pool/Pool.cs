namespace InfoGatherHub.HubCommon.Pool;

using System.Collections.Concurrent;


public class Pool<T> where T : class, IDisposable
{
    private readonly long maxSize;
    private long allocSize;
    private readonly Func<T> genFunc;
    private readonly ConcurrentQueue<T> q = new();

    public Pool(long maxSize, Func<T> genFunc)
    {
        this.maxSize = maxSize;
        this.genFunc = genFunc;
    }
    private bool CheckSize() => Interlocked.Read(ref allocSize) < maxSize;
    
    private T GetObject()
    {
        if(!CheckSize()) 
            throw new IndexOutOfRangeException("Pool Alloc Size is Max");

        bool deqRet = q.TryDequeue(out T? obj);

        obj ??= genFunc();
        Interlocked.Increment(ref allocSize);
        return obj;
    }
    public R Run<T2,R>(Func<T,T2,R> func, T2 args) where R : class?
    {
        R ret;
        T? obj = null;
        try
        {   
            obj = GetObject();
            ret = func(obj, args);
        }
        catch (System.Exception)
        {
            obj?.Dispose();
            Interlocked.Decrement(ref allocSize);
            throw;
        }
        q.Enqueue(obj);
        return ret;
    }
}