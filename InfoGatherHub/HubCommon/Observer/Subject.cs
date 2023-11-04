namespace InfoGatherHub.HubCommon.Observer;

using System;
using System.Collections.Generic;

class Subject<T> where T : struct
{
    private List<IObserver<T>> observers = new List<IObserver<T>>();
    private T data;

    public T Data
    {
        get { return data; }
        set
        {
            data = value;
            NotifyObservers();
        }
    }

    public void RegisterObserver(IObserver<T> observer)
    {
        observers.Add(observer);
    }

    public void RemoveObserver(IObserver<T> observer)
    {
        observers.Remove(observer);
    }

    private void NotifyObservers()
    {
        foreach (var observer in observers)
        {
            observer.Update(data);
        }
    }
}