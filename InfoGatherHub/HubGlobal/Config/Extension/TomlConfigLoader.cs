namespace InfoGatherHub.HubGlobal.Config.Extension.Toml;

using Tomlyn;

using InfoGatherHub.HubGlobal.Config;


static  public class TomlConfigLoader
{
    static private Object? config;

    static public void LoadToml<T>(this IConfigLoader<T> loader, string data) where T : class, new()
    {
        var loaded = Toml.ToModel<T>(data);
        TomlConfigLoader.config = loaded;
    }
    static public T GetToml<T>(this IConfigLoader<T> loader) where T : class, new()
    {
        return (T) TomlConfigLoader.config;
    }
}