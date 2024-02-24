namespace InfoGatherHub.HubSender;

using System;
using System.Net.Sockets;
using System.Threading;
using System.Collections.Generic;
using System.Collections.Concurrent;

using InfoGatherHub.HubCommon.Format;
using InfoGatherHub.HubGlobal.Logger.Extension.Line;
using InfoGatherHub.HubGlobal.Config.Extension.Toml;
using InfoGatherHub.HubGlobal;
using InfoGatherHub.HubSender;
using InfoGatherHub.HubSender.Ipc;
using InfoGatherHub.HubCommon.Display;
using InfoGatherHub.HubGlobal.Config;
using InfoGatherHub.HubSender.Worker;
using InfoGatherHub.HubGlobal.Logger;
using InfoGatherHub.HubSender.Worker.Format;
using InfoGatherHub.HubSender.Pusher;



internal class WorkerController : IDisposable
{
    private readonly Dictionary<string,IWorker> snapWorkers = new Dictionary<string,IWorker>();
    private readonly Dictionary<string,Timer> snapTimers = new Dictionary<string,Timer>();
    private IWorker sendDataWorker;
    private Timer sendDataTimer;
    private bool isRun = false;
    private bool isDispose = false;

    ConcurrentQueue<IFormat<WorkerFormatHeader>> sendQ = new ConcurrentQueue<IFormat<WorkerFormatHeader>>();

    public WorkerController(Config config)
    {
        sendDataWorker = new SendDataWorker(new RedisPusher(config.PusherSetting.ip, config.PusherSetting.port), sendQ);

        sendDataTimer = new Timer(
            state => ((IWorker?)state)?.Work(),
            sendDataWorker,
            TimeSpan.Zero,
            TimeSpan.FromMilliseconds(500)
        );

        InitSnapWorker(config);
    }

    private void InitSnapWorker(Config config)
    {
        if(config.osSnapSetting != null) {
            IWorker osSnapWorker = new ReadOsSnapWorker(new MemMapClient(config.osSnapSetting.path, config.osSnapSetting.size), sendQ);
            snapWorkers.Add("os", osSnapWorker);
        }


        if(config.redisSnapSetting != null) {
            IWorker redisSnapWorker = new ReadRedisSnapWorker(new MemMapClient(config.redisSnapSetting.path, config.redisSnapSetting.size), sendQ);
            snapWorkers.Add("redis", redisSnapWorker);
        }
    }

    public void RunAndNonBlocking()
    {
        if(isRun)
        {
            throw new Exception("Already Start");
        }
        isRun = true;
        foreach(var key in snapWorkers.Keys)
        {
            var worker = snapWorkers[key];
            Timer t = new Timer(
                state => ((IWorker?)state)?.Work(),
                worker,
                TimeSpan.Zero,
                TimeSpan.FromMilliseconds(1050)
            );

            snapTimers.Add(key, t);
        }

        sendDataTimer = new Timer(
            state => ((IWorker?)state)?.Work(),
            sendDataWorker,
            TimeSpan.Zero,
            TimeSpan.FromMilliseconds(2050)
        );
    }

    public void Dispose()
    {
        if(isDispose)
        {
            throw new Exception("already dispose");
        }
        isDispose = true;
        
        sendDataTimer.Dispose();
        sendDataWorker.Dispose();

        foreach(var timer in snapTimers.Values)
        {
            timer.Dispose();
        }

        foreach(var worker in snapWorkers.Values)
        {
            worker.Dispose();
        }

    }
}