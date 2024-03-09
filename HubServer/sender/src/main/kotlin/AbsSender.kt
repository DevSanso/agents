package devsanso.github.io.HubServer.sender

import devsanso.github.io.HubServer.common.data.DataCotent
import devsanso.github.io.HubServer.global.GlobalSingleTon

abstract class AbsSender {
    protected abstract fun sendImpl(data: DataCotent)

    fun send(data: DataCotent) {
        try {
            sendImpl(data)
        }catch(e : Exception) {
            GlobalSingleTon.logger.error(javaClass, e.message ?: "")
            throw e
        }
    }
}