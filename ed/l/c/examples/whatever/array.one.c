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

void digits_count_in_string()
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

void dynamic_2d_array()
{
    int r = 3, c = 4;
    int *arr = (int*)malloc(r * c * sizeof(int));

    // generate
    int i, j, v = 0;
    for (i = 0; i < r; i++) {
      for (j = 0; j < c; j++) {
         *(arr + i*c + j) = ++v;
      }
    }

    // print
    for (i = 0; i < r; i++) {
      for (j = 0; j < c; j++) {
         printf("%d ", *(arr + i*c + j));
      }
      printf("\n");
    }
}

void print_1d_str_array(char **arr, int n)
{
    for (int i = 0; i < n; i++) {
        printf("%s ", arr[i]);
    }
}

void print_1d_int_array(int *arr, int n)
{
    for (int i = 0; i < n; i++) {
        printf("%d ", arr[i]);
    }
}

void int_1d_array()
{
    int arr[] = {4, 5, 6};
    print_1d_int_array(arr, 3);
}

const int INT_2D_ARRAY_LEN = 100;

void print_2d_int_array(int arr[][INT_2D_ARRAY_LEN], int n, int m)
{
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < m; j++) {
            printf("%d ", arr[i][j]);
        }
        printf("\n");
    }
}

void int_2d_array()
{
    int arr[][INT_2D_ARRAY_LEN] = {
        {90, 91},
        {80, 81, 82},
    };
    print_2d_int_array(arr, 2, 3);
}

void str_1d_array()
{
    int n = 3;
    char *arr[] = {"foo", "bar", "123"};
    print_1d_str_array(arr, n);
}

void str_1d_array_v2()
{
    int n = 3;
    char **arr = {"Learning", "C", "is", "fun"}; // won't work
    print_1d_str_array(arr, n);
}

const int STR_2D_ARRAY_LEN = 100;

void print_2d_str_array(char *arr[][STR_2D_ARRAY_LEN], int n, int m)
{
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < m; j++) {
            printf("%s ", arr[i][j]);
        }
        printf("\n");
    }
}

void str_2d_array()
{
    char *arr[STR_2D_ARRAY_LEN][STR_2D_ARRAY_LEN] = {
        {"one", "foo"},
        {"two", "bar"},
        {"333", "ok"},
    };

    char *v = "this is 4";
    arr[3][0] = "#4";
    arr[3][1] = v;

    arr[4][0] = "#5";
    arr[4][1] = arr[4][0];

    print_2d_str_array(arr, 5, 2);
}

int main()
{
    // f1();
    // f2();
    // array_with_dynamic_length();
    // digits_count_in_string();
    // dynamic_2d_array();
    // str_1d_array();
    // str_2d_array();
    // int_1d_array();
    // int_2d_array();
    // str_2d_array();
    // str_1d_array_v2(); // error

    return 0;
}
