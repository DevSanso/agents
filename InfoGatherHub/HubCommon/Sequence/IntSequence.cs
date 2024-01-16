namespace InfoGatherHub.HubCommon.Sequence;

using System.Threading;

public class IntSequence : ISequence<int>
{
    private int num = 0;
    public int Next()
    {
        return Interlocked.Increment(ref num);
    }

}