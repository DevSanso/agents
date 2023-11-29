namespace InfoGatherHub.HubCommon.Compress;

public interface ICompress
{
    void Compress(byte[] data, out byte[] output);
    void Decompress(byte []data, out byte[] output);
}

