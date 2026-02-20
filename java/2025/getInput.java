import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;

class Main {
  public static void main(String[] args) {
    String url = "https://adventofcode.com/2025/day/1/input";
    String cookie = "";
    HttpClient client = HttpClient.newHttpClient();
    HttpRequest request =
        HttpRequest.newBuilder().uri(URI.create(url)).GET().header("Cookie", cookie).build();
    try {
      HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());

      System.out.println("Status code: " + response.statusCode());
      System.out.println("Response body: " + response.body());
    } catch (IOException | InterruptedException e) {
      e.printStackTrace();
    }
  }
}
