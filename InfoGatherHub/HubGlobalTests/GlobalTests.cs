namespace InfoGatherHub.HubGlobal.HubGlobalTests;

using Microsoft.VisualStudio.TestTools.UnitTesting;
using InfoGatherHub.HubGlobal;
using System;
using System.Diagnostics;
using InfoGatherHub.HubGlobal.Config;
using InfoGatherHub.HubGlobal.Config.Extension.Toml;

[TestClass]
public class GlobalTests
{
    class Sample
    {
        public int A {get;set;}
    }

    public void SetSample()
    {
        
        var sample = GlobalProvider<Sample>.Global().GetConfig();
        sample!.A = 1234;
    }
    
    [TestMethod]
    public void TestGlobal()
    {
        Global<Sample> g = GlobalProvider<Sample>.Init();
        g.LoadTemp();

        SetSample();
        Trace.Listeners.Add(new TextWriterTraceListener(Console.Out));
        Global<Sample> g1 = GlobalProvider<Sample>.Global();
        var sample = g1.GetConfig();
        Assert.AreEqual(sample!.A, 1234);
    }
}