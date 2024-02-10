namespace InfoGatherHub.HubServer.Config;


public record Config
{
    public DbConfig DBConfig = new();
    public TcpConfig ServerConfig = new();
}