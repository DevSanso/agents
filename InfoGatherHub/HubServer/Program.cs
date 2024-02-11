using System;
using System.Net.Sockets;
using System.Threading;
using System.Collections.Concurrent;


using InfoGatherHub.HubGlobal.Logger.Extension.Line;
using InfoGatherHub.HubGlobal.Config.Extension.Toml;
using InfoGatherHub.HubGlobal;
using InfoGatherHub.HubGlobal.Extend;
using InfoGatherHub.HubGlobal.Config;

using InfoGatherHub.HubCommon.Display;
using InfoGatherHub.HubServer.Config;
using InfoGatherHub.HubServer.Global.Extend;
using InfoGatherHub.HubServer.Server;
using InfoGatherHub.HubServer;

string configPath = args[1];

var g = GlobalProvider<Config, GlobalExtend>.Init();

g.LoadToml(configPath);


g.InitLogLine(new DisplayConsole());

var config = g.GetConfig();

var server = new TcpSocketServer(config!.ServerConfig.ServerAddress, config!.ServerConfig.Port);

var mainThread = new MainThread(server);
mainThread.Start();