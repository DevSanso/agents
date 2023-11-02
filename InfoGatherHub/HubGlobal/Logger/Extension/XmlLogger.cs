namespace InfoGatherHub.HubGlobal.Logger.Extension.Xml;

using System.Xml;
using System.Threading.Channels;

using InfoGatherHub.HubGlobal.Logger;
using InfoGatherHub.HubCommon.Cache;

public static class XmlLogger
{
    private record XmlLogData( LogLevel level,  LogCategory category, string message);
    private static  string? logPath = null;
    private static Channel<XmlLogData> channel = Channel.CreateUnbounded<XmlLogData>();
    private static LogLevel[] logLevels = new LogLevel[2];
    private static EnumToStringCache<LogLevel> levelCache = new EnumToStringCache<LogLevel>();
    private static EnumToStringCache<LogCategory> categoryCache = new EnumToStringCache<LogCategory>();

    public static void InitXml(this ILogger logger, string logPath)
    {
        XmlLogger.logPath = logPath;
    }
    public static void LogThread()
    {
        ChannelReader<XmlLogData> reader = channel.Reader;
        XmlDocument xml = new XmlDocument();

        XmlLogData? logData = new XmlLogData(LogLevel.Debug, LogCategory.ALL, "");
        while(true)
        {
            if(reader.TryRead(out logData) == false) continue;
            

            string current = DateTime.Now.ToString();
            var logtime = xml.CreateElement("Logtime");
            logtime.InnerText = current;

            string levelString = levelCache.get(logData.level);
            var level = xml.CreateElement("Level");
            level.InnerText = levelString;

            string categoryString = categoryCache.get(logData.category);
            var category = xml.CreateElement("Category");
            category.InnerText = categoryString;

            var message = xml.CreateElement("Message");
            message.InnerText = logData.message;

            var root = xml.CreateElement("HubSender");
            root.AppendChild(logtime);
            root.AppendChild(level);
            root.AppendChild(category);
            root.AppendChild(message);

            string output = xml.OuterXml;

            Console.WriteLine(output);
        }
    }

    public static void LogXml(this ILogger logger,  LogLevel level,  LogCategory category, String message)
    {
        channel.Writer.WriteAsync(new XmlLogData(level, category, message));
    }

}


