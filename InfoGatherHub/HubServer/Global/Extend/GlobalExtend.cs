namespace InfoGatherHub.HubServer.Global.Extend;

using InfoGatherHub.HubServer.Global.Extend.DB;
using InfoGatherHub.HubGlobal;
using InfoGatherHub.HubGlobal.Config;
using ps_config =  InfoGatherHub.HubServer.Config;
using InfoGatherHub.HubGlobal.Extend;

public class GlobalExtend
{
    public PgPool? DbPool; 
    private static string MakeConnStr(ps_config::DbConfig dbCfg) 
    {
        var cfg = dbCfg;
        return $"Host=${cfg.Ip};Username={cfg.ID};Password={cfg.Password};Database={cfg.Dbname}";
    }
    public void Init(Global<ps_config::Config> g)
    {
        ps_config::Config cfg = g.GetConfig()!;
        DbPool = new(cfg.DBConfig.MaxConn, MakeConnStr(cfg.DBConfig));
    }
    
}