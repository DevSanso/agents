namespace InfoGatherHub.HubSender.Pusher;

using System.Collections.Generic;

using NRedisStack;
using NRedisStack.RedisStackCommands;
using StackExchange.Redis;

public class RedisPusher : IPusher<byte[]>
{
    private readonly ConnectionMultiplexer redis;
    private readonly IDatabase db;
    private bool disposed = false;
    private Dictionary<string, RedisChannel> channels = new();
    public RedisPusher(string ip, int port)
    {
        redis = ConnectionMultiplexer.Connect($"{ip}:{port}");
        db = redis.GetDatabase();
        
    }
    private RedisChannel GetChannel(string key)
    {
        if(channels.ContainsKey(key))return channels[key];
        RedisChannel c = new(key, RedisChannel.PatternMode.Literal);
        channels.Add(key, c);
        return c;
    }
    public void Push(string key, byte[] data)
    {
        RedisChannel c = GetChannel(key);
        db.Publish(c, data);
    }
    public void Dispose()
    {
        if(disposed)return;
        disposed = true;
        redis.Dispose();
    }
}