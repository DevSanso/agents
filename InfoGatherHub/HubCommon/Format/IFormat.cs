namespace InfoGatherHub.HubCommon.Format;

public interface IFormat<T>
{
    int UTC();
    long UnixTime();

    string ObjectName();

    T Header();

    byte[] Data();
}

public sealed class Void 
{
    internal Void() { }
}