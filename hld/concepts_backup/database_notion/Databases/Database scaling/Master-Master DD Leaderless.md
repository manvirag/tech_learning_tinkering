# Master-Master DD / Leaderless

![Untitled](Master-Master%20DD%20Leaderless/Untitled.png)

Upon write, the client broadcasts the request to all replicas and waits for a certain number of ACKs. 

Upon read, the client contacts all replicas and waits for some responses.

This approach is called a quorum because the client waits for many responses. 

As we shall see later, how we configure the quorum is of critical importance as it determines the consistency of our database. ( This is called distributed consensus )

Let denote the read number as *r*, the write number as *w*, and the total replica as *T.* For instance, the quorum in the above figure has *r=1, w=2, T=3*. ( Weak Consistency)

why? 

![Untitled](Master-Master%20DD%20Leaderless/Untitled%201.png)

**Conflict Resolution:**

- **Last-Write-Wins (LWW):**
- **Vector Clocks:** A

Doubts ( First read about distributed consensus, still not clear then come here ):

Show some examples of r and w combinations and use cases. 

What if some nodes fail and are now not able to complete quorum? 

What if its temporary failure node can come back after the fix? 

What if it's a permanent failure? 

What if get the conflict in the worst case?