package devsanso.github.io.HubServer.receiver

import java.net.URI
import java.util.concurrent.ArrayBlockingQueue
import java.io.Closeable

import redis.clients.jedis.Jedis
import redis.clients.jedis.*

import devsanso.github.io.HubServer.receiver.data.RedisReceiverConfig


class RedisReceiver constructor(val channel : String, val config: RedisReceiverConfig)
    : BinaryJedisPubSub(), Closeable, Runnable {
    private val recvQ : ArrayBlockingQueue<ByteArray> = ArrayBlockingQueue<ByteArray>(100)
    private val connection : Jedis
    private fun connectUrl() : URI = URI("redis://${config.user}:${config.password}@"
            +"${config.ip}:${config.port}/${config.db}")

    init {
        connection = Jedis(connectUrl(), config)
    }

    fun recv() : ByteArray? =
        if(recvQ.size > 0) recvQ.take()
        else null

    override fun run() {
        try {
            connection.subscribe(this, channel.toByteArray())
        }catch(e : Exception) {
            throw e
        }
    }

    override fun onMessage(channel: ByteArray?, message: ByteArray?) {
        if (message != null) {
            recvQ.put(message)
        }
        super.onMessage(channel, message)
    }
    override fun close() {
        connection.close()
        recvQ.clear()
    }




}