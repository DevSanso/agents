namespace InfoGatherHub.HubServer.Global.Extend;

using System.Data.Odbc;
using System.Data;
using System.Collections.Concurrent;
using System.Data.Common;

using InfoGatherHub.HubGlobal;
using InfoGatherHub.HubGlobal.Config;
using InfoGatherHub.HubServer;

using G = InfoGatherHub.HubGlobal.GlobalProvider<Config,GlobalExtend>;

internal class OdbcPool
{
    private record OdbcBox
    {
        internal OdbcConnection? conn;
        internal bool IsUsed;
    }    
    Lazy<string> ip = new Lazy<string>(()=> G.Global().GetConfig()!.odbcConfig.Ip);
    Lazy<int> port = new Lazy<int>(()=> G.Global().GetConfig()!.odbcConfig.Port);
    Lazy<string> id = new Lazy<string>(() => G.Global().GetConfig()!.odbcConfig.ID);
    Lazy<string> password = new Lazy<string>(() => G.Global().GetConfig()!.odbcConfig.Password);
    Lazy<string> dbname = new Lazy<string>(() => G.Global().GetConfig()!.odbcConfig.Dbname);
    Lazy<string> driver = new Lazy<string>(() => G.Global().GetConfig()!.odbcConfig.Driver);

    object connListLock = new object();
    List<OdbcBox> connList = new List<OdbcBox>(50);
    private OdbcBox NewConnection()
    {
        OdbcBox b = new OdbcBox();
        OdbcConnection conn = new OdbcConnection(@$"Driver={driver.Value};
            Server={ip.Value};Port={port.Value};UID={id.Value};PWD={password.Value};DATABASE={dbname.Value}");

        conn.ConnectionTimeout = 60;
        conn.Open();
        b.conn = conn;
        return b;
    }

    private void OdbcBoxNotUsed(OdbcBox box)
    {
        box.IsUsed = false;
    }
    private OdbcBox GetConnFromPool()
    {
        OdbcBox? b = null;
        lock(connListLock)
        {
            b = connList.Where((b) => 
            {
                if(b.IsUsed == false  && (b.conn!.State == ConnectionState.Open || b.conn!.State == ConnectionState.Connecting))
                    return true;
                
                return false;
            })
            .First();
            if(b != null)b.IsUsed = true;
        }

        if(b == null)
        {
            lock(connListLock)
            {
                b = NewConnection();
                b.IsUsed = true;
                connList.Add(b);
            }    
        }

        return b;
    }
    private void FreeConnsToPool()
    {
        lock(connListLock)
        {
            int count = connList.Count;
            for(int i=0;i < count ; i++)
            {
                OdbcConnection conn = connList[i].conn!;
                if(conn.State == ConnectionState.Broken || conn.State == ConnectionState.Closed)
                {
                    conn.Dispose();
                    connList.RemoveAt(i);
                }
            }
        }
    }

    private void SetSession(OdbcConnection conn, string name)
    {
        using(OdbcCommand cmd = new OdbcCommand($"set application_name to {name}", conn))
        {
            cmd.ExecuteNonQuery();
        }    
    }        

    public void RunConnection(string session, string query, DbParameter[] parameters,  Action<DbDataReader> action)
    {
        OdbcBox conn = GetConnFromPool();
        SetSession(conn.conn!, session);
        
        using(OdbcCommand cmd = new OdbcCommand(query, conn.conn))
        {
            cmd.Parameters.AddRange(parameters);
            DbDataReader? reader = cmd.ExecuteReader();
            action(reader);
        }
        OdbcBoxNotUsed(conn);
        FreeConnsToPool();
    }
    public void RunConnection(string session, string execute,  DbParameter[] parameters, Action<int> action)
    {
       OdbcBox conn = GetConnFromPool();
       SetSession(conn.conn!, session);
        
        using(OdbcCommand cmd = new OdbcCommand(execute, conn.conn))
        {
            cmd.Parameters.AddRange(parameters);
            int rows = cmd.ExecuteNonQuery();
            action(rows);
        }
        OdbcBoxNotUsed(conn);
        FreeConnsToPool();
    }
}