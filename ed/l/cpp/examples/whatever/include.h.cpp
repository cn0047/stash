/*
@example:
d=ed/l/cpp/examples/whatever
g++ -w -o x $d/1.lib.cpp $d/include.h.cpp && ./x
*/

#include <iostream>

#include "1.lib.h"

int main()
{
  std::cout << "PI = " << getPi() << std::endl;

  return 0;
}
