namespace InfoGatherHub.HubCommon.Cache;

using System.Collections.Generic;

public class EnumToStringCache<T> where T :  System.Enum
{
    private Dictionary<T, string> dict = new Dictionary<T, string>();

    private bool isExist(T value)
    {
        return dict.ContainsKey(value);
    }
    public string get(T value)
    {
        if(isExist(value) == false)dict.Add(value, value.ToString());
        string? output = "";

        dict.TryGetValue(value, out output);
        return output!;
    }
}