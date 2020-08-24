public class str
{
  public static void main(String[] args)
  {
    // f1();
    f2();
  }

  private static void f2()
  {
    System.out.println(String.format("[f2] %,d", 1204032)); // [f2] 1,204,032
    System.out.println(String.format("[f2] % d", 1));       // [f2]  1
    System.out.println(String.format("[f2] % d", -2));      // [f2] -2
    System.out.println(String.format("[f2] %+d", 3));       // [f2] +3
  }

  private static void f1()
  {
    String s1 = "foo";
    String s2 = "foo";
    if (s1.intern() == s2.intern()) System.out.println("equals");
  }
}
