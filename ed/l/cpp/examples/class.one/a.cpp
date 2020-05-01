#include "a.h"

using namespace std;

Account::Account():balance(0),limit(100)
{
}

vector<string> Account::Report()
{
  vector<string> report;
  report.push_back("Balance is " + to_string(balance));
  report.push_back("Transactions: ");
  for (auto t:log)
  {
    report.push_back(t.Report());
  }
  report.push_back("------------------");

  return report;
}

bool Account::Deposit(int amt)
{
  if (amt >= 0) {
    balance += amt;
    log.push_back(Transaction(amt,"Deposit"));
    balance -= 1;
    log.push_back(Transaction(1,"Service Charge"));
    return true;
  } else {
    return false;
  }
}

bool Account::Withdraw(int amt)
{
  if (amt >= 0) {
    if (balance + limit  >= amt) {
      balance -= amt;
      log.push_back(Transaction(amt,"Withdraw"));
      balance -= 1;
      log.push_back(Transaction(1,"Service Charge"));
      return true;
    } else {
      return false;
    }
  } else {
    return false;
  }
}
