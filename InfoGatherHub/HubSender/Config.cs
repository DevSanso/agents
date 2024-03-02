namespace InfoGatherHub.HubSender;

using System.Collections.Generic;
record MMapSnapIpcConfig
{
    public string path = "";
    public int size = 0;
}

record TcpSnapIpcConfig
{
    public bool isUsed = false;
    public int port = 0;
}

record LogSetting
{
    public string type = "";
    public string? logPath = null;
}

record PusherSetting
{
    public string type = "";
    public string ip = "";
    public int port = 0;
}

record Config
{
    public LogSetting LogSetting = new LogSetting();
    public PusherSetting PusherSetting = new PusherSetting();
    public Dictionary<String, MMapSnapIpcConfig> MmapSnapSetting = new Dictionary<string, MMapSnapIpcConfig>();
    //public TcpSnapIpcConfig? tcpSnapServerSetting = new TcpSnapIpcConfig();
}