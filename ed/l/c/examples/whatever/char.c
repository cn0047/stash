#include <stdio.h>

void checkChar()
{
    char gender = 'b';
    printf("[checkChar] is g = %d \n", gender == 'g');
    printf("[checkChar] is b = %d \n", gender == 'b');

    for (int i = 1; i < 10; i+=2) {
        printf("-- %d \n", i);
    }
}

int main()
{
    checkChar();
    return 0;
}
