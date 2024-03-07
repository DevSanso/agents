
import org.junit.jupiter.api.*
import redis.clients.jedis.*

import devsanso.github.io.receiver.RedisReceiver
import devsanso.github.io.receiver.data.RedisReceiverConfig
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import java.util.concurrent.TimeUnit

internal class RedisReceiverTest {
    val host = "172.17.0.2"
    val port = 6379
    val db = 0
    val user = "test"
    val password = "test"

    val config = RedisReceiverConfig(host,port,user,password,db,10)
    @Test
    internal fun testReceiver() {
        val receiver = RedisReceiver("test", config)
        val sender = Jedis(host,port,config)

        CoroutineScope(Dispatchers.IO).launch {
            receiver.run()
        }
        TimeUnit.SECONDS.sleep(3)

        sender.use {
            it.publish("test".toByteArray(), "test".toByteArray())
            it.publish("test".toByteArray(), "test1".toByteArray())
            it.publish("test".toByteArray(), "test2".toByteArray())
        }

        TimeUnit.SECONDS.sleep(5)

        receiver.use {
            println(it.recv()?.toString(Charsets.UTF_8))
            println(it.recv()?.toString(Charsets.UTF_8))
            println(it.recv()?.toString(Charsets.UTF_8))
        }

    }
}