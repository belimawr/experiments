package demo

import com.sun.net.httpserver.HttpExchange
import com.sun.net.httpserver.HttpHandler
import com.sun.net.httpserver.HttpServer
import khttp.get
import org.junit.After
import org.junit.Test
import java.io.IOException
import java.net.InetSocketAddress
import java.security.SecureRandom
import kotlin.test.assertEquals


class TestMyHTTPClient {

    val server: HttpServer = HttpServer.create()
    val serverAddr: String

    init {
        val random = SecureRandom()
        var tries = 0
        var port = 8000 + random.nextInt(20000)

        while (tries < 2000) {
            tries ++
            try {
                server.bind(InetSocketAddress(port), 0)
                break
            } catch (e: Exception) {
                port = 8000 + random.nextInt(20000)
            }
        }

        server.executor = null
        server.start()

        serverAddr = "http://localhost:" + server.address.port
    }

    @After
    fun tearDown() {
        server.stop(0)
    }

    @Test fun shouldMakeTwoHTTPCalls() {
        val counter = Count(0)
        server.createContext("/shouldMakeTwoHTTPCalls", MockHTTPHandler(counter))

        val r = get(serverAddr + "/shouldMakeTwoHTTPCalls")
        val r2 = get(serverAddr + "/shouldMakeTwoHTTPCalls")

        assertEquals(2, counter.testCount)

        assertEquals(200, r.statusCode)
        assertEquals("This is the response", r.text)

        assertEquals(200, r2.statusCode)
    }

    @Test fun shouldMakeOneHTTPCalls() {
        val counter = Count(0)
        server.createContext("/test", MockHTTPHandler(counter))

        val r = get(serverAddr + "/test")
        assertEquals(200, r.statusCode)
        assertEquals("This is the response", r.text)

        val r2 = get(serverAddr + "/some/other/path")
        assertEquals(404, r2.statusCode)

        assertEquals(1, counter.testCount)
    }

    internal class MockHTTPHandler(val counter: Count) : HttpHandler {

        @Throws(IOException::class)
        override fun handle(t: HttpExchange) {
            val response = "This is the response"
            t.sendResponseHeaders(200, response.length.toLong())
            val os = t.responseBody
            os.write(response.toByteArray())
            os.close()
            counter.testCount++
        }
    }

    data class Count(var testCount: Int)
}
