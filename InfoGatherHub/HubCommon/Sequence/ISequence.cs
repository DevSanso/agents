namespace InfoGatherHub.HubCommon.Sequence;


public interface ISequence<T>
{
    T Next();
    T Current();
    void Reset();
}