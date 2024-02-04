namespace InfoGatherHub.HubGlobal.Extend;

using InfoGatherHub.HubGlobal;

public interface IGlobalExtend<E> {}

public interface IGenericExtend<T> 
{
    void Init(Global<T> g);
}
