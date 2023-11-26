namespace InfoGatherHub.HubSender;

record SnapInfoSetting
{
    public string path = "";
    public int size = 0;
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
    public SnapInfoSetting osSnapSetting = new SnapInfoSetting();
    public HubServerSetting hubServerSetting = new HubServerSetting();
}