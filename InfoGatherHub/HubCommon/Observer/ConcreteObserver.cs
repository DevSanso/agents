namespace InfoGatherHub.HubCommon.Observer;

using System;

class ConcreteObserver<T> : IObserver<T>
{
    private string name;
    private Action<T> action;

    public ConcreteObserver(string name, Action<T> action)
    {
        this.name = name;
        this.action = action;
    }

    public void Update(T data)
    {
        try 
        {
            this.action(data);
        }
        catch(Exception e)
        {
            throw new Exception($"Observer : {this.name} -> {e.ToString()}");
        }
    }
}
