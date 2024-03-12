package devsanso.github.io.HubServer.hub_server.data.config

typealias DataTarget = String
typealias DataType = String

class PgSenderConfig constructor() {
    class DataTableSchema constructor() {
        val tableName : String = ""
        val dataMapping : Map<String,String> = mapOf()
    }

    val ip : String = ""
    val port : Int = 0
    val dbName : String = ""
    val user : String = ""
    val password : String = ""

    val dataSchema : Map<Pair<DataTarget,DataType>, DataTableSchema> = mapOf()
}

class SenderConfig constructor(){
    val type : String = ""

    val pgSenderConfig : PgSenderConfig? = null
}