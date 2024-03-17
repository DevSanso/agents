package devsanso.github.io.HubServer.hub_server.worker

abstract class AbsWorker() {
    abstract fun update()
    abstract fun work()
    abstract fun fetch()
}
