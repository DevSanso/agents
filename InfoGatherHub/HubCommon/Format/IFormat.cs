namespace InfoGatherHub.HubCommon.Format;

public interface IFormat<T>
{
    T Header();

    byte[] Data();
}
