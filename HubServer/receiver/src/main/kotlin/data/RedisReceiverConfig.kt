package devsanso.github.io.receiver.data

import redis.clients.jedis.JedisClientConfig

data class RedisReceiverConfig(val ip : String,
                               val port : Int, val userName : String, val passwd : String,
                               val db : Int, val timeout : Int) : JedisClientConfig {

    override fun getUser(): String = userName
    override fun getPassword(): String = passwd
    override fun isSsl(): Boolean = false
    override fun getDatabase(): Int = db
    override fun getClientName(): String = "HubServerSub"
}