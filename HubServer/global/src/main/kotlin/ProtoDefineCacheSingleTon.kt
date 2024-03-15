package devsanso.github.io.HubServer.global

import java.util.HashMap

import com.google.protobuf.Descriptors.FieldDescriptor

import devsanso.github.io.HubServer.protos.agent.redis.ClientList.RedisClientInfo
import devsanso.github.io.HubServer.protos.agent.redis.InfoCpu.RedisCpuInfo
import devsanso.github.io.HubServer.protos.agent.redis.InfoStat.RedisStatsInfo
import devsanso.github.io.HubServer.protos.agent.redis.InfoMemory.RedisMemoryInfo
import devsanso.github.io.HubServer.protos.agent.redis.DbSizeOuterClass.DbSize

typealias Type = String
typealias Name = String

object ProtoDefineCacheSingleTon {
    class FieldNameCache internal constructor(private val list : List<FieldDescriptor>){
        private val cache : HashMap<Int, Pair<Type, Name>> = HashMap()
        init {
            list.forEach {
                cache[it.index] = Pair(it.type.name, it.name)
            }
        }
        
        operator fun get(index : Int) : Pair<Type, Name> = cache[index]!!
        fun size() : Int = cache.size
    }

    val RedisCache : HashMap<Type, FieldNameCache> by lazy {
        val ret: HashMap<Type, FieldNameCache> = HashMap()
        ret["RedisClientInfo"] = FieldNameCache(RedisClientInfo.getDescriptor().fields)
        ret["RedisCpuInfo"] = FieldNameCache(RedisCpuInfo.getDescriptor().fields)
        ret["RedisStatsInfo"] = FieldNameCache(RedisStatsInfo.getDescriptor().fields)
        ret["RedisMemoryInfo"] = FieldNameCache(RedisMemoryInfo.getDescriptor().fields)
        ret["DbSize"] = FieldNameCache(DbSize.getDescriptor().fields)
        ret
    }



}