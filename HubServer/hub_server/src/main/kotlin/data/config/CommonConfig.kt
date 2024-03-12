package devsanso.github.io.HubServer.hub_server.data.config

class CommonConfig constructor() {
    val logLevel : String = ""
    val logPath : String = ""

    val receiverConfig : ReceiverConfig = ReceiverConfig()
    val senderConfig : SenderConfig = SenderConfig()
}