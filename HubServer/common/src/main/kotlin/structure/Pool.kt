package devsanso.github.io.HubServer.common.structure

import java.io.Closeable
import java.util.LinkedList
import java.util.Queue
import java.util.concurrent.ArrayBlockingQueue
import java.util.concurrent.locks.ReentrantLock

class Pool<T> constructor(val max : Int, private val gen : () -> T) {

    class PoolElement<T> internal constructor(private val pool: Pool<T>, private val obj: T) : Closeable {
        internal lateinit var cls : Class<*>
        internal var used : Boolean = false
        private val lock : ReentrantLock = ReentrantLock()

        override fun close() {
            lock.lock()
            if(!used) {
                return
            }
            pool.push(this)
            lock.unlock()
        }

        fun runElement(block : (T) -> Unit) {
            lock.lock()

            if(!used) {
                lock.unlock()
                throw IllegalAccessException()
            }

            try {
                block(obj)
            }catch(e : Exception) {
                pool.removeAllocLog(this)
                throw e
            }finally {
                lock.unlock()
            }

        }
    }

    private val objects : Queue<PoolElement<T>> = LinkedList<PoolElement<T>>()
    val allocSet : HashSet<Class<*>> = HashSet()
    var allocObjectCnt : Int = 0

    @Synchronized
    fun pop(jClass : Class<*>) : PoolElement<T>? {
        if(objects.size > 0) {
            val o = objects.peek()

            if(o.used) {
                throw ConcurrentModificationException()
            }

            o.cls = jClass
            o.used = true
            return o
        }else if(allocObjectCnt < max) {
            val n = gen()
            val o = PoolElement<T>(this, n).also {
                it.cls = jClass
            }
            allocSet.add(jClass)
            allocObjectCnt++
            return o
        }else {
            return null
        }
    }

    @Synchronized
    internal fun removeAllocLog(obj : PoolElement<T>) {
        obj.used = false
        allocSet.remove(obj.cls)
        allocObjectCnt--
    }

    @Synchronized
    internal fun push(obj : PoolElement<T>) {
        obj.used = false
        allocSet.remove(obj.cls)
        objects.add(obj)
    }
}