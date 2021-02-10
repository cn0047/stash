import org.json.simple.JSONObject;

public class json
{
    public static void main(String[] args)
    {
        // f1();
        f2();
    }

    public static void f2()
    {
        Map<String, String> d = new HashMap<>();
        d.put("foo", "bar");

        ObjectMapper om = new ObjectMapper();
        String j = om.writeValueAsString(d);

        System.out.println("j-------");
        System.out.println(j);
        System.out.println("j-------");
    }

    public static void f1()
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
