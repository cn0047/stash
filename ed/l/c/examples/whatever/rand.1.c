#include <stdio.h>
#include <stdlib.h>
#include <time.h>

enum { ARR_LEN = 100 };

void f1()
{
    int i, *pNumbers = malloc(ARR_LEN * sizeof(int));
    if (pNumbers == NULL) {
        fprintf(stderr, "Insufficient memory.\n");
        exit(1);
    }

    srand((unsigned)time(NULL));

    printf("\n%d random numbers between 0 and 9999:\n", ARR_LEN);
    for (i = 0; i < ARR_LEN; ++i) {
        pNumbers[i] = rand() % 10000;
        // print
        printf("%6d", pNumbers[i]);
        if (i % 10 == 9) putchar('\n');
    }

    free(pNumbers);
}

int main()
{
    f1();

    return 0;
}
