#include <stdio.h>

static int z = 3;

void hello(int y)
{
    int x = 1;
    printf("hellow world: %d, %d, %d \n", x, y, z); // hellow world: 1, 2, 3
}

int main()
{
    hello(2);
    return 0;
}
