#include <stdio.h>

static int vals[] = {4, 5, 6};

void one() {
    size_t nib = sizeof(vals); // unsigned long - in bytes
    size_t length = sizeof(vals) / sizeof(vals[0]);
    printf("vals array size in bytes = %lu and length = %lu \n", nib, length);

    for (int i = 0; i < length; i++) {
        int v = vals[i];
        printf("array element with index %d has value %d \n", i, v);
    }
}

int main() {
    one();
    return 0;
}
