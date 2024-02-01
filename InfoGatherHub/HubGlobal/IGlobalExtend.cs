namespace InfoGatherHub.HubGlobal;

public interface IGlobalExtend<T>
{
    void Init<T2>(in Global<T, T2> g)  where T2 : IGlobalExtend<T>, new();
}