namespace InfoGatherHub.HubSender.Pusher;

using System;

public class ConsolePusher : IPusher<byte[]>
{
    private Console.AgentMapping mapping = new Console.AgentMapping();
    public void Dispose()
    {
    }

    public void Push(string key, byte[] data)
    {
        mapping.Run(key, data);
    }
}