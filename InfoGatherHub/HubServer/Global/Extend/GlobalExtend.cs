namespace InfoGatherHub.HubServer.Global.Extend;

using InfoGatherHub.HubServer.Global.Extend.DB;
using InfoGatherHub.HubGlobal;
using InfoGatherHub.HubGlobal.Config;
using InfoGatherHub.HubServer.Config;

public class GlobalExtend : IGlobalExtend<Config>
{
    public PgPool? DbPool; 
    private static string MakeConnStr(DbConfig dbCfg) 
    {
        var cfg = dbCfg;
        return $"Host=${cfg.Ip};Username={cfg.ID};Password={cfg.Password};Database={cfg.Dbname}";
    }
    public void Init<T2>(in Global<Config, T2> g) where T2 : IGlobalExtend<Config>, new()
    {
        Config cfg = g.GetConfig()!;
        DbPool = new(cfg.DBConfig.MaxConn, MakeConnStr(cfg.DBConfig));
    }
    
}