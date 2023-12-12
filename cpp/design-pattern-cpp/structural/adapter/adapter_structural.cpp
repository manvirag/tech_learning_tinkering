#include <iostream>
#include <string>
using namespace std;

// LegacyDatabase interface
class LegacyDatabase {
public:
    virtual void saveRecord(string key, string value) = 0;
    virtual string loadRecord(string key) = 0;
};

// Standard Database interface
class Database {
public:
    virtual void insert(string key, string value) = 0;
    virtual string select(string key) = 0;
};

// Adapter to convert LegacyDatabase to Database
class DatabaseAdapter : public Database {
private:
    LegacyDatabase* legacyDatabase;

public:
    DatabaseAdapter(LegacyDatabase* legacyDatabase) : legacyDatabase(legacyDatabase) {}

    void insert(string key, string value) override {
        legacyDatabase->saveRecord(key, value);
    }

    string select(string key) override {
        return legacyDatabase->loadRecord(key);
    }
};

// LegacyDatabase implementation
class LegacyDatabaseImpl : public LegacyDatabase {
public:
    void saveRecord(string key, string value) override {
        cout << "LegacyDatabaseImpl: Saving record with key " << key << " and value " << value << endl;
    }

/*
    The override keyword is used after the function signature (std::string loadRecord(std::string key)) to indicate that this function is intended to override a virtual function from a base class. It's a way of indicating to the compiler that you expect this function to match a virtual function declared in the base class exactly in terms of function name, parameters, and return type.
*/
    string loadRecord(string key) override {
        return "Loaded value for key " + key;
    }
};

int main() {
    LegacyDatabaseImpl legacyDatabase;
    DatabaseAdapter databaseAdapter(&legacyDatabase);

    databaseAdapter.insert("foo", "bar");
    string value = databaseAdapter.select("foo");
    cout << value << endl;

    return 0;
}
