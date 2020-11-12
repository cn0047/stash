#include <stdio.h>

void f1()
{
    char* s = NULL;
    printf("test=%s \n", s);
    // assert(s != NULL); // Undefined symbols for architecture x86_64: "_assert"
}

int main()
{
    f1();

    return 0;
}
