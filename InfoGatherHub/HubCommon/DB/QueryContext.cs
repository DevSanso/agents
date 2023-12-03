namespace InfoGatherHub.HubCommon.DB;

using System.Collections.Concurrent;
using System.Data.Odbc;
using System.Data.Common;
using System.Threading;


public class QueryContext<T>
{
    private OdbcConnection connection;
    private Action<List<T>, DbDataReader>? action;
    internal int timeout = 0;
    internal string query = "";
    internal string session = "";
    internal bool isTrans = false;
    internal object[]? args = null;

    internal QueryContext(OdbcConnection connection, Action<List<T>, DbDataReader>? memcpyAction)
    {   
        this.connection = connection;
        action = memcpyAction;
    }

    private void SetSession()
    {
        OdbcCommand command = new OdbcCommand();
        command.Connection = connection;
        command.CommandTimeout = 1;

        command.CommandText = $"SET SESSION my_session_name = {session};";
        command.ExecuteNonQuery();
        command.Dispose();
    }

    public void Exec()
    {
        SetSession();

        OdbcCommand command = new OdbcCommand();
        command.Connection = connection;
        command.CommandText = query;
        command.CommandTimeout = timeout;
        
        
        OdbcTransaction? trans = null;
        if(isTrans == true)
        {   
            trans = connection.BeginTransaction();
            command.Transaction = trans;
        }

        if(args != null)
        {
            for(int i = 0 ; i < args!.Length ; i++)
            {
                command.Parameters.Add(new OdbcParameter($"{i+1}" ,args[i]));
            }
        }

        try 
        {
            command.ExecuteNonQuery();
            if(isTrans == true)trans!.Commit();
            command.Dispose();
        }
        catch(Exception)
        {
            command.Dispose();
            if(isTrans == true)trans!.Rollback();
            throw;
        }
    }

    public T[] Query()
    {
        SetSession();

        OdbcCommand command = new OdbcCommand();
        command.Connection = connection;
        command.CommandText = query;
        command.CommandTimeout = timeout;
        
        OdbcTransaction? trans = null;
        if(isTrans == true)
        {   
            trans = connection.BeginTransaction();
            command.Transaction = trans;
        }

        if(args != null)
        {
            for(int i = 0 ; i < args!.Length ; i++)
            {
                command.Parameters.Add(new OdbcParameter($"{i+1}" ,args[i]));
            }
        }

        T[]? ret = null;
        
        try 
        {
            using(var reader = command.ExecuteReader())
            {
                if(isTrans == true)trans!.Commit();
                List<T> output = new List<T>();
                action!(output, reader);
                ret = output.ToArray();
            }
            command.Dispose();
        }
        catch(Exception)
        {
            command.Dispose();
            if(isTrans == true)trans!.Rollback();
            throw;
        }

        return ret!;
    }

}
