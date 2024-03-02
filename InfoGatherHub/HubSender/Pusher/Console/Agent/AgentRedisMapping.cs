namespace InfoGatherHub.HubSender.Pusher.Console.Agent;

using System.Reflection;
using System;

using InfoGatherHub.HubSender.Pusher.Console;
using InfoGatherHub.HubProtos.Agent.Redis;


public class AgentRedisMapping : IMapping<AgentRedisSnap>
{
    public void Run(string id, AgentRedisSnap snaps)
    {
        foreach(var snap in snaps.Datas)
        {
            switch(snap.Format)
            {
                case DataFormat.ClientLists:
                Print(RedisClientList.Parser.ParseFrom(snap.RawData).Clients);
                break;
                case DataFormat.Dbsize:
                Print(DbSize.Parser.ParseFrom(snap.RawData));
                break;
                case DataFormat.InfoCpu:
                Print(RedisCpuInfo.Parser.ParseFrom(snap.RawData));
                break;
                case DataFormat.InfoMemory:
                Print(RedisMemoryInfo.Parser.ParseFrom(snap.RawData));
                break;
                case DataFormat.InfoStat:
                Print(RedisStatsInfo.Parser.ParseFrom(snap.RawData));
                break;
            }
        }
    }

    private void Print(object obj)
    {
        var typeDesc = new TypeDelegator(obj.GetType());
        
        if(typeDesc.IsArray)
        {
            var array = (Array)obj;
            foreach(var item in array)
            {
                foreach(PropertyInfo prop in obj.GetType().GetProperties())
                {
                    var value = prop.GetValue(obj);
                    if(value != null)
                    {
                        Console.WriteLine($"{prop.Name}: {value}");
                    }
                }
            }
        }
        else
        {
            foreach(PropertyInfo prop in obj.GetType().GetProperties())
            {
                var value = prop.GetValue(obj);
                if(value != null)
                {
                    Console.WriteLine($"{prop.Name}: {value}");
                }
            }
        }
    }
}
