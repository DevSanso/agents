package devsanso.github.io.HubServer

import java.net.URL
import java.net.URLClassLoader

fun main() {

    val cl = ClassLoader.getSystemClassLoader()

    println("Hello World!")
    val com = Common()
    com.hello()
}