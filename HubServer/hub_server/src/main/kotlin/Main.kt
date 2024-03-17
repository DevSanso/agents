package devsanso.github.io.HubServer.hub_server

import java.net.URL

import devsanso.github.io.HubServer.hub_server.data.config.CommonConfig
import devsanso.github.io.HubServer.common.loader.ConfigLoader
import devsanso.github.io.HubServer.hub_server.worker.toolchain.RedisWorkerToolChain
fun main(args : Array<String>) {
    val config = ConfigLoader.load<CommonConfig>(URL(args[1]))

}