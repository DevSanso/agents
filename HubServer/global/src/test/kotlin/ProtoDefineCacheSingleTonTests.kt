import org.junit.jupiter.api.Test

import devsanso.github.io.HubServer.global.ProtoDefineCacheSingleTon

internal class ProtoDefineCacheSingleTonTests {
    @Test
    internal fun testProtoDefineCacheSingleTon() {
        ProtoDefineCacheSingleTon.RedisCache.forEach {
            println(it.key)
            for (index in 0..<it.value.size()) {
                println(it.value[index])
            }
        }
    }
}