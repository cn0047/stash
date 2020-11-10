#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void word_on_new_line()
{
    char *s = "How is that";
    int i = 0;
    while (s[i] > 0) {
        if (s[i] == 32) {
            printf("\n");
        } else {
            printf("%c", s[i]);
        }
        i++;
    }
}

void str_len()
{
    char *s = "foo";
    printf("str_len = %zu \n", strlen(s));
}

int main()
{
    // word_on_new_line();
    str_len();

    return 0;
}
