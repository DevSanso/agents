package devsanso.github.io.HubServer.common.loader

import devsanso.github.io.HubServer.common.adapter.logger.ILoggerAdapter
import devsanso.github.io.HubServer.common.adapter.logger.Log4jLoggerAdepter
import devsanso.github.io.HubServer.common.adapter.logger.LogLevel

class LoggerLoader private constructor(){
    companion object {
        fun createLogger(logName : String, filename : String?, level: LogLevel) : ILoggerAdapter {
            return Log4jLoggerAdepter(logName, filename, level)
        }
    }
}