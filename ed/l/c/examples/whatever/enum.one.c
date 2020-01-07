#include <stdio.h>

enum wt {
  TheInt,
  TheChar
};

void one() {
  enum wt v = TheInt;
  printf("[z] v = %d\n", v);
};

int main() {
    one();

    return 0;
}
