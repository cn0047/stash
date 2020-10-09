import org.json.simple.JSONObject;

public class json
{
    public static void main(String[] args)
    {
        JSONObject obj = new JSONObject();
        obj.put("name", "foo");
        obj.put("age", 100);
        String j = obj.toJSONString();

        ObjectMapper mapper = new ObjectMapper();
        Map<String, String> map = mapper.readValue(j, Map.class);

        System.out.println("-------");
        System.out.println(j);
        System.out.println("-------");
        System.out.println(map);
    }
}
