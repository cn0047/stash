#include <stdio.h>
#include <string.h>
#include <math.h>
#include <stdlib.h>

int int_to_str_and_sum_over_loop(int n)
{
    int l = 5; // count of chars in number string

    char s[l+1];
    sprintf(s, "%d", n); // int -> str

    int res = 0;
    for(int i = 0; i <= l; i++) {
        char c[2];              // temp string to hold only 1 char
        sprintf(c, "%c", s[i]); // char -> str
        int v = atoi(c);        // str -> int
        res += v;
    }

    return res;
}

int main() {
    int res = int_to_str_and_sum_over_loop(10564);
    printf("\nsum of digits = %d\n", res);

    return 0;
}
