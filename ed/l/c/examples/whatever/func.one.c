#include <stdio.h>

void f1()
{
    printf("[f1] v1: %s \n", __func__);
}

int main()
{
    f1();

    return 0;
}
