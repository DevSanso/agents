namespace InfoGatherHub.HubGlobal.HubGlobalTests;

using Microsoft.VisualStudio.TestTools.UnitTesting;
using InfoGatherHub.HubGlobal.Logger.Extension.Line;
using InfoGatherHub.HubGlobal.Logger;
using InfoGatherHub.HubGlobal;
using InfoGatherHub.HubCommon.Display;
using System.Threading;
using System;
using System.Diagnostics;

[TestClass]
public class LineLoggerTests
{
    static int lsRun = 0;
    public class Empty
    {

    }

    public class MockDisplay : IDisplay
    {
        public MockDisplay()
        {
        }

        public void Display(string message)
        {
            lsRun = 1;
            Trace.WriteLine(message);
        }
    }
    [TestMethod]
    public void TestLineLogger()
    {
        Trace.Listeners.Add(new TextWriterTraceListener(Console.Out));
        Trace.WriteLine("Hello World");
        Global<Empty> g = GlobalProvider<Empty>.Init();

        g.InitLogLine(new MockDisplay());

        g.Log(LogLevel.Error, LogCategory.ALL, "TestLineLogger");
        Thread.Sleep(3000);
        Assert.AreEqual(lsRun, 1);
    }
}