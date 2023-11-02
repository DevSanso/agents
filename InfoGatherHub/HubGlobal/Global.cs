namespace InfoGatherHub.HubGlobal;

using InfoGatherHub.HubGlobal.Config;
using InfoGatherHub.HubGlobal.Logger;

public interface Global<T> : IConfigLoader<T>, ILogger   
{
}

public class GlobalProvider<T> : Global<T> 
{
    private GlobalProvider()
    {
    }
    static Global<T>? instance = null;
    public static Global<T> Global()
    {
        if(instance == null)
            instance = new GlobalProvider<T>();
        
        return instance;
    }
}


