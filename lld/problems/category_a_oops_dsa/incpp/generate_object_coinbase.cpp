/*


You are designing a system to generate objects composed of multiple properties.

Each property has a finite set of possible values.

Part 1 – Random Object Generation

You are given a set of object properties and their possible values.

Example:

Object type: Tree

Properties:

Height → {Short, Medium, Tall}

LeafColor → {Green, Yellow, Red}

TrunkType → {Thin, Thick}

Task:

Generate N random objects, where each object randomly selects one value from each property.

Part 2 – Generate Only Unique Objects

Modify your solution so that no two generated objects are identical.

Two objects are considered identical if all property values match.

If all possible combinations are exhausted, stop generating more objects.

Part 3 – Property Rarity

Enhance the system to support rarity/weight for property values.

Example:

Height:

Short → Common

Medium → Common

Tall → Rare

Rare values should appear less frequently in the generated objects.
*/


#include <iostream>
#include <vector>
#include <string>
#include <map>

using namespace std;

/* =========================
   PropertyValue Class
   ========================= */
class PropertyValue {
public:
    string value;
    int weight;

    PropertyValue(const string& value, int weight = 1)
        : value(value), weight(weight) {}
};

/* =========================
   Property Class
   ========================= */
class Property {
public:
    string name;
    vector<PropertyValue> values;

    Property(const string& name) : name(name) {}

    void addValue(const string& value, int weight = 1) {
        values.push_back(PropertyValue(value, weight));
    }
};

/* =========================
   Object Class
   ========================= */
class TreeObject {
public:
    map<string, string> attributes;

    void print() const {
        for (map<string, string>::const_iterator it = attributes.begin();
             it != attributes.end(); ++it) {
            cout << it->first << "=" << it->second << " ";
        }
        cout << "\n";
    }
};

/* =========================
   Generator Class
   ========================= */
class ObjectGenerator {
private:
    vector<Property> properties;

public:
    ObjectGenerator(const vector<Property>& properties)
        : properties(properties) {}

    /* -------------------------
       1. Random Generation
       ------------------------- */
    TreeObject generateRandom() {
        TreeObject obj;
        for (int i = 0; i < properties.size(); i++) {
            int idx = rand() % properties[i].values.size();
            obj.attributes[properties[i].name] =
                properties[i].values[idx].value;
        }
        return obj;
    }

    /* -------------------------
       2. Unique Generation
       Using Index Enumeration
       ------------------------- */
    vector<TreeObject> generateAllUnique() {
        vector<TreeObject> result;
        vector<int> indices(properties.size(), 0);

        while (true) {
            TreeObject obj;
            for (int i = 0; i < properties.size(); i++) {
                obj.attributes[properties[i].name] =
                    properties[i].values[indices[i]].value;
            }
            result.push_back(obj);

            int pos = properties.size() - 1;
            while (pos >= 0) {
                indices[pos]++;
                if (indices[pos] < properties[pos].values.size())
                    break;
                indices[pos] = 0;
                pos--;
            }
            if (pos < 0)
                break;
        }
        return result;
    }

    /* -------------------------
       3. Weighted Random Generation
       ------------------------- */
    TreeObject generateWeightedRandom() {
        TreeObject obj;

        for (int i = 0; i < properties.size(); i++) {
            int totalWeight = 0;
            for (int j = 0; j < properties[i].values.size(); j++) {
                totalWeight += properties[i].values[j].weight;
            }

            int r = rand() % totalWeight;
            for (int j = 0; j < properties[i].values.size(); j++) {
                if (r < properties[i].values[j].weight) {
                    obj.attributes[properties[i].name] =
                        properties[i].values[j].value;
                    break;
                }
                r -= properties[i].values[j].weight;
            }
        }
        return obj;
    }
};

/* =========================
   Main
   ========================= */
int main() {
    srand(time(0));

    Property height("Height");
    height.addValue("Short", 5);
    height.addValue("Medium", 5);
    height.addValue("Tall", 1);   // rare

    Property leaf("LeafColor");
    leaf.addValue("Green", 5);
    leaf.addValue("Yellow", 3);
    leaf.addValue("Red", 2);

    Property trunk("TrunkType");
    trunk.addValue("Thin", 4);
    trunk.addValue("Thick", 2);

    vector<Property> properties;
    properties.push_back(height);
    properties.push_back(leaf);
    properties.push_back(trunk);

    ObjectGenerator generator(properties);

    cout << "=== Random Generation ===\n";
    for (int i = 0; i < 5; i++) {
        generator.generateRandom().print();
    }

    cout << "\n=== Unique Generation ===\n";
    vector<TreeObject> uniqueObjs = generator.generateAllUnique();
    for (int i = 0; i < uniqueObjs.size(); i++) {
        uniqueObjs[i].print();
    }

    cout << "\n=== Weighted Random Generation ===\n";
    for (int i = 0; i < 10; i++) {
        generator.generateWeightedRandom().print();
    }

    return 0;
}
