
import org.junit.jupiter.api.*

import devsanso.github.io.HubServer.common.adapter.logger.LogLevel
import devsanso.github.io.HubServer.common.loader.LoggerLoader

internal class LoggerAdepterTests {
    private final val tempDir : String = "/tmp"
    @Test
    internal fun testLogger() {
        val logger = LoggerLoader.createLogger("test logger",null, level = LogLevel.Debug)
        logger.debug(this.javaClass, "test logger")
        logger.error(this.javaClass, "test error logger")
        logger.info(this.javaClass, "test info logger")
    }

    @Test
    internal fun testFileLogger() {
        val logger = LoggerLoader.createLogger("test file logger", "$tempDir/temp.log", LogLevel.Error)
        logger.debug(this.javaClass, "test it can't write")
        logger.error(this.javaClass, "test error logger")
    }

}