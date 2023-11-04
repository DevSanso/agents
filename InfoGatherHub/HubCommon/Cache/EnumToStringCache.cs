namespace InfoGatherHub.HubCommon.Cache;

using System.Collections.Generic;

public class EnumToStringCache<T> where T :  System.Enum
{
    private Dictionary<T, string> dict = new Dictionary<T, string>();

    public string get(T value)
    {
        string? output = "";

        bool isCatched = dict.TryGetValue(value, out output);
        if(isCatched == true)
        {
            return output!;
        }
        else
        {
            dict.Add(value, value.ToString());
            return value.ToString();
        }

        
    }
}