package devsanso.github.io.HubServer.sender

import devsanso.github.io.HubServer.sender.data.PgSenderConfig
import devsanso.github.io.HubServer.common.data.DataCotent
import devsanso.github.io.HubServer.common.structure.Pool
import org.apache.ibatis.jdbc.SQL
import java.sql.Connection
import java.sql.DriverManager
import java.util.*

class PgSender constructor(val config : PgSenderConfig, poolMax : Int) : AbsSender() {
    private val connPool : Pool<Connection> = Pool(poolMax, this.genConnection())
    private val props = Properties().apply {
        this.setProperty("user", config.user)
        this.setProperty("password", config.password)
        this.setProperty("ssl", "false")
        this.setProperty("connectTimeout", "60")
        this.setProperty("ApplicationName", "Agent - Hub Server")
    }

    init {
        Class.forName("org.postgresql.Driver")
    }

    private fun genConnection() : () -> Connection {
        val url = config.connurl
        return {
            val conn = DriverManager.getConnection(url, props)

            conn.autoCommit = false
            conn
        }
    }

    private fun makeQuery(data : DataCotent) : String = data.data.let {
        val s = SQL()
        s.INSERT_INTO(data.objectName)

        for(pair in it) {
            s.VALUES(pair.key, pair.value.toString())
        }

        s.toString()
    }

    override fun sendImpl(data: DataCotent) {
        val sql = makeQuery(data)
        val conn = connPool.pop(this.javaClass)

        conn?.runElement { c ->
            try {
                c.createStatement().use {
                    it.execute(sql)
                }
                c.commit()
            }catch(e : Exception) {
                c.rollback()
                throw e
            }
        }

        conn?.close()
    }
}