public class class1
{
  public static void main(String[] args)
  {
    // f1();
    staticField();
  }

  private static void staticField()
  {
    staticFieldClass c = new staticFieldClass();
    c.incCount();
    System.out.println(c.getCount());
  }

  private static void f1()
  {
    System.out.println("It works!");
  }

  private static void f1(String str)
  {
  }
}

public class staticFieldClass
{
  private static int count;

  public staticFieldClass()
  {
    count = 1;
  }

  public int getCount()
  {
    return count;
  }

  public void incCount()
  {
    count += 1;
  }
}
