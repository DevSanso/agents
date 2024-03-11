import devsanso.github.io.HubServer.common.loader.ConfigLoader
import org.junit.jupiter.api.Test
import java.io.File
import kotlin.test.assertEquals

class SampleConfig constructor() {
    val name : String = ""
    val version : String = ""
    val port : Int = 0

    val props : Map<String,String> = mapOf()
}

class ConfigLoggerTests {
    val data = """
    name = "Sample Name"
    version = "1.0.0"
    port = 8080

    [props]
    key1 = "value1"
    key2 = "value2"
    key3 = "value3"
    """
    @Test
    internal fun ConfigReadTest() {
        val config = ConfigLoader.load<SampleConfig>(data.toByteArray())

        assertEquals(config.name, "Sample Name")
        assertEquals(config.port, 8080)
        assertEquals(config.version, "1.0.0")
        assertEquals(config.props["key1"],"value1")
        assertEquals(config.props["key2"],"value2")
        assertEquals(config.props["key3"],"value3")
    }
}