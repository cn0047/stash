using System;
using System.Net.Http;
using System.Threading.Tasks;

namespace Main
{
    public class Program
    {
        static async Task Main()
        {
            await SendPostRequest();
            // await SendJSONPostRequest();
        }

        static async Task SendPostRequest()
        {
            string url = "https://webhook.site/1f95ac34-45df-4003-8108-cb3c9868202f";
            var client = new HttpClient();
            try
            {
                var request = new HttpRequestMessage(HttpMethod.Post, url);
                var collection = new List<KeyValuePair<string, string>>();
                collection.Add(new("clientId", "foo"));
                collection.Add(new("secretKey", "bar"));
                var content = new FormUrlEncodedContent(collection);
                request.Content = content;
                var response = await client.SendAsync(request);
                response.EnsureSuccessStatusCode();
                Console.WriteLine(await response.Content.ReadAsStringAsync());
            }
            catch (Exception ex)
            {
                Console.WriteLine($"Exception: {ex.Message}");
            }
        }

        static async Task SendJSONPostRequest()
        {
            string url = "https://webhook.site/1f95ac34-45df-4003-8108-cb3c9868202f";
            using (HttpClient httpClient = new HttpClient())
            {
                try
                {
                    string jsonPayload = "{\"foo\": \"bar\"}";
                    var content = new StringContent(jsonPayload, System.Text.Encoding.UTF8, "application/json");
                    HttpResponseMessage response = await httpClient.PostAsync(url, content);
                    if (response.IsSuccessStatusCode)
                    {
                        Console.WriteLine("POST request sent successfully!");
                    }
                    else
                    {
                        Console.WriteLine($"Error: {response.StatusCode}");
                    }
                    Console.WriteLine(await response.Content.ReadAsStringAsync());
                }
                catch (Exception ex)
                {
                    Console.WriteLine($"Exception: {ex.Message}");
                }
            }
        }
    }
}
