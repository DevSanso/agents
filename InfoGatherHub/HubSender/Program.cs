using System;
using System.Net.Sockets;
using System.Threading;
using System.Collections.Concurrent;

using InfoGatherHub.HubCommon.Format;
using InfoGatherHub.HubGlobal.Logger.Extension.Xml;
using InfoGatherHub.HubGlobal.Config.Extension.Toml;
using InfoGatherHub.HubGlobal;
using InfoGatherHub.HubSender;
using InfoGatherHub.HubSender.Snap;
using InfoGatherHub.HubCommon.Display;
using InfoGatherHub.HubGlobal.Config;
using InfoGatherHub.HubSender.Worker;
using InfoGatherHub.HubGlobal.Logger;

string configPath = args[1];

var g = GlobalProvider<Config,object>.Init(null);

g.LoadToml(configPath);

g.InitXml(new DisplayConsole());

var config = g.GetConfig()!;

ConcurrentQueue<IFormat<InfoGatherHub.HubCommon.Format.Void>> q = new ConcurrentQueue<IFormat<InfoGatherHub.HubCommon.Format.Void>>();

IWorker osSnapWorker = new ReadOsSnapWorker(new MemMapClient(config.osSnapSetting.path, config.osSnapSetting.size), q);
IWorker sendDataWorker = new SendDataWorker(new TcpClient(config.hubServerSetting.ip, config.hubServerSetting.port), q);

Timer osSnapTimer = new Timer(
    state => ((IWorker?)state)?.Work(),
    osSnapWorker,
    TimeSpan.Zero,
    TimeSpan.FromMilliseconds(1050)
);

Timer sendDataTimer = new Timer(
    state => ((IWorker?)state)?.Work(),
    sendDataWorker,
    TimeSpan.Zero,
    TimeSpan.FromMilliseconds(500)
);

Console.CancelKeyPress += (sender, e) =>
{
    g.Log(LogLevel.Debug, LogCategory.ALL, "HubSender signal ctrl + c");

    osSnapTimer.Dispose();
    sendDataTimer.Dispose();
    osSnapWorker.Dispose();
    sendDataWorker.Dispose();

    g.Log(LogLevel.Debug, LogCategory.ALL, "HubSender is shutdown now");
    Environment.Exit(0);
};

while(true)
{
    Thread.Sleep(100);
}




