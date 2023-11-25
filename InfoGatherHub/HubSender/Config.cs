namespace InfoGatherHub.HubSender;

record SnapInfoSetting
{
    public string path = "";
}
record LogSetting
{
    public string type = "";
    public string? logPath = null;
}
record Config
{
    public LogSetting logSetting;
    public SnapInfoSetting osSnapSetting;
}