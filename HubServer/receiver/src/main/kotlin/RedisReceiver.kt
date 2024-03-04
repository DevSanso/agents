package devsanso.github.io.receiver

import java.net.URI
import java.util.concurrent.ArrayBlockingQueue

import redis.clients.jedis.Jedis
import redis.clients.jedis.*

import devsanso.github.io.receiver.data.RedisReceiverConfig
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.Job
import kotlinx.coroutines.launch
import java.io.Closeable

class RedisReceiver constructor(val channel : String, val config: RedisReceiverConfig) : BinaryJedisPubSub(), Closeable {
    private val recvQ : ArrayBlockingQueue<ByteArray> = ArrayBlockingQueue<ByteArray>(100)
    private val connection : Jedis
    private fun connectUrl() : URI = URI("redis://${config.user}:${config.password}@"
            +"${config.ip}:${config.port}/${config.db}")

    init {
        connection = Jedis(connectUrl(), config)

        val that = this
        CoroutineScope(Dispatchers.IO).launch {
            connection.subscribe(that, channel.toByteArray())
        }
    }

    fun recv() : ByteArray? =
        if(recvQ.size > 0) recvQ.take()
        else null


    override fun onMessage(channel: ByteArray?, message: ByteArray?) {
        if (message != null) {
            recvQ.put(message)
        }
        super.onMessage(channel, message)
    }
    override fun close() = connection.close()




}