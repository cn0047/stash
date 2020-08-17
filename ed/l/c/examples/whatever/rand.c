#include <stdio.h>
#include <stdlib.h>
#include <time.h>

void f1()
{
    srand((unsigned)time(NULL)); // init
    int r = rand();
    printf("[f1] %d | %d \n", r, r % 10000); // [f1] 660723882 | 3882
}

int main()
{
    f1();

    return 0;
}
