namespace InfoGatherHub.HubSender.Snap;

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
        snap = MemoryMappedFile.CreateFromFile(pathname,FileMode.Open);
        this.size = size;
        buffer = new byte[size];
    }
    public void fetchSnapData()
    {
        using(var accessor = snap?.CreateViewAccessor(0, size,MemoryMappedFileAccess.Read))
        {
            accessor?.ReadArray<byte>(0, buffer, 0, size);
        }
    }
    public byte[] getSnapData()
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