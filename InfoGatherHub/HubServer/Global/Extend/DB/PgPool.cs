namespace InfoGatherHub.HubServer.Global.Extend.DB;

using Npgsql;

using InfoGatherHub.HubCommon.Pool;
using System.Data.Common;

public class PgPool : Pool<NpgsqlConnection>
{
    private static NpgsqlConnection Gen(string connstr)
    {
        throw new NotImplementedException();
    }
    public PgPool(long size, string connstr) : base(size, ()=>Gen(connstr))
    {
    }
    private object? ExecuteImpl(NpgsqlConnection conn, string query)
    {
        using var cmd = new NpgsqlCommand(query, conn);
        cmd.ExecuteNonQuery();

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
    public void Execute(string query)
    {
        base.Run<string,object?>(ExecuteImpl, query);
    }

    public T Query<T>(string query, Func<DbDataReader?,T> readFunc) where T : class
    {
        DbDataReader? reader = base.Run<string,DbDataReader?>(QueryImpl, query);
        T ret = readFunc(reader);
        Task disposeTask = reader!.DisposeAsync().AsTask();
        disposeTask.Wait();
        disposeTask.Dispose();
        return ret;
    }
}