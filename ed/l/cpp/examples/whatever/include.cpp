/*
@example:
d=ed/l/cpp/examples/whatever
g++ -w -o x $d/1.lib.cpp $d/include.cpp && ./x
*/

#include <iostream>

float getPi();

int main()
{
  std::cout << "PI = " << getPi() << std::endl;

  return 0;
}
