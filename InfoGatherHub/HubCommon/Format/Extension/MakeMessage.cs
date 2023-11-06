namespace InfoGatherHub.HubCommon.Format.Extension;

using System.Text.Json;
using System.Text;

public static class IFormatMakeMessage
{

    private class JsonFormat
    {
        public long UnixTime {get;set;}
        public int UTC {get;set;}
        public byte[] Data {get;set;}
    }
    public static byte[] MakeBasicMessage(this IFormat<Void> format)
    {
        JsonFormat msg = new JsonFormat()
        {
            UnixTime = format.UnixTime(),
            UTC = format.UTC(),
            Data = format.Data()
        };
        return JsonSerializer.Serialize(msg);
    }
}