namespace InfoGatherHub.HubGlobal.Logger.Extension.Line;

using System.Threading.Channels;
using System.Threading;

using InfoGatherHub.HubGlobal.Logger;
using InfoGatherHub.HubCommon.Cache;
using InfoGatherHub.HubCommon.Display;

public static class LineLogger
{
    private record LineLogData(LogLevel level,  LogCategory category, string message);
    public static bool IsInit {get;private set;}
    private static Channel<LineLogData> channel = Channel.CreateUnbounded<LineLogData>();
    private static int logLevel = (int)LogLevel.Error;
    private static EnumToStringCache<LogLevel> levelCache = new EnumToStringCache<LogLevel>();
    private static EnumToStringCache<LogCategory> categoryCache = new EnumToStringCache<LogCategory>();
    private static IDisplay? display = null;
 
    public static void InitLogLine(this ILogger logger, IDisplay display)
    {
        if(IsInit == true)
        {
            throw new Exception();
        }
        LineLogger.display = display;

        new Thread(new ThreadStart(LogThread))
        {
            IsBackground = true
        };
        IsInit = true;
    }
    private static void LogThread()
    {
        ChannelReader<LineLogData> reader = channel.Reader;
        LineLogData? logData = null;
        while(true)
        {
            if(reader.TryRead(out logData) == false) 
            {
                Thread.Sleep(10);
                continue;
            }

            string current = DateTime.Now.ToString();
            string level = levelCache.Get(logData.level);
            string category = categoryCache.Get(logData.category);

            string output = $"[{current}]-[{level}]-[{category}]:{logData.message}";
            display?.Display(output);
            logData = null;
        }
    }
    internal static void LogLine(this ILogger logger,  LogLevel level,  LogCategory category, String message)
    {
        if((int)level <= logLevel) channel.Writer.WriteAsync(new LineLogData(level, category, message));
    }

}