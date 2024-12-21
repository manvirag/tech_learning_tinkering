# Multi-master in db

![Untitled](Multi-master%20in%20db/Untitled.png)

So in Multi-Leader replication, **we are going to have one leader in each of my data centers** and each data center’s leader replicates its changes to the leaders in other data centers **asynchronously**. ***Each Leader also act as a Follower to other Leaders***

Advantages of Multi-Leader Replication

Better Performance as compared to Single leader replication as we have now reduced both Read & Write Latency of our application

High Fault Tolerance as each data center can continue operating independently of others if any data center goes down. Also, replication catches up when the failed datacenter comes back online (Might be by promoting a follower as a leader and asking for the latest changes from the leaders of other data centers).**Yes, there is a trade-off between Performance and High Availability here.**

Issues (**Write Conflicts**.):

![Untitled](Multi-master%20in%20db/Untitled%201.png)

if you have previously worked with **GIT** as a **Version Control System** then you can also connect this scenario similar to **merge conflicts**

Solution ? [Need to check this more]

**Conflict Detection**

In the **Multi-Leader replication** scheme, when two leaders in different data centers have conflicts then those write **conflicts can be detected asynchronously at some later point in time without blocking any user**

But at that time, it may be too late to ask the user to resolve the conflict. So this approach isn’t widely accepted.

**Conflict Avoidance (Recommended approach)**

according to this technique if the **application can ensure that all writes for a particular record go through the same leader (data center), then conflicts cannot occur.**

Probably that datacenter is in his geographical location or is the nearest datacenter to his geographical location.

what if a user has moved to a different location Yes,**in this situation, the conflict avoidance technique will break**

**Conclusion**

So as we always know [**There is no such thing as a perfect system](http://www.jeff-hester.com/reality-straight-up/saving-capitalism-thing-perfect-system/),** but yes there are good and bad systems.

But we know [**Prevention is always better than cure**](https://dictionary.cambridge.org/dictionary/english/prevention-is-better-than-cure). So here **Prevention** corresponds to **Avoidance** and **Cure** corresponds to a situation when conflict has already occurred and we need to eliminate it by **Detecting** it**.**

**Hence we can conclude that Conflict Avoidance is better than Conflict Detection.**

Sample Use-case of this technique in google calender.

We might want to use it even though we are not connected to the internet (storing in local storage)

This is similar to the Multi-Leader replication scheme where each device acts as a Leader since it accepts the reads and writes made by its clients. There is an asynchronous multi-leader replication scheme between the replicas of our calendar present in all of our devices.

![Untitled](Multi-master%20in%20db/Untitled%202.png)