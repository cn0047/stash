import java.util.Set;
import java.util.HashSet;

public class set
{
    public static void main(String[] args)
    {
        print();
        inSet();
    }

    public static Set<String> getStrSet()
    {
        Set<String> hs = new HashSet<String>();
        hs.add("foo");
        hs.add("bar");
        return hs;
    }

    public static void inSet()
    {
        Set<String> s = Set.of("foo", "boo");
        System.out.println("\n[inSet]");
        System.out.println(s.contains("boo"));
    }

    public static void print()
    {
        Set<String> hs = getStrSet();
        System.out.println("\n[print]");
        System.out.print(hs);
    }
}
