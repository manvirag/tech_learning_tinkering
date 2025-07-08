# hld

**Decision Making  | Trade-off | Which to select and why to select.**

**Step:**
→ **Functional Requirements.**

what the software system should do?

They outline the specific functions, features, and capabilities that the software must provide to meet the user's needs

for e.g., we want to make a cache

1. get(key)
2. put(key, value) etc.
3. in short features part comes in this.
→ API designing. CRUD
→ API error handling.
→ Data modeling.
→ Query.

→ **Non-Function Requirements. (Jo ki tume batai nahi jayagi , but understood hota hai)**

These requirements answer questions like "How should the software perform?”

for e.g. performance, scalability, availability, security, etc.

→ load balancing
→ consistent hashing
→ sharding
→ data backup, replication, distribution
→ security firewall vpn etc. authentication + authorisation + access control + rate limiting
→ caching
→ monitoring and alerting.
→ metrics and analytics. 
→ logs analysis.
→ continuous deployment → deployment without downtime.
→ Https using
→ testing before deployment  → load testing + unit test + integration test
→ async messaging system, queuing system
→ DNS domain name system.

→ **Data Estimations.**

The back of the envelope estimation:

→ This is a rough estimation
→ Don’t spend much time, because the interview says scalable. so it mostly will have scalable components like a load balancer etc.
→ Keep number simple 2^x

| Zeros | Traffic | Storage |
| --- | --- | --- |
| 3 | thousand | KB |
| 6 | million | MB |
| 9 | billion | GB |
| 12 | Trillian | TB |
| 15 | quadrillion | PB |

char →  2 bytes

long/double → 8 bytes

image avg → 300kb

trade → cap theorem

…
5 millions user * 1kb → 5 gb storage

**Steps:**

**Traffic Estimations: 
       →** total no of requests per second on service.

 → total no of queries per second on storage.

![Untitled](hld%208da3438652384c5489859d575a2ac5c8/Untitled.png)

DAU → daily active user
MAU → monthly active user.
→ Write request
→ Read request
→ read  heavy or write heavy ? 

![Untitled](hld%208da3438652384c5489859d575a2ac5c8/Untitled%201.png)

**Storage Estimations:
      →** 5-year storage 

![Untitled](hld%208da3438652384c5489859d575a2ac5c8/Untitled%202.png)

**Bandwidth estimations:**

amount of data transfer per second.

![Untitled](hld%208da3438652384c5489859d575a2ac5c8/Untitled%203.png)

**Ram Estimations(for e.g cache):**

→ total ram 

![Untitled](hld%208da3438652384c5489859d575a2ac5c8/Untitled%204.png)

**No of Machines with RAM:**

![Untitled](hld%208da3438652384c5489859d575a2ac5c8/Untitled%205.png)

| Availability (percentage) | Downtime per year |
| --- | --- |
| 99 | 3.6 days |
| 99.99 | 52 minutes |
| 99.999 | 5 minutes |
| 99.9999 | 31 seconds |

Ask the interviewer about the trade-off of CAP 
****

→ **High-level design (API + diagraming)**

****

→ **Deep dive into the design and its components.** 
→ **Identify bottlenecks and scalability.**

### Resource:

 1. https://github.com/donnemartin/system-design-primer

1. [https://leetdesign.com/](https://leetdesign.com/)
2. [https://www.mydistributed.systems/](https://www.mydistributed.systems/)

### Concepts:

1. 

### Design Blogs

1. Uber -> [https://lnkd.in/dFh5V6UW](https://lnkd.in/dFh5V6UW)

2. Pinterest ->[https://lnkd.in/dQWFTwas](https://lnkd.in/dQWFTwas)

3. Netflix -> [https://lnkd.in/dADsZZpJ](https://lnkd.in/dADsZZpJ)

4. Snapchat -> [https://eng.snap.com/blog](https://eng.snap.com/blog)

5. Dropbox -> [https://dropbox.tech/](https://dropbox.tech/)

6. Google -> [https://lnkd.in/dBwy8_G7](https://lnkd.in/dBwy8_G7)

7. Twitter -> [https://lnkd.in/d6RqN-Cq](https://lnkd.in/d6RqN-Cq)

8. Meta -> [https://lnkd.in/dbm-dQWT](https://lnkd.in/dbm-dQWT)

9. Jane Street -> [https://lnkd.in/dV6xDTht](https://lnkd.in/dV6xDTht)

10. Longer List -> [https://lnkd.in/dYniH2rX](https://lnkd.in/dYniH2rX)

1. Airbnb: [https://lnkd.in/dAPjjaA3](https://lnkd.in/dAPjjaA3)
2. Amazon: [https://lnkd.in/dyp43Yqp](https://lnkd.in/dyp43Yqp)
3. Bittorrent: [https://lnkd.in/dfZPa6Ma](https://lnkd.in/dfZPa6Ma)
4. Asana: [https://lnkd.in/dWqZxf6Y](https://lnkd.in/dWqZxf6Y)
5. Atlassian: [https://lnkd.in/d-i34bUQ](https://lnkd.in/d-i34bUQ)
6. Cloudera: [https://blog.cloudera.com](https://blog.cloudera.com/)
7. Docker: [https://blog.docker.com](https://blog.docker.com/)
8. Dropbox: [https://lnkd.in/dUQJTxac](https://lnkd.in/dUQJTxac)
9. eBay: [https://lnkd.in/dnmca2uT](https://lnkd.in/dnmca2uT)
10. Facebook: [https://lnkd.in/dbwkUDjN](https://lnkd.in/dbwkUDjN)
11. GitHub: [https://lnkd.in/dSC9StzD](https://lnkd.in/dSC9StzD)
12. Google: [https://lnkd.in/ddPVy6Zj](https://lnkd.in/ddPVy6Zj)
13. Groupon: [https://lnkd.in/dsyGvUWF](https://lnkd.in/dsyGvUWF)
14. Highscalability: [http://highscalability.com](http://highscalability.com/)
15. Instacart: [https://tech.instacart.com](https://tech.instacart.com/)
16. Instagram: [https://lnkd.in/dEs6FyGn](https://lnkd.in/dEs6FyGn)
17. Linkedin: [https://lnkd.in/d_yQe9g6](https://lnkd.in/d_yQe9g6)
18. Mixpanel: [https://mixpanel.com/blog](https://mixpanel.com/blog)
19. Netflix: [https://lnkd.in/dKhbQqxd](https://lnkd.in/dKhbQqxd)
20. Nextdoor: [https://lnkd.in/dDdGPQgR](https://lnkd.in/dDdGPQgR)
21. PayPal: [https://lnkd.in/d9YkeE_h](https://lnkd.in/d9YkeE_h)
22. Pinterest: [https://lnkd.in/duz8a8vq](https://lnkd.in/duz8a8vq)
23. Quora: [https://lnkd.in/d-iuzYZq](https://lnkd.in/d-iuzYZq)
24. Reddit: [https://redditblog.com](https://redditblog.com/)
25. Salesforce: [https://lnkd.in/dV9unb47](https://lnkd.in/dV9unb47)
26. Shopify: [https://lnkd.in/dQtK4TME](https://lnkd.in/dQtK4TME)
27. Slack: [https://slack.engineering](https://slack.engineering/)
28. Soundcloud: [https://lnkd.in/dgWK_v4h](https://lnkd.in/dgWK_v4h)
29. Spotify: [https://labs.spotify.com](https://labs.spotify.com/)
30. Stripe: [https://lnkd.in/dm-WBTgr](https://lnkd.in/dm-WBTgr)
31. System design primer: [https://lnkd.in/dnUnsQE9](https://lnkd.in/dnUnsQE9)
32. Twitter: [https://lnkd.in/d9tmm5wj](https://lnkd.in/d9tmm5wj)
33. Thumbtack: [https://lnkd.in/d6QTWF_p](https://lnkd.in/d6QTWF_p)
34. Uber: [http://eng.uber.com](http://eng.uber.com/)
35. Yahoo: [https://lnkd.in/dKgyhbNE](https://lnkd.in/dKgyhbNE)
36. Yelp: [https://lnkd.in/d_6hhMS4](https://lnkd.in/d_6hhMS4)
37. Zoom: [https://lnkd.in/dquH3cKY](https://lnkd.in/dquH3cKY)

Mock Interview Playlist:

[https://www.youtube.com/playlist?list=PLrtCHHeadkHp92TyPt1Fj452_VGLipJnL](https://www.youtube.com/playlist?list=PLrtCHHeadkHp92TyPt1Fj452_VGLipJnL)

[Problems:](hld%208da3438652384c5489859d575a2ac5c8/Problems%2067713130bb074f3787cd7e6032521461.md)

[Concepts | Problems](hld%208da3438652384c5489859d575a2ac5c8/Concepts%20Problems%203fcc2452aed5470baaa3b6ab0a428740.md)