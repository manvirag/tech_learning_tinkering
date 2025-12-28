#include <iostream>
#include <map>
#include <string>
#include <unordered_set>
#include <set>
using namespace std;
/*
assumption : 
- we do have valid raw. 
- no cycle right now. 

- we can have dependency 
- circular ? -> yes ignore that update 
- can referent single cell multiple times ? -> yes 
- extendable to multiple operator currently + and - only
- rawvalue - > number, or formula or dependency
- can mix of both number and dependency -> a1+7+a2-5
*/
class Cell {
    public:
      string key;
      string rawValue;
      long int calculatedValue;
      set<pair<string, long int> > dependencies;
      set<pair<string, long int> > dependents;
      Cell() {
        rawValue = "";
        calculatedValue = 0;
      }
};
class ExpressionParser {
    public:
      ExpressionParser() = default;

      bool isFormula(string value) {
        return value[0] == '=';
      }
      set<pair<string, long int>> getDependencies(string value) {
        set<pair<string, long int>> dependencies;
        if(!isFormula(value)) return dependencies;
        map<string, long int> coefficients;
        string formula = value.substr(1);
        for(int i = 0; i < formula.size(); i++) {
            if(formula[i] >= 'A' && formula[i] <= 'Z') {
                string currentKey = ""; 
                while(i < formula.size() && formula[i] != '+' && formula[i] != '-') {
                    currentKey += formula[i];
                    i++;
                }
                cout<<"currentKey: "<<currentKey<<endl;
                coefficients[currentKey]++;
            }
        }

        for(auto &coefficient : coefficients) {
            dependencies.insert({coefficient.first, coefficient.second});
        }
        return dependencies;
      }
      
      long int calculateValue(string value, map<string, Cell> cells) {
        
        if(!isFormula(value)) return stoi(value);
        string formula = value.substr(1);
        long int result = 0;
        bool isPositive = true;
        string previousToken = "";
        for(int i = 0; i < formula.size();) {
            if(formula[i] == '+' || formula[i] == '-') {
                isPositive = formula[i] == '+';
                i++;
            } else if(formula[i] >= 'A' && formula[i] <= 'Z') {
                string currentKey = "";
                while(i < formula.size() && formula[i] != '+' && formula[i] != '-') {
                    currentKey += formula[i];
                    i++;
                }
                cout<<"currentKey: "<<currentKey<<endl;
                result += isPositive ? cells[currentKey].calculatedValue : -cells[currentKey].calculatedValue;

            } else if(formula[i] >= '0' && formula[i] <= '9') {
                string currentNumber = "";
                while(i < formula.size() && formula[i] >= '0' && formula[i] <= '9') {
                    currentNumber += formula[i];
                    i++;
                }
                cout<<"currentNumber: "<<currentNumber<<endl;
                result += isPositive ? stoi(currentNumber) : -stoi(currentNumber);
            }
        }
        return result;
      }
    
};
class ExcelSheet {
    public: 
      map<string, Cell> cells;
      ExpressionParser parser;
      ExcelSheet() {
        parser = ExpressionParser();
        cells.clear();
      }
      void print() {
        for(auto &cell : cells) {
            cout<<cell.first<<" "<<cell.second.rawValue<<" "<<cell.second.calculatedValue<<endl;
        }
      }
      void set(string cell, string value) {
        if(cells.find(cell) == cells.end()) {
            cells[cell] = Cell();
        }
        cells[cell].rawValue = value;
        cells[cell].key = cell;
        for(auto dependent : cells[cell].dependencies) {
            cells[dependent.first].dependents.erase({cell, dependent.second});
        }
        cells[cell].dependencies.clear();
        long int previousValue = cells[cell].calculatedValue;
        cells[cell].dependencies = parser.getDependencies(value);
        for(auto dependency : cells[cell].dependencies) {
            cells[dependency.first].dependents.insert({cell, dependency.second});
        }
        cells[cell].calculatedValue = parser.calculateValue(value, cells);
        long int delta = cells[cell].calculatedValue - previousValue;
        for(auto dependent : cells[cell].dependents) {
            updateDependent(dependent.first, dependent.second , delta);
        }
      }
      void reset(string cell) {
        cells[cell].rawValue = "";
        
        for(auto dependent : cells[cell].dependencies) {
            cells[dependent.first].dependents.erase({cell, dependent.second});
        }
        cells[cell].dependencies.clear();
        for(auto dependent : cells[cell].dependents) {
            updateDependent(dependent.first, dependent.second, -cells[cell].calculatedValue);
        }
        cells[cell].calculatedValue = 0;
      }
      void updateDependent(string cell, long int coefficient, long int value) {
        long int initialValue = cells[cell].calculatedValue;
        cells[cell].calculatedValue += coefficient * value;
        for( wauto dependent : cells[cell].dependents) {
            updateDependent(dependent.first, dependent.second, value - initialValue);
        }
      }     
};
int main() {

    ExcelSheet sheet = ExcelSheet();
    sheet.set("A1", "1");
    sheet.set("A2", "2");
    sheet.set("A3", "=3+6");
    sheet.set("A4", "=A1+A2+A3");
    sheet.set("A2", "3");
    sheet.set("A1", "3");
    sheet.reset("A3");
    sheet.print();
    return 0;
}

/*

excelsheet 
   expression 
    - number 
    - formula  : =7+3-10 
    - dependency : A1+A2+A3+A2 
      - cycle in dependency -> ignore that update 
    - mix of both number and dependency : 1+A1+A2+A3+A2 
    
cell 




*/