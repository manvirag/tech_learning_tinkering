#include <iostream> 
#include <map>
#include <vector>
using namespace std;

/*

rule stratedy -> evaluate(Expense)
expense class 
facade 

*/
/*
"expenseId": "1",
  "itemId": "Item1",
  "expenseType": "Food",
  "amountInUsd": 250,
  "sellerType": "restaurant",
  "sellerName": "ABC restaurant"
*/
/*


1. **Total expense** should not be more than **175 USD**.
2. **Restaurant-type sellers** should not have expenses more than **45 USD**.
3. **Entertainment-type expenses** are **not allowed**.


*/

enum ExpenseType {
    FOOD,
    ENTERTAINMENT,
    OTHER
};
enum SellerType {
    RESTAURANT,
    HOTEL
};

class Expense {
    public:
    string expenseId;
    string itemId;
    ExpenseType expenseType;
    int amountInUsd;
    SellerType sellerType;
    string sellerName;
}; 
class Rule {
    public:
    virtual string name() = 0;
    virtual bool evaluate(Expense expense) = 0;
};
class TotalExpenseRule : public Rule {
    public:
    bool evaluate(Expense expense){
        return expense.amountInUsd <= 175;
    }
    string name(){
        return "TotalExpenseRule";
    }
};
class RestaurantExpenseRule : public Rule {
    public:
    bool evaluate(Expense expense){
        return expense.sellerType == RESTAURANT && expense.amountInUsd <= 45;
    }
    string name(){
        return "RestaurantExpenseRule";
    }
};
class EntertainmentExpenseRule : public Rule {
    public:
    bool evaluate(Expense expense){
        return expense.expenseType == ENTERTAINMENT;
    }
    string name(){
        return "EntertainmentExpenseRule";
    }
};


class ExpenseFacade {
    public:
    vector<string> evaluateRules(vector<Rule*> rules, vector<Expense> expenses){
        map<string, bool> result;
        for(Rule* rule : rules) {
            for(Expense expense : expenses) {
                if(!rule->evaluate(expense)) {
                    result[rule->name()] = true;
                    
                }
            }
        }
        vector<string> finalResult;
        for(auto &r : result) {
            if(r.second) {
                finalResult.push_back(r.first);
            }
        }
        return finalResult;
    }
};
int main() {
    Expense expense;
    expense.expenseType = ENTERTAINMENT;
    expense.sellerType = RESTAURANT;
    expense.amountInUsd = 250;
    expense.sellerName = "ABC restaurant";
    expense.expenseId = "1";
    expense.itemId = "Item1";


    // Expense expense2;
    // expense2.expenseType = FOOD;
    // expense2.sellerType = RESTAURANT;
    // expense2.amountInUsd = 250;
    // expense2.sellerName = "ABC restaurant";
    // expense2.expenseId = "2";
    // expense2.itemId = "Item2";
    ExpenseFacade expenseFacade;
    vector<Rule*> rules = {new TotalExpenseRule(), new RestaurantExpenseRule(), new EntertainmentExpenseRule()};
    vector<Expense> expenses = {expense};
    vector<string> result = expenseFacade.evaluateRules(rules, expenses);
    for(auto &r : result) {
        cout<<r<<endl;
    }
    cout<<"Hello World"<<endl;
    return 0;
}