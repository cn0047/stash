import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class list
{
  public static void main(String[] args)
  {
    print();
    inList();
  }

  public static void inList()
  {
    List<String> list = getList();
    System.out.printf("501");
  }

  public static List<String> getList()
  {
    List<String> list = Arrays.asList("foo", "bar");
    return list;
  }

  public static void print()
  {
    List<String> list = getList();
    // list.add("x");
    list.forEach(System.out::println);
  }
}
