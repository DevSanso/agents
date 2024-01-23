namespace InfoGatherHub.HubCommon.Sequence;

using System.Threading;

public class LongSequence : ISequence<long>
{
    private long num = 0;
    public long Next()
    {
        return Interlocked.Increment(ref num);
    }
    public void Reset()
    {
        Interlocked.Exchange(ref num, 0);
    }
    public long Current()
    {
        return Interlocked.Read(ref num);
    }

}