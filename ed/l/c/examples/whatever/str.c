#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void print_1d_str_array(char **arr, int n)
{
    for (int i = 0; i < n; i++) {
        printf("%s ", arr[i]);
    }
}

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

char** str_split(char *text, char delim) {
    char **result = malloc(1 * sizeof(char*));
    int size = 1;

    char *temp = strtok(text, &delim);
    *result = temp;

    while (temp != NULL) {
        size++;
        result = realloc(result, size * sizeof(char*));
        temp = strtok(NULL, &delim);
        result[size - 1] = temp;
    }

    return result;
}

int main()
{
    // word_on_new_line();
    // str_len();

    // split
    char str[20] = "Hello world";
    char **arr = str_split(str, ' ');
    print_1d_str_array(arr, 2);

    return 0;
}
