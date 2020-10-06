import org.json.simple.JSONObject;

public class json
{
  public static void main(String[] args)
  {
      JSONObject obj = new JSONObject();
      obj.put("name", "foo");
      obj.put("age", 100);
      System.out.println(obj.toJSONString());
  }
}
