namespace InfoGatherHub.HubGlobal.Logger.Extension.Line;

using System;
using System.Threading.Channels;
using System.Threading;

using InfoGatherHub.HubGlobal.Logger;
using InfoGatherHub.HubCommon.Cache;
using InfoGatherHub.HubCommon.Display;

public static class LineLogger
{
    private record LineLogData(LogLevel Level,  LogCategory Category, string Message);
    public static bool IsInit {get;private set;}
    private static readonly Channel<LineLogData> Channel = System.Threading.Channels.Channel.CreateUnbounded<LineLogData>();
    private static int logLevel = 0;
    private static readonly EnumToStringCache<LogLevel> levelCache = new();
    private static readonly EnumToStringCache<LogCategory> categoryCache = new();
    private static IDisplay? display = null;
    private static Thread? thread = null;
 
    public static void InitLogLine(this ILogger _, IDisplay display, LogLevel level = LogLevel.Error)
    {
        if(IsInit == true)
        {
            throw new Exception();
        }
        LineLogger.display = display;
        LineLogger.logLevel = (int)level;

        thread = new Thread(new ThreadStart(LogThread))
        {
            IsBackground = true
        };
        IsInit = true;
        thread.Start();
    }
    private static void LogThread()
    {
        ChannelReader<LineLogData> reader = Channel.Reader;
        while (true)
        {
            if (reader.TryRead(out LineLogData? logData) == false)
            {
                Thread.Sleep(10);
                continue;
            }

            string current = DateTime.Now.ToString();
            string level = levelCache.Get(logData.Level);
            string category = categoryCache.Get(logData.Category);

            string output = $"[{current}]-[{level}]-[{category}]:{logData.Message}";
            display?.Display(output);
        }
    }
    
    internal static void LogLine(this ILogger _,  LogLevel level,  LogCategory category, String message)
    {
        if((int)level <= logLevel)
        {
             var valueT = Channel.Writer.WriteAsync(new LineLogData(level, category, message));
             valueT.AsTask().Wait();
        }
    }

}