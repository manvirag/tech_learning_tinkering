#include <iostream>
#include <vector>
#include <stack>
#include <queue>

// Iterator interface
class Iterator {
public:
    virtual int next() = 0;
    virtual bool hasNext() = 0;
};

// Iterator for an array
class ArrayIterator : public Iterator {
private:
    std::vector<int> array;
    int currentIndex;

public:
    ArrayIterator(const std::vector<int>& arr) : array(arr), currentIndex(0) {}

    int next() override {
        return array[currentIndex++];
    }

    bool hasNext() override {
        return currentIndex < array.size();
    }
};

// Iterator for a stack
class StackIterator : public Iterator {
private:
    std::stack<int> stack;

public:
    StackIterator(const std::stack<int>& st) : stack(st) {}

    int next() override {
        int value = stack.top();
        stack.pop();
        return value;
    }

    bool hasNext() override {
        return !stack.empty();
    }
};

// Iterator for a queue
class QueueIterator : public Iterator {
private:
    std::queue<int> queue;

public:
    QueueIterator(const std::queue<int>& q) : queue(q) {}

    int next() override {
        int value = queue.front();
        queue.pop();
        return value;
    }

    bool hasNext() override {
        return !queue.empty();
    }
};

// Iterator for a binary tree
class TreeNode {
public:
    int val;
    TreeNode* left;
    TreeNode* right;
    TreeNode(int value) : val(value), left(nullptr), right(nullptr) {}
};

class TreeIterator : public Iterator {
private:
    TreeNode* root;
    std::stack<TreeNode*> stack;

public:
    TreeIterator(TreeNode* node) : root(node) {
        if (root != nullptr) {
            stack.push(root);
        }
    }

    int next() override {
        TreeNode* curr = stack.top();
        stack.pop();

        if (curr->right != nullptr) {
            stack.push(curr->right);
        }

        if (curr->left != nullptr) {
            stack.push(curr->left);
        }

        return curr->val;
    }

    bool hasNext() override {
        return !stack.empty();
    }
};

int main() {
    // Example usage
    std::vector<int> numbers = {1, 2, 3, 4, 5};
    Iterator* arrayIterator = new ArrayIterator(numbers);

    while (arrayIterator->hasNext()) {
        std::cout << arrayIterator->next() << " ";
    }
    std::cout << std::endl;

    std::stack<int> stackNumbers;
    stackNumbers.push(5);
    stackNumbers.push(4);
    stackNumbers.push(3);
    Iterator* stackIterator = new StackIterator(stackNumbers);

    while (stackIterator->hasNext()) {
        std::cout << stackIterator->next() << " ";
    }
    std::cout << std::endl;

    std::queue<int> queueNumbers;
    queueNumbers.push(1);
    queueNumbers.push(2);
    queueNumbers.push(3);
    Iterator* queueIterator = new QueueIterator(queueNumbers);

    while (queueIterator->hasNext()) {
        std::cout << queueIterator->next() << " ";
    }
    std::cout << std::endl;

    TreeNode* root = new TreeNode(1);
    root->left = new TreeNode(2);
    root->right = new TreeNode(3);
    root->left->left = new TreeNode(4);
    root->left->right = new TreeNode(5);
    Iterator* treeIterator = new TreeIterator(root);

    while (treeIterator->hasNext()) {
        std::cout << treeIterator->next() << " ";
    }
    std::cout << std::endl;

    delete arrayIterator;
    delete stackIterator;
    delete queueIterator;
    delete treeIterator;

    return 0;
}
