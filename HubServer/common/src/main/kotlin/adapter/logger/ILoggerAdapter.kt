package devsanso.github.io.HubServer.common.adapter.logger

abstract class ILoggerAdapter internal constructor(){
    abstract fun debug(jClass : Class<*>, message : String)
    abstract fun info(jClass : Class<*>, message : String)
    abstract fun error(jClass : Class<*>, message : String)
    abstract fun panic(jClass : Class<*>, message : String)
}