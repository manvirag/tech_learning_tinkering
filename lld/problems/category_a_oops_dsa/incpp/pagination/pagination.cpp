/*

Transaction Search & Pagination System

You are given an in-memory list of transactions.
Each transaction has the following fields:

id (unique integer)

time (integer timestamp)

userId (integer)

amount (integer, can be negative)

currency (integer)

Part 1 – Generic Filtering

Design a generic search API that allows filtering transactions based on multiple conditions.

Each filter contains:

Field name (id, time, userId, amount, currency)

Operator (=, <, >, <=, >=)

Value

Filters are combined using AND logic.

Part 2 – Pagination

The dataset can be very large and should not be returned all at once.

Implement two pagination strategies:
1️⃣ Offset-Based Pagination

Input:

offset

limit

Output:

Subset of filtered transactions

2️⃣ Cursor-Based Pagination

Input:

cursor (last seen transaction ID)

limit

Output:

Subset of filtered transactions

nextCursor if more data exists

Part 3 – Discussion (Conceptual)

Why pagination is needed

Difference between offset-based and cursor-based pagination

Trade-offs of each approach

How to choose a cursor column
*/



#include <iostream>
#include <vector>
#include <string>

using namespace std;

/* =========================
   Transaction Class
   ========================= */
class Transaction {
public:
    int id;
    int time;
    int userId;
    int amount;
    int currency;

    Transaction(int id, int time, int userId, int amount, int currency)
        : id(id), time(time), userId(userId), amount(amount), currency(currency) {}
};

/* =========================
   Filter Operator Enum
   ========================= */
enum class Op {
    EQ, LT, GT, LE, GE
};

/* =========================
   Filter Class
   ========================= */
class Filter {
public:
    string field;
    Op op;
    int value;

    Filter(const string& field, Op op, int value)
        : field(field), op(op), value(value) {}
};

/* =========================
   Offset Pagination Result
   ========================= */
class OffsetPage {
public:
    vector<Transaction> data;
    int totalCount;

    OffsetPage() : totalCount(0) {}
};

/* =========================
   Cursor Pagination Result
   ========================= */
class CursorPage {
public:
    vector<Transaction> data;
    int nextCursor;   // -1 means no more pages

    CursorPage() : nextCursor(-1) {}
};

/* =========================
   Apply Single Filter
   ========================= */
class FilterEngine {
public:
    static bool applyFilter(const Transaction& tx, const Filter& f) {
        int fieldValue;

        if (f.field == "id") fieldValue = tx.id;
        else if (f.field == "time") fieldValue = tx.time;
        else if (f.field == "userId") fieldValue = tx.userId;
        else if (f.field == "amount") fieldValue = tx.amount;
        else if (f.field == "currency") fieldValue = tx.currency;
        else return false;

        switch (f.op) {
            case Op::EQ: return fieldValue == f.value;
            case Op::LT: return fieldValue < f.value;
            case Op::GT: return fieldValue > f.value;
            case Op::LE: return fieldValue <= f.value;
            case Op::GE: return fieldValue >= f.value;
        }
        return false;
    }

    static bool matchesAll(const Transaction& tx, const vector<Filter>& filters) {
        for (int i = 0; i < filters.size(); i++) {
            if (!applyFilter(tx, filters[i]))
                return false;
        }
        return true;
    }
};

/* =========================
   Query Engine
   ========================= */
class TransactionQueryEngine {
public:
    static OffsetPage queryWithOffset(
        const vector<Transaction>& transactions,
        const vector<Filter>& filters,
        int offset,
        int limit
    ) {
        OffsetPage result;
        vector<Transaction> filtered;

        for (int i = 0; i < transactions.size(); i++) {
            if (FilterEngine::matchesAll(transactions[i], filters)) {
                filtered.push_back(transactions[i]);
            }
        }

        result.totalCount = filtered.size();

        for (int i = offset; i < filtered.size() && result.data.size() < limit; i++) {
            result.data.push_back(filtered[i]);
        }

        return result;
    }

    static CursorPage queryWithCursor(
        const vector<Transaction>& transactions,
        const vector<Filter>& filters,
        int cursorId,   // -1 means start
        int limit
    ) {
        CursorPage result;
        int count = 0;

        for (int i = 0; i < transactions.size(); i++) {
            const Transaction& tx = transactions[i];

            if (cursorId != -1 && tx.id <= cursorId)
                continue;

            if (!FilterEngine::matchesAll(tx, filters))
                continue;

            result.data.push_back(tx);
            count++;

            if (count == limit) {
                result.nextCursor = tx.id;
                break;
            }
        }

        return result;
    }
};

/* =========================
   Demo / Main
   ========================= */
int main() {
    vector<Transaction> transactions;
    transactions.push_back(Transaction(1,11,1,10,1));
    transactions.push_back(Transaction(2,12,1,11,3));
    transactions.push_back(Transaction(3,13,2,-10,1));
    transactions.push_back(Transaction(4,14,1,12,2));
    transactions.push_back(Transaction(5,5,1,10,1));
    transactions.push_back(Transaction(6,6,1,13,1));
    transactions.push_back(Transaction(7,7,1,-5,2));
    transactions.push_back(Transaction(8,8,1,10,2));
    transactions.push_back(Transaction(9,9,1,10,1));
    transactions.push_back(Transaction(10,10,1,15,1));
    transactions.push_back(Transaction(11,21,1,16,1));
    transactions.push_back(Transaction(12,22,1,-3,2));
    transactions.push_back(Transaction(13,23,1,5,1));
    transactions.push_back(Transaction(14,24,1,6,2));
    transactions.push_back(Transaction(15,25,1,10,1));

    vector<Filter> filters;
    filters.push_back(Filter("userId", Op::EQ, 1));
    filters.push_back(Filter("amount", Op::GT, 10));

    cout << "=== Offset-Based Pagination ===\n";
    OffsetPage offsetPage =
        TransactionQueryEngine::queryWithOffset(transactions, filters, 0, 3);

    for (int i = 0; i < offsetPage.data.size(); i++) {
        cout << "id=" << offsetPage.data[i].id
             << " amount=" << offsetPage.data[i].amount << "\n";
    }
    cout << "Total matching: " << offsetPage.totalCount << "\n\n";

    cout << "=== Cursor-Based Pagination ===\n";
    int cursor = -1;

    while (true) {
        CursorPage page =
            TransactionQueryEngine::queryWithCursor(transactions, filters, cursor, 3);

        for (int i = 0; i < page.data.size(); i++) {
            cout << "id=" << page.data[i].id
                 << " amount=" << page.data[i].amount << "\n";
        }

        if (page.nextCursor == -1)
            break;

        cout << "--- NEXT PAGE ---\n";
        cursor = page.nextCursor;
    }

    return 0;
}
