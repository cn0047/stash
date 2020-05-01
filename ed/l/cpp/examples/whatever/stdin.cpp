#include <iostream>
#include <string>

using namespace std;

int one()
{
  int i;
  cout << "Enter value i: ";
  cin >> i;
  cout << i << endl;
}

int two()
{
  string n;
  cout << "What is your name? ";
  cin >> n;
  cout << "Hello " << n << endl;

  return 0;
}

int main()
{
  two();

  return 0;
}
