namespace InfoGatherHub.HubSender;

record SnapInfo
{
    string type = "";
    string path = "";
}
record Config
{
    string logPath = "";
    SnapInfo? osSnapInfo;
}