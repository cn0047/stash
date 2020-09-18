#include <stdio.h>
#include <stdlib.h>

void word_on_new_line()
{
    char * s = "How is that";
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

int main()
{
    word_on_new_line();

    return 0;
}
