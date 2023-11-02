namespace InfoGatherHub.HubGlobal.Logger;



public enum LogLevel
{
    Debug,
    Error
}
public enum LogCategory
{
    ALL,
    Network,
    IO,
    Code
}
public interface ILogger {}