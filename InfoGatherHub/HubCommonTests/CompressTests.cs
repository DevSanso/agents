namespace InfoGatherHub.HubCommon.Tests;

using Microsoft.VisualStudio.TestTools.UnitTesting;
using InfoGatherHub.HubCommon.Compress;

[TestClass]
public class CompressTests
{
    [TestMethod]
    public void CompressTest()
    {
        var msg = "Hello World!";
        var output = new byte[256];
        var decompress = new byte[256];
        var compress = new Lz4Compress();
        compress.Compress(System.Text.Encoding.UTF8.GetBytes(msg), out output);
        compress.Decompress(output, out decompress);

        Assert.AreEqual(msg, System.Text.Encoding.UTF8.GetString(decompress).TrimEnd('\0'));
    }
}