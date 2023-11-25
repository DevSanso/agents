namespace InfoGatherHub.HubCommon.Format;

public class Format : IFormat<Void>
{
    private string objectName;
    private byte[] data;
    public Format(string objectName, byte[] data)
    {
        this.objectName = objectName;
        this.data = data;
    }
    public byte[] Data()
    {
        return data;
    }
    public Void Header()
    {
        return new Void();
    }
    public string ObjectName()
    {
        return objectName;
    }
}