package demo

import com.sun.net.httpserver.HttpExchange
import com.sun.net.httpserver.HttpHandler
import com.sun.net.httpserver.HttpServer
import khttp.get
import org.junit.Test
import java.io.IOException
import java.net.InetSocketAddress
import kotlin.test.assertEquals


class TestMyKttpClient {


//    internal data class Count(var testCount: Int, var count2: Int)

    public var count = 0

    @Test fun shouldMakeHttpCall() {
        val server = HttpServer.create(InetSocketAddress(8888), 0)
        server.createContext("/test", MyHandler(this))
        server.executor = null
        server.start()

        val r = get("http://localhost:8888/test")
        assertEquals(200, r.statusCode)
        assertEquals("This is the response", r.text)

        val r2 = get("http://localhost:8888/notFound")
        assertEquals(404, r2.statusCode)

        assertEquals(1, count)
    }

    internal class MyHandler(t: TestMyKttpClient) : HttpHandler {

        val myt: TestMyKttpClient = t

        @Throws(IOException::class)
        override fun handle(t: HttpExchange) {
            val response = "This is the response"
            t.sendResponseHeaders(200, response.length.toLong())
            val os = t.responseBody
            os.write(response.toByteArray())
            os.close()
            myt.count++
        }
    }

}
