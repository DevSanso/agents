package devsanso.github.io.HubServer.hub_server.thread

import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch

import devsanso.github.io.HubServer.hub_server.data.config.ReceiverConfig
import devsanso.github.io.HubServer.hub_server.data.dto.convert
import devsanso.github.io.HubServer.hub_server.worker.AbsWorker
import devsanso.github.io.HubServer.hub_server.global.QueueMap
import devsanso.github.io.HubServer.receiver.RedisReceiver

class ReceiverWorker constructor(private val config: ReceiverConfig) : AbsWorker() {
    private var receiver : RedisReceiver? = null
    private var isRunThread : Boolean = false
    private var isNeedRunThread : Boolean = false
    private var getData : ByteArray? = null

    private val sendQueue = QueueMap.get(QueueMap.QueueMapKey.RECEIVTER_QUEUE)

    @Synchronized
    override fun update() {
        if(!isRunThread) isNeedRunThread = true
    }
    @Synchronized
    override fun work() {
        if(!isNeedRunThread) {
            CoroutineScope(Dispatchers.IO).launch {
                receiver = RedisReceiver(config.channel, config.convert())
                isRunThread = true
                receiver!!.run()
                receiver!!.close()
                receiver = null
                isRunThread = false
            }
        }
        else {
            getData = receiver!!.recv()
        }
    }
    @Synchronized
    override fun fetch() {
        if(getData != null)sendQueue.add(getData!!)
    }

}