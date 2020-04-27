#include <iostream>

int x()
{
  return 7;
}

int x(int v)
{
  return v + x();
}

int main()
{
  int v = x(3);
  std::cout << "got: " << v << std::endl;

  return 0;
}
