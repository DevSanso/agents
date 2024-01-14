namespace InfoGatherHub.HubGlobal.Logger;

using InfoGatherHub.HubGlobal.Logger.Extension.Line;

public enum LogLevel
{
    Error = 0,
    Info = 1,
    Debug = 2
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
        if(LineLogger.IsInit == true)logger.LogLine(level, category, message);
    }
}