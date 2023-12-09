namespace InfoGatherHub.HubServer;

record OdbcConfig
{
    public string Ip = "";
    public int Port = 0;
    public string ID = "";
    public string Password = "";
    public string Dbname = "";
    public string Driver = "";
    public int MaxConn = 0;
}
record Config
{
    public OdbcConfig odbcConfig = new OdbcConfig();
}