namespace InfoGatherHub.HubGlobal.Config;

using InfoGatherHub.HubGlobal.Config.Extension.Toml;

public interface IConfigLoader<T>
{
}

public static class IConfigLoaderPorps
{
    static private string? IsUsed = null;

    static internal void SetIsUsed<T>(this IConfigLoader<T> loader, string use)
    {
        IsUsed = use;
    }
    static public string? GetIsUsed<T>(this IConfigLoader<T> loader)
    {
        return IsUsed;
    }
}

public static class IConfigLoaderCaller
{
    static public T? GetConfig<T>(this IConfigLoader<T> loader) where T : class, new()
    {
        string? used = loader.GetIsUsed();

        if(used == "tomi") return loader.GetToml<T>();
        else if(used == "temp") return loader.GetTemp<T>();

        throw new NullReferenceException("Global Config is not setting");
    }
}