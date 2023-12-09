using System;
using System.Net.Sockets;
using System.Threading;
using System.Collections.Concurrent;


using InfoGatherHub.HubCommon.Format;
using InfoGatherHub.HubGlobal.Logger.Extension.Xml;
using InfoGatherHub.HubGlobal.Config.Extension.Toml;
using InfoGatherHub.HubGlobal;
using InfoGatherHub.HubCommon.Display;
using InfoGatherHub.HubGlobal.Config;
using InfoGatherHub.HubGlobal.Logger;
using InfoGatherHub.HubServer;
using InfoGatherHub.HubServer.Mapping;
using InfoGatherHub.HubServer.Global.Extend;

string configPath = args[1];

var g = GlobalProvider<Config, GlobalExtend>.Init(new GlobalExtend());

g.LoadToml(configPath);

g.InitXml(new DisplayConsole());

var config = g.GetConfig();