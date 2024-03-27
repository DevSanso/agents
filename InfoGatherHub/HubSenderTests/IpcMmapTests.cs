namespace InfoGatherHub.HubSenderTests;

using System.Drawing;
using  InfoGatherHub.HubSender.Ipc;

using System.IO;
using System.IO.MemoryMappedFiles;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using NetTopologySuite.Operation.Buffer;
using System.Threading;

[TestClass]
public class IpcMmapTests
{
    static private readonly int Size = 32;
    static public string GetTempPath()
    {
        return "/tmp/hubsender_ipc_test.tmp";
    }
    static public void readFile(string text, MemoryMappedFile file)
    {
        using var acc = file.CreateViewAccessor(0, Size, MemoryMappedFileAccess.Write);
        var b = System.Text.Encoding.UTF8.GetBytes(text);
        acc.WriteArray(0, b, 0, text.Length);
        acc.Flush();
        
    }
    [TestMethod]
    public void TestIpcMmapTest()
    {   
        using var f = File.Open(GetTempPath(),FileMode.Create, FileAccess.ReadWrite, FileShare.ReadWrite);
        f.SetLength(Size);
        using var file = MemoryMappedFile.CreateFromFile(f, 
            null, f.Length, MemoryMappedFileAccess.ReadWrite, HandleInheritability.None, false);
        string text = "Hello World!".PadRight(Size,'0');
        readFile(text, file);

        ISnapClient client = new MemMapClient(GetTempPath(), Size);
        
        client.FetchSnapData();
        var data = client.GetSnapData();

        Assert.AreEqual(text, System.Text.Encoding.UTF8.GetString(data!));
        
    }
}