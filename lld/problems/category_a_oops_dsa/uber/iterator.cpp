// implementing simplest 

#include<iostream> 
using namespace std; 

class Iterator {
    public: 
        virtual ~Iterator() {}
        virtual bool hasNext() = 0;
        virtual int next() = 0;
};

class ListIterator: public Iterator {
    public: 
        vector<int> list; 
        int currentIndex;
    ListIterator(vector<int> givenList): list(givenList), currentIndex(-1) {}

    bool hasNext() override {
        // for(auto x: list){
        //     cout<<x<<" ";
        // }
        // cout<<endl;
        int diff = (list.size()-1)-currentIndex;
        if(diff>=1)
         return true;
        return false; 
    }

    int next() override{
        if(hasNext()) {
            currentIndex = currentIndex+1;
            
            return list[currentIndex];
        } else {
            return -1;
        }
    }

};
int main() {
    ListIterator li({1,2,3,4,5});
    
    while(li.hasNext()) {
        
        cout<<li.next()<<" ";
    }
    cout<<endl;
    return 0; 
}