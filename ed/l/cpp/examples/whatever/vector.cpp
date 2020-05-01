#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

using namespace std;

int one()
{
  vector<int> v;
  for (int i=0; i<10; i++) {
    v.push_back(i);
  }
  for (auto item:v) {
    cout << item << " ";
  }
  cout << endl;

  return 0;
}

int main()
{
  one();

  return 0;
}
