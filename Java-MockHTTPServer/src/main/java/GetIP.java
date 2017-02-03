import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;
import org.json.simple.JSONObject;
import org.json.simple.parser.JSONParser;
import org.json.simple.parser.ParseException;

import java.io.IOException;

public class GetIP {
    private OkHttpClient client;
    private JSONParser parser;
    private String url;

    public GetIP(String url) {
        parser = new JSONParser();
        client = new OkHttpClient();

        this.url = url;
    }

    String get() {
        Request request = new Request.Builder()
                .url(this.url + "/json")
                .build();

        Response response;
        String jsonData;
        try {
            response = client.newCall(request).execute();
            if (response.code() != 200) {
                return "";
            }
            jsonData = response.body().string();
        } catch (IOException e) {
            return "";
        }

        JSONObject jsonObject;
        try {
            jsonObject = (JSONObject) parser.parse(jsonData);

            if (jsonObject.containsKey("ip")) {
                return (String) jsonObject.get("ip");
            }
        } catch (ParseException e) {
            return "";
        }
        return "";
    }
}