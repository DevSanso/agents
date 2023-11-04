namespace InfoGatherHub.HubCommon.Display;

using System.IO;
using System.Text;
public class DisplayFile : IDisplay, IDisposable
{
    private FileStream file;
    public DisplayFile(string path)
    {
        file = File.Open(path, FileMode.CreateNew, FileAccess.Write);
    }

    public void Dispose()
    {
        file.Close();
    }

    public void Display(string message)
    {
        byte []output = Encoding.UTF8.GetBytes(message + Environment.NewLine);
        file.Write(output, 0, output.Length);
    }
}