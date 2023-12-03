namespace InfoGatherHub.HubCommon.DB;

using System.Collections.Concurrent;
using System.Data.Odbc;
using System.Data.Common;
using System.Threading;
public class ConnPool
{
    private readonly string dataSource = "";
    private readonly int connTimeout = 0;
    private readonly int connMaxCount = 0;
    private ConcurrentBag<OdbcConnBox> bag = new ConcurrentBag<OdbcConnBox>();

    private bool CheckIsCanUseConn(OdbcConnBox box)
    {
        var state = box.Conn!.State;
        return box.IsUsed == OdbcConnBox.OdbcConnBoxNotUsed && (state == System.Data.ConnectionState.Connecting || state == System.Data.ConnectionState.Open);
    }
    private OdbcConnBox GetConn()
    {
        var findEnum = bag.Where(CheckIsCanUseConn);
        if( findEnum.Count() == 0 && bag.Count >= connMaxCount ) throw new Exception("");

        OdbcConnBox connBox;
        if( findEnum.Count() == 0)
        {
            var emptyBox = bag.Where((OdbcConnBox box) => box.Conn == null).First();
            Interlocked.Exchange(ref emptyBox.IsUsed, OdbcConnBox.OdbcConnBoxUsed);
            emptyBox.Conn = new OdbcConnection(dataSource);
            emptyBox.Conn.Open();
            connBox = emptyBox;
        }
        else
        {
            connBox = findEnum.First();
            Interlocked.Exchange(ref connBox.IsUsed, OdbcConnBox.OdbcConnBoxUsed);
        }
        
        return connBox;
    }
    private void FreeConnBox()
    {
        for(int i = 0; i < connMaxCount; i++)
        {
            var conn = bag.ElementAt(i);
            if (conn.IsUsed == OdbcConnBox.OdbcConnBoxUsed && conn.Conn == null) continue;

            var state = conn.Conn!.State;
            if(state == System.Data.ConnectionState.Closed || state == System.Data.ConnectionState.Broken)
            {
                conn.Conn!.Dispose();
                conn.Conn = null;
            }
        }
    }

    private void PutConnToPool(OdbcConnBox conn)
    {
        if(conn.Conn!.State == System.Data.ConnectionState.Executing || conn.Conn!.State == System.Data.ConnectionState.Fetching)
        {
            throw new Exception("");
        }
        Interlocked.Exchange(ref conn.IsUsed, OdbcConnBox.OdbcConnBoxNotUsed);
        FreeConnBox();
    }

    public void Exec(string session, string query, bool isTrans, int timeoutSecond, object[]? args)
    {
        var conn = GetConn();
        var queryExec = new QueryContext<object>(conn.Conn!, null)
        {
            query = query,
            timeout = timeoutSecond,
            session = session,
            args = args,
            isTrans = isTrans
        };
        queryExec.Exec();
        PutConnToPool(conn);
    }

    public void Query<T>(string session, string query, int timeoutSecond, bool isTrans, object[]? args, Action<List<T>, DbDataReader> memcpyAction, out T[] output) where T : new()
    {
        var conn = GetConn();
        var queryExec = new QueryContext<T>(conn.Conn!, memcpyAction)
        {
            query = query,
            timeout = timeoutSecond,
            session = session,
            args = args,
            isTrans = isTrans
        };
        output = queryExec.Query();
        PutConnToPool(conn);
    }
}