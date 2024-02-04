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


string configPath = args[1];

List<IWorker> snapWorkers = new List<IWorker>();
List<Timer> snapTimers = new List<Timer>();

var g = GlobalProvider<Config>.Init();

g.LoadToml(configPath);

g.InitLogLine(new DisplayConsole());

var config = g.GetConfig()!;

ConcurrentQueue<IFormat<WorkerFormatHeader>> q = new ConcurrentQueue<IFormat<WorkerFormatHeader>>();

if(config.osSnapSetting != null) {
    IWorker osSnapWorker = new ReadOsSnapWorker(new MemMapClient(config.osSnapSetting.path, config.osSnapSetting.size), q);
    Timer osSnapTimer = new Timer(
        state => ((IWorker?)state)?.Work(),
        osSnapWorker,
        TimeSpan.Zero,
        TimeSpan.FromMilliseconds(1050)
    );
    snapTimers.Add(osSnapTimer);
    snapWorkers.Add(osSnapWorker);
}

if(config.redisSnapSetting != null) {
    IWorker redisSnapWorker = new ReadRedisSnapWorker(new MemMapClient(config.redisSnapSetting.path, config.redisSnapSetting.size), q);
    Timer redisSnapTimer = new Timer(
        state => ((IWorker?)state)?.Work(),
        redisSnapWorker,
        TimeSpan.Zero,
        TimeSpan.FromMilliseconds(1050)
    );
    snapTimers.Add(redisSnapTimer);
    snapWorkers.Add(redisSnapWorker);
}

IWorker sendDataWorker = new SendDataWorker(new TcpClient(config.hubServerSetting.ip, config.hubServerSetting.port), q);

Timer sendDataTimer = new Timer(
    state => ((IWorker?)state)?.Work(),
    sendDataWorker,
    TimeSpan.Zero,
    TimeSpan.FromMilliseconds(500)
);

Console.CancelKeyPress += (sender, e) =>
{
    g.Log(LogLevel.Debug, LogCategory.ALL, "HubSender signal ctrl + c");

    sendDataTimer.Dispose();
    sendDataWorker.Dispose();

    foreach(var timer in snapTimers)
    {
        timer.Dispose();
    }

    foreach(var worker in snapWorkers)
    {
        worker.Dispose();
    }

    g.Log(LogLevel.Debug, LogCategory.ALL, "HubSender is shutdown now");
    Environment.Exit(0);
};

while(true)
{
    Thread.Sleep(1000);
}




