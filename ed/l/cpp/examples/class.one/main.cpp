#include <iostream>
#include "a.h"

using namespace std;

int main()
{
  Account a1;
  a1.Deposit(90);
  cout << "After depositing $90 " <<
    "balance is " << a1.GetBalance() << endl;
  for(auto s:a1.Report()) {
    cout << s << endl;
  }
  a1.Withdraw(50);
  if (a1.Withdraw(100)) {
    cout << "second withdraw succeeds" << endl;
  }
  cout << "After withdrawing $50 then $100" << endl;
  for(auto s:a1.Report()) {
    cout << s << endl;
  }

  return 0;
}
