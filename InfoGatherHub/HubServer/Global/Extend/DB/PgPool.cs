namespace InfoGatherHub.HubServer.Global.Extend.DB;

using Npgsql;

using InfoGatherHub.HubCommon.Pool;
using InfoGatherHub.HubServer.Config;
using InfoGatherHub.HubGlobal;
using InfoGatherHub.HubGlobal.Config;
using System.Data;
using System.Data.Common;

internal class PgPool
{
    private static Lazy<DbConfig> config = new(() => 
    {
        var g = GlobalProvider<Config,GlobalExtend>.Global();
        return g.GetConfig()!.DBConfig;
    });
    private static string MakeConnStr() 
    {
        var cfg = config.Value;
        return $"Host=${cfg.Ip};Username={cfg.ID};Password={cfg.Password};Database={cfg.Dbname}";
    }
    private static NpgsqlConnection Gen()
    {
        throw new NotImplementedException();
    }
    Pool<NpgsqlConnection> p = new(50, Gen);
    private object? ExecuteImpl(NpgsqlConnection conn, string query)
    {
        using(var cmd = new NpgsqlCommand(query, conn))
        {
            cmd.ExecuteNonQuery();
        }

        return null;
    }
    private DbDataReader? QueryImpl(NpgsqlConnection conn, string query)
    {
        DbDataReader? reader = null;
        using(var cmd = new NpgsqlCommand(query, conn))
        {
            reader = cmd.ExecuteReader();
        }
        return reader;
    }
    internal PgPool(DbConfig config)
    {

    }
    public void Execute(string query)
    {
        p.Run<string,object?>(ExecuteImpl, query);
    }

    public T Query<T>(string query, Func<DbDataReader?,T> readFunc) where T : class
    {
        DbDataReader? reader = p.Run<string,DbDataReader?>(QueryImpl, query);
        T ret = readFunc(reader);
        reader!.DisposeAsync();
        return ret;
    }

}