#include <iostream>
#include <map>
#include <string>
#include <unordered_set>
#include <iomanip>
#include <cmath>
using namespace std;


long double currencyConvertion(string from, string to, map<string, vector<pair<string, long double >>> &currencyMap, long double amount, unordered_set<string> &visited) {
    if(from == to) {
        return amount;
    }
    visited.insert(from);
    cout<<from<<" "<<to<<" "<<endl;
    cout<<fixed<<setprecision(10)<<amount<<endl;
    for(auto node: currencyMap[from]) {
        if(visited.find(node.first) != visited.end()) {
            continue;
        }
        cout<<amount<<" "<<node.second<<endl;
        long double result = currencyConvertion(node.first, to, currencyMap, amount * node.second, visited);
        if(result != -1) {
            return result;
        }
    }
    return -1;

}
int main() {
    
    map<string, vector<pair<string, long double >>> currencyMap;
    currencyMap["USD"].push_back(make_pair("JPY", (long double)110));
    currencyMap["JPY"].push_back(make_pair("USD", (long double)1/(long double)110));
    currencyMap["USD"].push_back(make_pair("AUD", (long double)1.45));
    currencyMap["AUD"].push_back(make_pair("USD", (long double)1/(long double)1.45));
    currencyMap["JPY"].push_back(make_pair("GBP", (long double)0.0070));
    currencyMap["GBP"].push_back(make_pair("JPY", (long double)1/(long double)0.0070));
    long double amount = 1;
    unordered_set<string> visited;
    long double result = currencyConvertion("GBP", "AUD", currencyMap, amount, visited);
    result = ceil(result * 100.0) / 100.0;
    cout << fixed << setprecision(2) << result << endl;
    cout << "Hello, World!" << endl;
    return 0;
}