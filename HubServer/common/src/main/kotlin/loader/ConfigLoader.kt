package devsanso.github.io.HubServer.common.loader

import com.fasterxml.jackson.dataformat.toml.TomlMapper
import java.io.File
import java.net.URL

class ConfigLoader private constructor() {
    companion object {
        inline fun <reified T>load(path : URL) :T {
            val mapper = TomlMapper()
            return mapper.readValue(File(path.toString()), T::class.java)
        }

        inline fun <reified T>load(bytes : ByteArray) :T {
            val mapper = TomlMapper()
            return mapper.readValue(bytes, T::class.java)
        }
    }
}