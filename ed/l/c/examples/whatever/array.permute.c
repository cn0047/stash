#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char** get_array_1()
{
    char **arr;
    int n = 3;

    arr = calloc(n, sizeof(char*));

    arr[0] = calloc(11, sizeof(char));
    arr[0] = "ab";
    arr[1] = calloc(11, sizeof(char));
    arr[1] = "bc";
    arr[2] = calloc(11, sizeof(char));
    arr[2] = "cd";

    return arr;
}

void print_array(char **arr, int n)
{
    for (int i = 0; i < n; i++) {
        printf("%s ", arr[i]);
    }
}

void swap(char **arr, int a, int b)
{
    char *tmp;
    tmp = arr[a];
    arr[a] = arr[b];
    arr[b] = tmp;
}

void reverse(char **arr, int x, int y)
{
    while (x < y) {
        swap(arr, x, y);
        x++;
        y--;
    }
}

void print_permutations(char **arr, int l, int r)
{
    if (l == r) {
        print_array(arr, 3);
        printf("\n");
    } else {
        for (int i = l; i <= r; i++) {
            swap(arr, l, i);
            print_permutations(arr, l+1, r);
            swap(arr, l, i);
        }
    }
}

const int ARR_LEN = 100;
char *p_arr[ARR_LEN][ARR_LEN];
int p_arr_len = 0;

void permute_array(char **arr, int n, int l, int r)
{
    if (l == r) {
        for (int i = 0; i < n; i++) {
            p_arr[p_arr_len][i] = arr[i];
        }
        p_arr_len++;
    } else {
        for (int i = l; i <= r; i++) {
            swap(arr, l, i);
            permute_array(arr, n, l+1, r);
            swap(arr, l, i);
        }
    }
}

void permute_and_print_array(char **arr, int n)
{
    permute_array(arr, n, 0, n-1);
    for (int i = 0; i < p_arr_len; i++) {
        print_array(p_arr[i], n);
        printf("\n");
    }
}

// int p_arr_i = 0; // in regular case
int p_arr_i = 1;    // for hackerrank

int already_permuted = 0;

int next_permutation_1(int n, char **arr)
{
    if (already_permuted == 0) {
        permute_array(arr, n, 0, n-1);
        already_permuted = 1;
    }

    if (p_arr_i >= p_arr_len) {
        return 0;
    }

    // copy
    for (int i = 0; i < n; i++) {
        arr[i] = p_arr[p_arr_i][i];
    }

    p_arr_i++;

    return 1;
}

// ‼️ permutation.
// On each return 1 this func makes array (arr) to contain unique permutation.
int next_permutation_2(int n, char **arr)
{
    int inv = -1;

    for (int i = 0; i < n - 1; i++) {
        if(strcmp(arr[i],arr[i+1]) < 0) {
            inv = i;
        }
    }

    if (inv == -1) {
        return 0;
    }

    for (int i = n - 1; i > inv; i--) {
        if (strcmp(arr[inv], arr[i]) < 0) {
            swap(arr, inv, i);
            break;
        }
    }

    reverse(arr, inv+1, n-1);

    return 1;
}

int next_permutation(int n, char **arr)
{
    next_permutation_1(n, arr);
    next_permutation_2(n, arr);
}

void permute_loop(char **arr, int n)
{
    do {
        for (int i = 0; i < n; i++) {
            printf("%s%c", arr[i], i == n - 1 ? '\n' : ' ');
        }
    } while (next_permutation(n, arr));
}

int main()
{
    int n = 3;
    char **arr = get_array_1();

    // Test print.
    // print_array(arr, n);
    // printf("\n");

    // Test swap.
    // swap(arr, 0, 2);
    // print_array(arr, 3);
    // printf("\n");

    // Test reverse.
    print_array(arr, n);
    printf("\n");
    reverse(arr, 0, 2);
    print_array(arr, n);
    printf("\n");

    // Test permute func.
    // print_permutations(arr, 0, n-1);

    // permute_and_print_array(arr, n);

    // permute_loop(arr, n);

    return 0;
}
