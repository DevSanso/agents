package devsanso.github.io.HubServer.global

import devsanso.github.io.HubServer.common.adapter.logger.ILoggerAdapter
import devsanso.github.io.HubServer.common.adapter.logger.LogLevel
import devsanso.github.io.HubServer.common.loader.LoggerLoader

object GlobalSingleTon {
    val props : Map<String, String> = mapOf()

    private fun String.logLevelConvert() : LogLevel = when {
        this == "DEBUG" -> LogLevel.Debug
        this == "INFO" -> LogLevel.Info
        this == "PANIC" -> LogLevel.Panic
        this == "ERROR" -> LogLevel.Error
        else -> throw Exception("GlobalSingleTon : LogLevel not supported $this")
    }

    val logger : ILoggerAdapter by lazy {
        val filepath = props["logfile"]
        val logLevel = props["loglevel"]!!.logLevelConvert()
        LoggerLoader.createLogger("logger",filepath, logLevel)
    }
}