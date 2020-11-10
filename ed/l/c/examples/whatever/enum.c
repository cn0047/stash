#include <stdio.h>

enum {
    X_VAL = 100
};

void f1()
{
    printf("[f1] %d \n", X_VAL); // [f1] 100
}

int main()
{
    f1();

    return 0;
}
