namespace InfoGatherHub.HubCommon.Tests;

using Microsoft.VisualStudio.TestTools.UnitTesting;
using InfoGatherHub.HubCommon.Cache;


enum TestEnum
{
    Test1,
    Test2,
    Test3
}
[TestClass]
public class CacheTests
{
    EnumToStringCache<TestEnum> cache = new EnumToStringCache<TestEnum>();
    [TestMethod]
    public void TestMethod1()
    {
        Assert.AreEqual("Test1", cache.Get(TestEnum.Test1));
        cache.Get(TestEnum.Test2);
        cache.Get(TestEnum.Test3);
        Assert.AreEqual("Test1", cache.Get(TestEnum.Test1));
        Assert.AreEqual("Test2", cache.Get(TestEnum.Test2));
        Assert.AreEqual("Test3", cache.Get(TestEnum.Test3));

    }
}