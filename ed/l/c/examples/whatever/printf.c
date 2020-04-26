#include <stdio.h>

void one() {
    printf("[1] %-10.2f value \n", 12.34);
    printf("[1] %-10.2d value \n", 123456789);
}

int main() {
    one();

    return 0;
}
