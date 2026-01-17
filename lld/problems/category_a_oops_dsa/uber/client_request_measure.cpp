/*

→ https://leetcode.com/discuss/post/6557379/uber-interview-concurrency-design-a-clas-ybbr/

I was asked to design a class to measure ongoing requests per clientId

was provided an interface with these 3 methods:

1. processRequest(reqId, clientId)
2. processResponse(reqId, clientId)
3. getOngoingRequests(clientId) -> reqId -> which didn't have processResponse fhn called ?? -> yes 


*/

#include<iostream> 
#include<map> 
#include<string> 
#include<unordered_set>
#include<shared_mutex>
using namespace std; 


class OngoingRequestsManagementTool {
    public: 
        map<string, unordered_set<string>> mp; // for simpliciy assume client and request as string 
        shared_mutex mu;
        

    OngoingRequestsManagementTool() {}

    void processRequest(string reqId, string clientId) {
        unique_lock<shared_mutex> lock(mu);
        mp[clientId].insert(reqId);
    }
    void processResponse(string reqId, string clientId) {
        unique_lock<shared_mutex> lock(mu);
        mp[clientId].erase(reqId);
    }
    int getOngoingRequests(string clientId ) {
        shared_lock<shared_mutex> lock(mu);
        return mp[clientId].size();
    }
};
int main() {
    OngoingRequestsManagementTool tool; 
    tool.processRequest("1", "a");
    tool.processRequest("2", "a");
    tool.processRequest("3", "a");
    tool.processRequest("4", "b");
    cout<<tool.getOngoingRequests("a")<<endl;
    tool.processResponse("1", "a");
    cout<<tool.getOngoingRequests("a")<<endl;
    cout<<tool.getOngoingRequests("b")<<endl;

    cout<<"hello world"<<endl;
    return 0;
}

