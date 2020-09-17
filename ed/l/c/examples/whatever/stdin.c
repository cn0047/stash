#include <stdio.h>
#include <string.h>
#include <math.h>
#include <stdlib.h>

void scan_and_print_sum_and_diff()
{
    int i1, i2;
    float fv1, fv2;

    scanf("%d %d", &i1, &i2);
    scanf("%f %f", &fv1, &fv2);

    printf("%d %d\n", i1+i2, i1-i2);
    printf("%.1f %.1f\n", fv1+fv2, fv1-fv2);
}

void f1()
{
    char s[100];
    scanf("%[^\n]%*c", &s);
    printf("got:%s", s);
}

int main()
{
    // f1();
    scan_and_print_sum_and_diff();
}
