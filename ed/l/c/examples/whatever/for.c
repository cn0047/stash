#include <stdio.h>
#include <string.h>
#include <math.h>
#include <stdlib.h>

void print_names_for_range()
{
    char *list[10];
    list[1] = "one";
    list[2] = "two";
    list[3] = "three";
    list[4] = "four";
    list[5] = "five";
    list[6] = "six";
    list[7] = "seven";
    list[8] = "eight";
    list[9] = "nine";

    int a = 8, b = 11;
    for(int i = a; i <= b; i++) {
        if (i <= 9) {
            printf("%s\n", list[i]);
        } else {
            if (i%2 == 0) {
                printf("even\n");
            } else {
                printf("odd\n");
            }
        }
    }
}

int main()
{
    print_names_for_range();
    return 0;
}
