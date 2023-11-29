namespace InfoGatherHub.HubServer.Mapping;

using Google.Protobuf;

using InfoGatherHub.HubProtos.Agent;
using InfoGatherHub.HubProtos.Agent.Os;
using InfoGatherHub.HubServer.Mapping.Agent;

public class AgentMapping : IMapping<SnapData>
{
    AgentOsMapping osMapping = new AgentOsMapping();
    public void Run(SnapData snap)
    {
        switch(snap.Format)
        {
            case SnapFormat.Os:
                osMapping.Run(AgentOsSnap.Parser.ParseFrom(snap.RawSnap));
            break;
        }
    }
}