/*
→Design similar to org chart which we have on teams. There were multiple requirement. Like allocate budget. Add manager, add IC, remove manager, remove IC and few more requirements



→  https://leetcode.com/discuss/post/7313588/uber-sse-technical-phone-round-by-get_si-w5km/

1. You are given a list of edges [[a,b], [b,c], [d,f], [f,g]]. Here a,b means a is manager of b. You have to implement 3 functions:a. Return total number of employees under a manager (both direct and indirect)b. Change the manager of an employee. If that employee has some reportees then those will also be affected.c. Add new employees. Eg: [c,d] needs to be added.

Input format was not hardcoded. You have to write the main function on how input will be given to these apis.

Basically if you create a tree structure you would be able to solve this question. The key was to have optimized solution such that (a) part should have TC - O(1)

I couldn't come up with optimized solution in the given time.




→ https://leetcode.com/discuss/post/6880181/uber-sde-2-interview-experience-verdict-8e53t/

Asked to design an Employee & Team Directory Management System with:

- Add employee
- Assign manager
- Add to team
- Hierarchical view from any employee
- Calculate employee and team CTC

Was able to implement the core logic and run it on given test cases.

Due to time limits, couldn’t add all validations and edge-case handling.
*/