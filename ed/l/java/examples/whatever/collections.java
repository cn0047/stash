import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.HashMap;
import java.util.Map;
import java.util.Set;
import java.util.HashSet;

public class collections
{
  public static void main(String[] args)
  {
    // list1();
    // map1();
    set1();
  }

  public static void list1()
  {
    List<String> list = Arrays.asList("foo", "bar");
    // list.add("x");
    list.forEach(System.out::println);
  }

  public static void map1()
  {
    Map<String, String> map = new HashMap<>();
    map.put("f", "foo");
    map.put("b", "bar");
    map.put("x", "x");
    map.remove("x");
    map.forEach((k, v) -> System.out.printf("%s %s%n", k, v));
  }

  public static void set1()
  {
    Set<String> hs = new HashSet<String>();
    hs.add("foo");
    hs.add("bar");
    System.out.println(hs);
  }
}
