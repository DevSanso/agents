namespace InfoGatherHub.HubSender.Ipc;

using System;
using System.IO;
using System.IO.MemoryMappedFiles;
public class MemMapClient : ISnapClient
{
    private MemoryMappedFile? snap;
    private int size = 0;
    private byte[] buffer;
    public MemMapClient(string pathname, int size)
    {  
        snap = MemoryMappedFile.CreateFromFile(File.Open(pathname, FileMode.Open, FileAccess.Read), 
            null, 0, MemoryMappedFileAccess.Read, HandleInheritability.None, false);
        this.size = size;
        buffer = new byte[size];
    }
    public void FetchSnapData()
    {
        using(var accessor = snap!.CreateViewStream(0, this.size, MemoryMappedFileAccess.CopyOnWrite))
        {
            accessor!.Read(buffer,0, size);
        }
    }
    public byte[]? GetSnapData()
    {
        byte []ret = new byte[size];
        ret.CopyTo(this.buffer, 0);

        return ret;
    }
    public void Dispose()
    {
        snap?.Dispose();
    }
}