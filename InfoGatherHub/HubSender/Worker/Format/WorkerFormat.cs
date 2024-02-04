namespace InfoGatherHub.HubSender.Worker.Format;

using InfoGatherHub.HubCommon.Format;

public record WorkerFormatHeader 
{
    public readonly string Id = ""; 
    public readonly string ObjectName = "";

    public WorkerFormatHeader(string id, string objectName)
    {
        this.Id = id;
        this.ObjectName = objectName;
    }
}
public class WorkerFormat : IFormat<WorkerFormatHeader>
{
    readonly WorkerFormatHeader header;
    readonly byte[] data;
    public WorkerFormat(WorkerFormatHeader header, byte[] data)
    {
        this.header = header;
        this.data = data;
    }
    public byte[] Data()
    {
        return data;
    }

    public WorkerFormatHeader Header()
    {
        return header;
    }
}