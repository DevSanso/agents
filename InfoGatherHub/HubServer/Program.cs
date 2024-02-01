using System;
using System.Net.Sockets;
using System.Threading;
using System.Collections.Concurrent;


using InfoGatherHub.HubCommon.Format;
using InfoGatherHub.HubGlobal.Logger.Extension.Line;
using InfoGatherHub.HubGlobal.Config.Extension.Toml;
using InfoGatherHub.HubGlobal;
using InfoGatherHub.HubCommon.Display;
using InfoGatherHub.HubGlobal.Config;
using InfoGatherHub.HubServer.Config;
using InfoGatherHub.HubServer.Global.Extend;

string configPath = args[1];

var g = GlobalProvider<Config, GlobalExtend>.Init();

g.LoadToml(configPath);


g.InitLogLine(new DisplayConsole());

var config = g.GetConfig();