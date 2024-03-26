namespace InfoGatherHub.HubGlobal.Config.Extension.Toml;

using InfoGatherHub.HubGlobal.Config;

static  public class TempConfigLoader
{
    static private Object? config;

    static public void LoadTemp<T>(this IConfigLoader<T> loader) where T : class, new()
    {
        TempConfigLoader.config = new T();
        loader.SetIsUsed("temp");
    }
    static internal T? GetTemp<T>(this IConfigLoader<T> loader) where T : class, new()
    {
        return  (T?) config;
    }
}