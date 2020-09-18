#include <stdio.h>
#include <stdlib.h>

static int vals[] = {4, 5, 6};

int len(int arr[])
{
    return 0;
}

void f1()
{
    size_t nib = sizeof(vals); // unsigned long - in bytes
    size_t length = sizeof(vals) / sizeof(vals[0]);
    printf("vals array size in bytes = %lu and length = %lu \n", nib, length);

    for (int i = 0; i < length; i++) {
        int v = vals[i];
        printf("array element with index %d has value %d \n", i, v);
    }
}

void f2()
{
    int arr<::> = <% 1, 2, 3 %>; // digraphs
    printf("[f2] el 2 = %d\n", arr<:1:>);

    int arr2[] = { 1, 2, 3 };
    printf("[f2] el 2 = %d\n", arr2[1]);
}

void array_with_dynamic_length()
{
    int n = 5;
    int *arr = (int*)malloc(n * sizeof(int));
    for (int i = 0; i < n; i++) {
      printf("\n\t%d", arr[i]);
    }
    free(arr);
}

void digits_count()
{
    int arr[10]; // from 0 to 9
    // Init arr with 0.
    for (int i = 0; i < 10; i++) {
      arr[i] = 0;
    }

    // String to find digits in.
    char s[100] = "a11472o5t6";

    int i = 0;
    while (s[i] > 0) {
        if (s[i] >= 48 && s[i] <= 57) { // 48 - digit 0, 57 - digit 9
            char c[2];                  // temp string to hold only 1 char
            sprintf(c, "%c", s[i]);     // char -> str
            int v = atoi(c);            // str -> int
            arr[v]++;
        }
        i++;
    }

    for (int i = 0; i < 10; i++) {
      printf("\n\t%d", arr[i]);
    }
}

int main()
{
    // f1();
    // f2();
    // array_with_dynamic_length();
    digits_count();

    return 0;
}
