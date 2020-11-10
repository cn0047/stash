#include <stdio.h>

union otoa {
    int Intgr;
    float RealNum;
};

void one()
{
    printf("[1] size = %d\n", (int)sizeof(union otoa)); // [1] size = 4
};

int main()
{
    one();

    return 0;
}
