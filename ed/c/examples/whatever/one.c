#include <stdio.h>

static int z = 3;
static int vals[] = {4, 5, 6};

void hello(int y)
{
  int x = 1;
  printf("hellow world: %d, %d, %d \n", x, y, z);
}

int main()
{
  hello(2);
  return 0;
}
