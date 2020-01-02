#include <stdio.h>

int main()
{
    int a = 5;
    int * p = &a;
    printf("a = %d \n", a);
    if (p) {
        printf("p = %p \n", (void *)p);
    }
    printf("a + 2 = %d \n", *p+2);

    return 0;
}
