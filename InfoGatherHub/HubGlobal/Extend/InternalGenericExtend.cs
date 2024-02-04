namespace InfoGatherHub.HubGlobal.Extend;

using InfoGatherHub.HubGlobal;

public static class InternalGenericExtend
{
    private static object extend;
    public static Global<T> InitExtend<T,E>(this IGlobalExtend<E> gp, E extend) where E : IGenericExtend<T>
    {
        InternalGenericExtend.extend = extend;
        var g = GlobalProvider<T>.Init()!;
        extend.Init(g);
        return g;
    }

    public static E GetExtend<T,E>(this Global<T,E> g) where E : class
    {
        return (E)extend;
    }
}