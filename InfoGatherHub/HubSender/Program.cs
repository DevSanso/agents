using InfoGatherHub.HubGlobal.Logger;
using InfoGatherHub.HubGlobal.Logger.Extension.Line;
using InfoGatherHub.HubGlobal.Config.Extension.Toml;
using InfoGatherHub.HubGlobal;
using InfoGatherHub.HubSender;
using InfoGatherHub.HubCommon.Display;
using InfoGatherHub.HubGlobal.Config;

if (args.Length != 2)
{
    Console.WriteLine("Usage: HubSender <configPath>");
    Environment.Exit(1);
}

string configPath = args[1];

var g = GlobalProvider<Config>.Init();

g.LoadToml(configPath);

g.InitLogLine(new DisplayConsole());

var config = g.GetConfig()!;

var contorller = new WorkerController(config);
contorller.RunAndNonBlocking();

Console.CancelKeyPress += (sender, e) =>
{
    g.Log(LogLevel.Debug, LogCategory.ALL, "HubSender signal ctrl + c");
    contorller.Dispose();
    g.Log(LogLevel.Debug, LogCategory.ALL, "HubSender is shutdown now");
    Environment.Exit(0);
};

while(true)
{
    Thread.Sleep(5000);
}




