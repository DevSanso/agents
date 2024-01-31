namespace InfoGatherHub.HubServer.Config;

public record DbConfig
{
    public string Ip = "";
    public int Port = 0;
    public string ID = "";
    public string Password = "";
    public string Dbname = "";
    public string Driver = "";
    public int MaxConn = 0;
}