#include<iostream>
using namespace std;

class Singleton{
    static Singleton* instance;
    int dbData;
    // Private constructor to prevent external instantiation
    Singleton() : dbData(0) {}
    public:
    static Singleton* getInstance(){
         if(!instance){
            instance  = new Singleton();
         }
         return instance;
         
    }
    void setData(int value){
        dbData = value;
    }

    int getData(){
        return dbData;
    }
};

Singleton* Singleton::instance = NULL;

int  main(){
    Singleton* temp = Singleton::getInstance();
    Singleton* temp2 = Singleton::getInstance();
    temp->setData(45);
    cout<<temp2->getData()<<endl;
    return 0;
}