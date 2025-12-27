#include <iostream>
#include <vector>
#include <queue>
using namespace std;

/* =========================
   Common Iterator Interface
   ========================= */
class Iterator {
public:
    virtual bool hasNext() = 0;
    virtual int next() = 0;
    virtual ~Iterator() {}
};

/* =========================
   1. Basic List Iterator
   ========================= */
class ListIterator : public Iterator {
    vector<int> data;
    int index;

public:
    ListIterator(const vector<int>& nums) : data(nums), index(0) {}

    bool hasNext() override {
        return index < data.size();
    }

    int next() override {
        return data[index++];
    }
};

/* =========================
   2. Range Iterator
   Supports positive & negative step
   ========================= */
class RangeIterator : public Iterator {
    int current;
    int end;
    int step;
    bool valid;

public:
    RangeIterator(int start, int end, int step)
        : current(start), end(end), step(step) {

        if (step == 0) valid = false;
        else if (start < end && step < 0) valid = false;
        else if (start > end && step > 0) valid = false;
        else valid = true;
    }

    bool hasNext() override {
        if (!valid) return false;
        if (step > 0) return current <= end;
        return current >= end;
    }

    int next() override {
        int val = current;
        current += step;
        return val;
    }
};

/* =========================
   3. Interleaving Iterator
   (Works on iterator objects)
   ========================= */
class InterleavingIterator : public Iterator {
    queue<Iterator*> q;

public:
    InterleavingIterator(const vector<Iterator*>& iterators) {
        for (auto it : iterators) {
            if (it && it->hasNext())
                q.push(it);
        }
    }

    bool hasNext() override {
        return !q.empty();
    }

    int next() override {
        Iterator* it = q.front();
        q.pop();

        int val = it->next();

        if (it->hasNext())
            q.push(it);

        return val;
    }
};

/* =========================
   4. Interleaving Iterator
   (List of Lists version)
   ========================= */
class InterleavingListIterator {
    vector<vector<int>> lists;
    queue<pair<int,int>> q;

public:
    InterleavingListIterator(const vector<vector<int>>& input)
        : lists(input) {

        for (int i = 0; i < lists.size(); i++) {
            if (!lists[i].empty())
                q.push({i, 0});
        }
    }

    bool hasNext() {
        return !q.empty();
    }

    int next() {
        auto [i, j] = q.front();
        q.pop();

        int val = lists[i][j];

        if (j + 1 < lists[i].size())
            q.push({i, j + 1});

        return val;
    }
};

/* =========================
   5. Filtered Iterator
   ========================= */
class FilteredIterator : public Iterator {
    Iterator* base;
    function<bool(int)> predicate;
    bool hasCached;
    int cachedValue;

public:
    FilteredIterator(Iterator* it, function<bool(int)> pred)
        : base(it), predicate(pred), hasCached(false) {}

    bool hasNext() override {
        if (hasCached) return true;

        while (base->hasNext()) {
            int val = base->next();
            if (predicate(val)) {
                cachedValue = val;
                hasCached = true;
                return true;
            }
        }
        return false;
    }

    int next() override {
        hasNext();
        hasCached = false;
        return cachedValue;
    }
};

/* =========================
   MAIN — Sample Usage
   ========================= */
int main() {
    cout << "=== List Iterator ===\n";
    ListIterator it1({1,2,3});
    while (it1.hasNext())
        cout << it1.next() << " ";
    cout << "\n\n";

    cout << "=== Range Iterator ===\n";
    RangeIterator it2(0, 10, 2);
    while (it2.hasNext())
        cout << it2.next() << " ";
    cout << "\n\n";

    cout << "=== Interleaving Lists ===\n";
    vector<vector<int>> lists = {{1,2,3},{4,5},{6}, {}, {7,8,9}};
    InterleavingListIterator il(lists);
    while (il.hasNext())
        cout << il.next() << " ";
    cout << "\n\n";

    cout << "=== Interleaving Iterators ===\n";
    ListIterator a({1,2,3});
    ListIterator b({4,5,6});
    RangeIterator c(1,6,1);

    vector<Iterator*> its = {&a, &b, &c};
    InterleavingIterator zigzag(its);

    while (zigzag.hasNext())
        cout << zigzag.next() << " ";
    cout << "\n\n";

    cout << "=== Filtered Iterator (Even Numbers) ===\n";
    RangeIterator r(1, 10, 1);
    FilteredIterator even(&r, [](int x) { return x % 2 == 0; });

    while (even.hasNext())
        cout << even.next() << " ";
    cout << "\n";

    return 0;
}
