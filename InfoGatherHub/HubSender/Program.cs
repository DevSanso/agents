using System;

using InfoGatherHub.HubCommon.Observer;
using InfoGatherHub.HubGlobal.Logger.Extension.Xml;
using InfoGatherHub.HubGlobal.Config.Extension.Toml;
using InfoGatherHub.HubGlobal;
using InfoGatherHub.HubSender;
using InfoGatherHub.HubSender.Snap;
using InfoGatherHub.HubCommon.Display;

string configPath = args[1];

var g = GlobalProvider<Config>.Global();

g.LoadToml(configPath);

var config = g.GetToml();

g.InitXml(config.logSetting.logPath, new DisplayConsole());






