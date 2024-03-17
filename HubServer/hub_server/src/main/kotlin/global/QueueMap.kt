package devsanso.github.io.HubServer.hub_server.global

import java.util.*
import java.util.concurrent.ArrayBlockingQueue

object QueueMap {
    enum class QueueMapKey {
        RECEIVTER_QUEUE
    }
    private val queueMap : Map<QueueMapKey, ArrayBlockingQueue<Any>> =
        EnumMap<QueueMapKey, ArrayBlockingQueue<Any>>(QueueMapKey::class.java).apply {
            for(i in this.keys) {
                this[i] = ArrayBlockingQueue(100)
            }
        }

    fun get(index : QueueMapKey) : ArrayBlockingQueue<Any> {
        val q = queueMap[index] ?: throw Exception("")
        return q
    }
}