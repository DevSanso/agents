namespace InfoGatherHub.HubCommon.Display;

public class DisplayConsole : IDisplay
{
    public void Display(string message)
    {
        Console.WriteLine(message);
    }
}