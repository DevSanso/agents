namespace InfoGatherHub.HubCommon.Compress;

using K4os.Compression.LZ4;
public class Lz4Compress : ICompress
{
    public void Compress(byte[] data, out byte[] output)
    {
        byte []compress = new byte[LZ4Codec.MaximumOutputSize(data.Length)];
        LZ4Codec.Encode(data, 0, data.Length, compress, 0, compress.Length);

        output = compress;
    }
    public void Decompress(byte []data, out byte[] output)
    {
        byte []decompress = new byte[LZ4Codec.MaximumOutputSize(data.Length)];
        LZ4Codec.Decode(data, 0, data.Length, decompress, 0, decompress.Length);

        output = decompress;
    }
}