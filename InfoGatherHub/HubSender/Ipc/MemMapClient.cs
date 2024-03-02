namespace InfoGatherHub.HubSender.Ipc;

using System;
using System.IO;
using System.IO.MemoryMappedFiles;
using K4os.Compression.LZ4.Internal;

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
        using var file = MemoryMappedFile.CreateFromFile(File.Open(this.path, FileMode.Open, FileAccess.Read), 
            null, 0, MemoryMappedFileAccess.Read, HandleInheritability.None, false);

        using var accessor = file!.CreateViewStream(0, this.size, MemoryMappedFileAccess.CopyOnWrite);

        lock(lockObj)
            accessor!.Read(buffer!, 0, size);
        
    }
    public byte[]? GetSnapData()
    {
        byte []ret = new byte[size];
        
        lock(lockObj)
            ret.CopyTo(this.buffer!, 0);
        

        return ret;
    }
    public void Dispose()
    {
        this.buffer = null;
    }
}