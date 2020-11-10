#include <stdio.h>

struct hello {
    int index;
    float code;
};

void one()
{
    struct hello h = {5, 200};
    printf("[1] index = %d; code = %0.1f \n", h.index, h.code);
};

typedef unsigned char byte;

struct Color
{
    byte Red;
    byte Green;
    byte Blue;
};

// typedef struct {
//   short First;
//   short Second;
// } Layout;

void two()
{
    struct Color c = {225, 128};
    printf("[2] %d \n", c.Green);

    // struct Layout l = {1, 2};
    // printf("[2] %d \n", l.First);
};

int main()
{
    one();
    two();

    return 0;
}
