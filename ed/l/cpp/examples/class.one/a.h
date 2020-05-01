#include <vector>
#include <string>
#include "t.h"

class Account
{
private:
  int balance;
  std::vector<Transaction> log;
  int limit;
public:
  Account();
  std::vector<std::string> Report();
  bool Deposit(int amt);
  bool Withdraw(int amt);
  int GetBalance(){return balance;} // inline function
};

