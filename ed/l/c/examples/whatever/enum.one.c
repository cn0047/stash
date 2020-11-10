#include <stdio.h>

enum wt {
    TheInt,
    TheChar
};

void one()
{
    enum wt v1 = TheInt;
    enum wt v2 = TheChar;
    printf("[1] v1: %d, v2: %d \n", v1, v2); // [1] v1: 0, v2: 1
}

int main()
{
    one();

    return 0;
}
