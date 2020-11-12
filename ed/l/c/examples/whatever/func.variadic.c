#include <stdio.h>
#include <stdarg.h>

void print_int_vals(int n, ...)
{
    va_list args;
    va_start(args, n);
    for (int i = 0; i < n; i++) {
        int v = va_arg(args, int);
        printf("v(%d) = %d \n", i, v);
    }
    va_end(args);
}

int main()
{
    int n = 3;
    print_int_vals(n, 5, 6, 7);

    return 0;
}
