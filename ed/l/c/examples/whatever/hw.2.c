#include <stdio.h>

static int z = 3;

void hello(int y)
{
    unsigned int x = 1u;
    char *str = "hellow world";
    printf("%s: %d, %d, %d \n", str, x, y, z); // hellow world: 1, 2, 3
}

int main()
{
    hello(2);
    return 0;
}
