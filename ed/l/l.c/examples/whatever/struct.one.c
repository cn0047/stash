#include <stdio.h>

struct hello
{
    int index;
    float code;
};

int main()
{
    struct hello h = {5, 200};
    printf("index = %d; code = %0.1f \n", h.index, h.code);

    return 0;
}
