namespace InfoGatherHub.HubServer.Mapping.Agent;

using Google.Protobuf;

using InfoGatherHub.HubServer.Mapping;
using InfoGatherHub.HubProtos.Agent.Os;
using InfoGatherHub.HubProtos.Agent.Os.Net;

public class AgentOsMapping : IMapping<AgentOsSnap>
{
    public void Run(AgentOsSnap snaps)
    {
        foreach(var snap in snaps.Datas)
        {
            switch(snap.Format)
            {
                case DataFormat.NetArp:
                ArpInfos.Parser.ParseFrom(snap.RawData);
                break;
                case DataFormat.NetDev:
                NetDevInfos.Parser.ParseFrom(snap.RawData);
                break;
                case DataFormat.NetSockStat:
                SockStatInfo.Parser.ParseFrom(snap.RawData);
                break;
                case DataFormat.NetTcp4Stat:
                Tcp4Stats.Parser.ParseFrom(snap.RawData);
                break;
            }
        }
    }
}
