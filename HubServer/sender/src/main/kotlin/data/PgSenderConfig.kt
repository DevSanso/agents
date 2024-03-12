package devsanso.github.io.HubServer.sender.data

data class PgSenderConfig(val ip : String, val port : Int, val dbName : String,
                          val user : String, val password : String) {

    val connurl: String = "jdbc:postgresql://$ip:$port/$dbName"
}
