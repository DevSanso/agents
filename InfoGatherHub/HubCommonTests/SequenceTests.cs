namespace InfoGatherHub.HubCommon.Tests;

using Microsoft.VisualStudio.TestTools.UnitTesting;
using InfoGatherHub.HubCommon.Sequence;

[TestClass]
public class SequenceTests
{
    [TestMethod]
    public void SequenceTest()
    {
        ISequence<long> sequence = new LongSequence();
        Assert.AreEqual(1, sequence.Next());
        Assert.AreEqual(2, sequence.Next());
        Assert.AreEqual(3, sequence.Next());
        sequence.Reset();
        Assert.AreEqual(1, sequence.Next());
    }
}