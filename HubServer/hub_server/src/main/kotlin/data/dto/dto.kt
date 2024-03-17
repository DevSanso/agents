package devsanso.github.io.HubServer.hub_server.data.dto

import devsanso.github.io.HubServer.hub_server.data.config.*
import devsanso.github.io.HubServer.receiver.data.RedisReceiverConfig

fun ReceiverConfig.convert() : RedisReceiverConfig = RedisReceiverConfig(
    ip = this.ip,
    port = this.port,
    userName = this.username,
    passwd = this.password,
    db = this.dbName,
    timeout = 3
)