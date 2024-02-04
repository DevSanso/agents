namespace InfoGatherHub.HubSender;

record MMapSnapIpcConfig
{
    public bool isUsed = false;
    public string path = "";
    public int size = 0;
    public string physicsIp = "";
    public string virtualIp = "";
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

record HubServerSetting
{
    public string ip = "";
    public int port = 0;
}

record Config
{
    public LogSetting logSetting = new LogSetting();
    public MMapSnapIpcConfig? osSnapSetting = new MMapSnapIpcConfig();
    public MMapSnapIpcConfig? redisSnapSetting = new MMapSnapIpcConfig();
    //public TcpSnapIpcConfig? tcpSnapServerSetting = new TcpSnapIpcConfig();
    public HubServerSetting hubServerSetting = new HubServerSetting();
}