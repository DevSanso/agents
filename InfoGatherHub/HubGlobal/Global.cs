namespace InfoGatherHub.HubGlobal;

using InfoGatherHub.HubGlobal.Config;
using InfoGatherHub.HubGlobal.Logger;

public class Global<T, T2> : IConfigLoader<T>, ILogger   
{
    public T2? Extend  {get; protected set;}
    protected Global()
    {
    }

}

public class GlobalProvider<T, T2> : Global<T, T2> where T2 : IGlobalExtend<T>, new()
{
    private GlobalProvider() : base()
    {
    }
    static Global<T, T2>? instance = null;
    public static Global<T, T2> Init()
    {
        if(instance != null)
            throw new Exception("already init global");

        var tmp = new GlobalProvider<T, T2>();
        var e = new T2();
        e.Init(tmp);
        tmp.Extend = e;
        instance = tmp;
        return instance;
    }
    public static Global<T,T2> Global()
    {        
        return instance!;
    }
}


