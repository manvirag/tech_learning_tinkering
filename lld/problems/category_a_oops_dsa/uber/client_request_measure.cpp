/*

→ https://leetcode.com/discuss/post/6557379/uber-interview-concurrency-design-a-clas-ybbr/

I was asked to design a class to measure ongoing requests per clientId

was provided an interface with these 3 methods:

1. processRequest(reqId, clientId)
2. processResponse(reqId, clientId)
3. getOngoingRequests(clientId)


*/