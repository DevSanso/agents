namespace InfoGatherHub.HubCommon.DB;

using System.Collections.Concurrent;
using System.Data.Odbc;
using System.Data.Common;
using System.Threading;

internal class OdbcConnBox
{
    internal OdbcConnBox()
    {
        
    }
    internal OdbcConnection? Conn;
    internal int IsUsed = 0;
    public const int OdbcConnBoxUsed = 1;
    public const int OdbcConnBoxNotUsed = 0;
}