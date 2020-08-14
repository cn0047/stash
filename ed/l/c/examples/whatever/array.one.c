#include <stdio.h>

static int vals[] = {4, 5, 6};

void f1()
{
    size_t nib = sizeof(vals); // unsigned long - in bytes
    size_t length = sizeof(vals) / sizeof(vals[0]);
    printf("vals array size in bytes = %lu and length = %lu \n", nib, length);

    for (int i = 0; i < length; i++) {
        int v = vals[i];
        printf("array element with index %d has value %d \n", i, v);
    }
}

void f2()
{
  int arr<::> = <% 1, 2, 3 %>; // digraphs
  printf("[f2] el 2 = %d\n", arr<:1:>);

  int arr2[] = { 1, 2, 3 };
  printf("[f2] el 2 = %d\n", arr2[1]);
}

int main()
{
    // f1();
    f2();

    return 0;
}
