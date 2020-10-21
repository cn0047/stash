import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class list
{
  public static void main(String[] args)
  {
    // print();
    // inList();
    // loop();
    get();
  }

  public static void get()
  {
    List<String> l = getList();
    System.out.printf("get 0 = %s \n", l.get(0));
  }

  public static void loop()
  {
    List<String> l = getList();
    for (String el : l) {
      System.out.printf("%s \n", el);
    }
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
