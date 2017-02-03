import okhttp3.HttpUrl;
import okhttp3.mockwebserver.MockResponse;
import okhttp3.mockwebserver.MockWebServer;
import org.junit.Assert;
import org.junit.Test;



public class GetIPTest {

    @Test
    public void run() throws Exception {
        MockWebServer server = new MockWebServer();

        // Schedule some responses.
        server.enqueue(new MockResponse().setBody("{\"ip\":\"179.35.167.131\"}\n"));
        server.enqueue(new MockResponse().setBody("{\"ip\":\"179.35.167.250\"}\n").setResponseCode(404));

        server.start();

        HttpUrl baseUrl = server.url("");

        GetIP me = new GetIP(baseUrl.url().toString());

        Assert.assertEquals("179.35.167.131", me.get());
        Assert.assertEquals("", me.get());

        Assert.assertEquals("", me.get());
        server.close();
    }
}