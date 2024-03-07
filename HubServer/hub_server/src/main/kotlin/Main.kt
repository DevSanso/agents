package devsanso.github.io.HubServer.hub_server

import java.net.URL
import java.net.URLClassLoader
import devsanso.github.io.HubServer.common.Common

fun main() {

    val cl = ClassLoader.getSystemClassLoader()

    println("Hello World!")
    val com = Common()
    com.hello()
}