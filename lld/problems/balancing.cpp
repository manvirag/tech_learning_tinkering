#include <iostream>
#include <map>
#include <cmath>
using namespace std;
class Transaction {
    public:
     string from ; 
     string to;
     int amount;
     Transaction() = default;
     Transaction(string from, string to, int amount) {
        this->from = from;
        this->to = to;
        this->amount = amount;
     }
};
int main() {
     vector<Transaction> transactions;
     transactions.push_back(Transaction("A", "B", 50));
     transactions.push_back(Transaction("B", "C", 30));
     transactions.push_back(Transaction("A", "C", 20));

     map<string, int> netBalance;
     for(auto &transaction : transactions) {
        netBalance[transaction.from] -= transaction.amount;
        netBalance[transaction.to] += transaction.amount;
     }
     for(auto &b : netBalance) {
        cout<<b.first<<" "<<b.second<<endl;
     }

     vector<string> positiveBalanceUsers;
     vector<string> negativeBalanceUsers;
     for(auto &b : netBalance) {
        if(b.second > 0) {
            positiveBalanceUsers.push_back(b.first);
        } else {
            negativeBalanceUsers.push_back(b.first);
        }
     }
     vector<Transaction> finalTransactions;
     while(positiveBalanceUsers.size() > 0 && negativeBalanceUsers.size() > 0) {
        string positiveUser = positiveBalanceUsers.back();
        string negativeUser = negativeBalanceUsers.back();
        int amount = min(netBalance[positiveUser], -netBalance[negativeUser]);
        Transaction transaction;
        transaction.to = positiveUser;
        transaction.from = negativeUser;
            
        if(abs(netBalance[negativeUser])  == amount) {
            negativeBalanceUsers.pop_back();
        }
        if (netBalance[positiveUser]  == amount) {
            positiveBalanceUsers.pop_back();
        }
        netBalance[negativeUser] += amount;
        netBalance[positiveUser] -= amount;

        transaction.amount = amount;
        finalTransactions.push_back(transaction);
        
     }

     for(auto &t : finalTransactions) {
        cout<<t.to<<" "<<t.from<<" "<<t.amount<<endl;
     }
     return 0; 
 
}