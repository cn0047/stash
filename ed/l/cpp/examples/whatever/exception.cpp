#include <iostream>

void one()
{
  try {
    throw 404;
  } catch (int errc) {
    std::cout << "Error: " << errc << "\n";
  }
}

int main()
{
  one();

  return 0;
}
