package devsanso.github.io.HubServer.receiver

import java.net.URI
import java.util.concurrent.ArrayBlockingQueue
import java.io.Closeable

import redis.clients.jedis.Jedis
import redis.clients.jedis.*

import devsanso.github.io.HubServer.receiver.data.RedisReceiverConfig


class RedisReceiver(channel : String, config: RedisReceiverConfig)
    : Runnable, Closeable {
    private class RedisReceiverImpl(val channel : String, val config: RedisReceiverConfig)
        : BinaryJedisPubSub(), Closeable {
        val recvQ : ArrayBlockingQueue<ByteArray> = ArrayBlockingQueue<ByteArray>(100)
        val connection : Jedis
        private fun connectUrl() : URI = URI("redis://${config.user}:${config.password}@"
                +"${config.ip}:${config.port}/${config.db}")
        init {
            connection = Jedis(connectUrl(), config)
        }
        override fun close() {
            connection.close()
            recvQ.clear()
        }

        override fun onMessage(channel: ByteArray?, message: ByteArray?) {
            if (message != null) {
                recvQ.put(message)
            }
            super.onMessage(channel, message)
        }
    }

    private val impl : RedisReceiverImpl = RedisReceiverImpl(channel, config)

    fun recv() : ByteArray? =
        if(impl.recvQ.size > 0) impl.recvQ.take()
        else null

    override fun run() {
        try {
            impl.connection.subscribe(impl, impl.channel.toByteArray())
        }catch(e : Exception) {
            throw e
        }
    }

    override fun close() {
        impl.close()
    }

}