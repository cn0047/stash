#include <stdio.h>

union otoa {
  int Intgr;
  float RealNum;
};

void one() {
  printf("[z] size = %d\n", (int)sizeof(union otoa));
};

int main() {
    one();

    return 0;
}
