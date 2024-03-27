namespace InfoGatherHub.HubSender.Ipc;

using System;
using System.IO;
using System.IO.MemoryMappedFiles;

public class MemMapClient : ISnapClient
{
    private readonly string path = "";
    private readonly int size = 0;
    private byte[]? buffer;
    private readonly object lockObj = new();
    public MemMapClient(string pathname, int size)
    {  
        this.path = pathname;
        this.size = size;
        buffer = new byte[size];
    }
    public void FetchSnapData()
    {
        using var f = File.Open(this.path, FileMode.Open, FileAccess.Read, FileShare.ReadWrite);
        using var file = MemoryMappedFile.CreateFromFile(f, null, this.size, MemoryMappedFileAccess.Read, HandleInheritability.None, false);

        using var accessor = file!.CreateViewAccessor(0, this.size, MemoryMappedFileAccess.Read);

        lock(lockObj)
            accessor!.ReadArray(0, buffer!, 0, this.size);     
        
    }
    public byte[]? GetSnapData()
    {
        byte []ret = new byte[size];
        
        lock(lockObj)
            buffer!.CopyTo(ret, 0);

        return ret;
    }
    public void Dispose()
    {
        this.buffer = null;
    }
}