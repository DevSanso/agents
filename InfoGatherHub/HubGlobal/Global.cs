namespace InfoGatherHub.HubGlobal;

using InfoGatherHub.HubGlobal.Config;
using InfoGatherHub.HubGlobal.Logger;
using InfoGatherHub.HubGlobal.Extend;

public class Global<T> : IConfigLoader<T>, ILogger
{
}

public class Global<T,E> : IConfigLoader<T>, ILogger
{
}

public class GlobalProvider<T> 
{
    private GlobalProvider()
    {
    }
    
    static Global<T>? instance = null;
    public static Global<T> Init()
    {
        if(instance != null)
            throw new Exception("already init global");

        instance = new();
        return instance;
    }
    public static Global<T> Global()
    {        
        return instance!;
    }
}

public class GlobalProvider<T,E> : IGlobalExtend<E>
{
    private GlobalProvider()
    {
    }
    
    static Global<T,E>? instance = null;
    public static Global<T,E> Init()
    {
        if(instance != null)
            throw new Exception("already init global");

        instance = new();
        return instance;
    }
    public static Global<T,E> Global()
    {        
        return instance!;
    }
}


