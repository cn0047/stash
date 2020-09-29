import java.util.Map;
import java.util.HashMap;

public class map1
{
    public static void main(String[] args)
    {
       simple();
    }

    private static void simple()
    {
        Map<Integer,String> map = new HashMap<Integer,String>();
        map.put(1, "a");
        map.put(2, "b");
        map.put(3, "c");

        for (Map.Entry m:map.entrySet()) {
            System.out.print(m.getKey()+":"+m.getValue()+", ");
        }

        System.out.printf("\nby key: %d got: %s\n", 3, map.get(3));
    }
}
