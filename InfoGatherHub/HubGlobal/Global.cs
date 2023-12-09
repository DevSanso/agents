namespace InfoGatherHub.HubGlobal;

using InfoGatherHub.HubGlobal.Config;
using InfoGatherHub.HubGlobal.Logger;

public class Global<T, T2> : IConfigLoader<T>, ILogger   
{
    public T2? Extend  {get; private set;}
    protected Global(T2? extend)
    {
        this.Extend = extend;
    }
}

public class GlobalProvider<T, T2> : Global<T, T2> where T2 : class
{
    private GlobalProvider(T2? extend) : base(extend)
    {
    }
    static Global<T, T2>? instance = null;
    public static Global<T, T2> Init(T2 ?extend)
    {
        if(instance != null)
            throw new Exception("already init global");

        instance = new GlobalProvider<T,T2>(extend);
        return instance;
    }
    public static Global<T,T2> Global()
    {        
        return instance!;
    }
}


