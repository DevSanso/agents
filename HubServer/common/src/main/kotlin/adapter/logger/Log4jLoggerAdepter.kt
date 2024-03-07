package devsanso.github.io.HubServer.common.adapter.logger


import org.apache.logging.log4j.Level
import org.apache.logging.log4j.core.LoggerContext
import org.apache.logging.log4j.core.appender.FileAppender
import org.apache.logging.log4j.core.appender.ConsoleAppender
import org.apache.logging.log4j.core.config.AppenderRef
import org.apache.logging.log4j.core.layout.JsonLayout
import org.apache.logging.log4j.core.layout.PatternLayout
import kotlin.math.log


class Log4jLoggerAdepter internal constructor(logName : String, filename : String?, level: LogLevel) : ILoggerAdapter(){
    val context : LoggerContext

    private fun LogLevel.convert() : Level = when {
        this == LogLevel.Info -> Level.INFO
        this == LogLevel.Error -> Level.ERROR
        this == LogLevel.Panic -> Level.FATAL
        else -> Level.DEBUG
    }

    private fun layout() : PatternLayout = PatternLayout.newBuilder()
        .withCharset(Charsets.UTF_8)
        .withPattern("%d{yyyy-MM-dd}-%t-%-5p-%-10c:%m%n")
        .build()

    private fun levelLogger(filename : String) : FileAppender = FileAppender.newBuilder()
        .withFileName(filename)
        .setLayout(layout())
        .setName(FileAppender.PLUGIN_NAME)
        .build()

    private fun allLogger() : ConsoleAppender = ConsoleAppender.newBuilder()
        .setLayout(layout())
        .setName(ConsoleAppender.PLUGIN_NAME)
        .build()

    init {
        context = LoggerContext(logName)
        val config = context.configuration
        val loggerConfig = config.getLoggerConfig(logName)

        loggerConfig.level = level.convert()

        context.setConfiguration(config)
        context.updateLoggers()

        if (filename != null) {
            val fileAppender = levelLogger(filename)
            fileAppender.start()
            loggerConfig.addAppender(fileAppender, level.convert(), null)
        }

        val consoleAppender = allLogger()
        consoleAppender.start()
        loggerConfig.addAppender(consoleAppender, level.convert(), null)
    }

    override fun debug(jClass: Class<*>, message: String) {
        context.getLogger(jClass).debug(message)
    }

    override fun info(jClass: Class<*>, message: String) {
        context.getLogger(jClass).info(message)
    }

    override fun error(jClass: Class<*>, message: String) {
        context.getLogger(jClass).error(message)
    }

    override fun panic(jClass: Class<*>, message: String) {
        context.getLogger(jClass).fatal(message)
    }
}