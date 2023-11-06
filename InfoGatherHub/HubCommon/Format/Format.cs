namespace InfoGatherHub.HubCommon.Format;

public class Format : IFormat<Void>
{
    private long second;
    private int utc;
    private string objectName;
    private byte[] data;
    private Format(long sec, int utc,string objectName, byte[] data)
    {
        second = sec;
        this.objectName = objectName;
        this.utc = utc;
        this.data = data;
    }
    private readonly DateTime unix = new DateTime(1970, 1, 1, 0, 0, 0, DateTimeKind.Local);

    static public IFormat<Void> Now(string objectName, byte[] data)
    {
        DateTime local = DateTime.Now;

        long sec = (long)(local - new DateTime(1970, 1, 1, 0, 0, 0, DateTimeKind.Local)).TotalSeconds;
        int utc = local.ToLocalTime().Hour - local.ToUniversalTime().Hour;

        return new Format(sec, utc, objectName, data);
    }
    public int UTC()
    {
        return utc;
    }
    public long UnixTime()
    {
        return second;
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