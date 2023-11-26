namespace InfoGatherHub.HubGlobal.Logger;

using InfoGatherHub.HubGlobal.Logger.Extension.Xml;

public enum LogLevel
{
    Debug,
    Error
}
public enum LogCategory
{
    ALL,
    Network,
    IO,
    Code
}
public interface ILogger {}

public static class LoggerCaller
{
    public static void Log(this ILogger logger, LogLevel level,  LogCategory category, String message)
    {
        if(XmlLogger.IsInit == true)logger.LogXml(level, category, message);
    }
}