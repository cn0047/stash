#include <stdio.h>
#include <stdlib.h>

// It prints:
//
// n=5, result:
//    5 5 5 5 5 5 5 5 5
//    5 4 4 4 4 4 4 4 5
//    5 4 3 3 3 3 3 4 5
//    5 4 3 2 2 2 3 4 5
//    5 4 3 2 1 2 3 4 5
//    5 4 3 2 2 2 3 4 5
//    5 4 3 3 3 3 3 4 5
//    5 4 4 4 4 4 4 4 5
//    5 5 5 5 5 5 5 5 5
//
// n=3, result:
//    3 3 3 3 3
//    3 2 2 2 3
//    3 2 1 2 3
//    3 2 2 2 3
//    3 3 3 3 3
int main()
{
    int n = 5;
    int m = n + (n - 1);
    int *arr = (int*)malloc(m * m * sizeof(int));
    int i, j = 0;

    // Init.
    for (i = 0; i < m; i++) {
      for (j = 0; j < m; j++) {
        *(arr + i * m + j) = 0;
      }
    }

    // Generate.
    for (i = 0; i < n; i++) {
        int v = n - i;
        int steps = m - i*2;
        printf("v = %d, %d \n", v, steps);
        for (j = 0; j < steps; j++) {
            *(arr + i * m + j + i) = v;             //   ⮕
            *(arr + (i + j) * m + i) = v;           // ⬇
            *(arr + (m - i - 1) * m + j + i) = v;   //   ⮕
            *(arr + (i + j) * m + (m - i - 1)) = v; //     ⬇
        }
    }

    // Print.
    printf("\n");
    for (i = 0; i < m; i++) {
      for (j = 0; j < m; j++) {
        printf("%d ", *(arr + i*m + j));
      }
      printf("\n");
    }

    return 0;
}
