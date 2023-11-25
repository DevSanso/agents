namespace InfoGatherHub.HubCommon.Format;

public interface IFormat<T>
{
    string ObjectName();

    T Header();

    byte[] Data();
}

public sealed class Void 
{
    internal Void() { }
}