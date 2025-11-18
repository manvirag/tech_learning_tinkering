#include <iostream>
#include <map>
#include <vector>
using namespace std;

class PortManager {
    private:
        vector<int> freeStack;
        vector<bool> isFree;
    
    public:
        PortManager(int N) : isFree(N, true) {
            freeStack.reserve(N);
            // Initialize with all ports as free
            for (int i = 0; i < N; i++) {
                freeStack.push_back(i);
            }
        }
    
        // O(1)
        int get() {
            if (freeStack.empty()) return -1;  // No free ports available
            int port = freeStack.back();
            freeStack.pop_back();
            isFree[port] = false;
            return port;
        }
    
        // O(1)
        void free(int port_id) {
            // Ignore if out-of-range OR already free
            if (port_id < 0 || port_id >= (int)isFree.size()) return;
            if (isFree[port_id]) return;
    
            isFree[port_id] = true;
            freeStack.push_back(port_id);
        }
    };
    
int main() {
   
    return 0;
}