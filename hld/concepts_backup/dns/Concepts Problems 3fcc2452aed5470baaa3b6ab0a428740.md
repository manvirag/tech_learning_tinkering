# Concepts | Problems

**Book:** System Design Interview: An Insider’s Guide by Alex Xu.

### **DNS(Domain Name System):**

→ If I am creating a website from scratch, what I have to do with the DNS server? 

![Untitled](Concepts%20Problems%203fcc2452aed5470baaa3b6ab0a428740/Untitled.png)

![Untitled](Concepts%20Problems%203fcc2452aed5470baaa3b6ab0a428740/Untitled%201.png)

![Untitled](Concepts%20Problems%203fcc2452aed5470baaa3b6ab0a428740/Untitled%202.png)

  1. browser checks in cache, not → operating system cache , not → operating system call dns resolver.

1. we can configure custom DNS resolve in operator system like cloudfare or google. Else it depends on the DNS resolver server provider by isp.
2. isp provider these configurations by dhcp protocol.
3. DNS resolvers know about the root name servers through a process of initial configuration and continuous updates. Here's how this works
**a.** When setting up DNS resolvers, operators manually configure the IP addresses of the root name servers, which are well-known and stable.
**b.** DNS resolvers use a "root hints" file to store the IP addresses of root name servers. This file is regularly updated by resolver operators to reflect changes in the root server infrastructure. 
**c.** **ICANN and IANA** organizations provide information about the IP addresses of the root name servers to resolver operators and keep them informed about any changes or updates to the root server infrastructure.
**d.** Anycast routing is a technique used for root name servers in DNS, where multiple servers share the same IP address. It improves performance and reduces latency by directing queries to the nearest server based on network proximity.
**e.** **Root Zone Change Notifications:** DNS resolvers may subscribe to notification services provided by ICANN and IANA to stay updated about changes in the root zone, such as the addition of new top-level domains (TLDs) or changes in root server IP addresses. This ensures that resolvers are aware of the latest information regarding the root zone
4. How does tld server know about authoritative server, what is amazon route s3 , what is go daddy what is domain registrar.

       → **going with this hypothesis (not 100% sure → yes correct by chat gpt)** → there are some organizations that are called authoritative servers.  there are some organisations which called domain registrar which have informatiy7on about which domain occupied which are not . These are connected to tld server as well as authoritative servers. when any one register domain, they connect the authoritative server for this and add the information into tld that this domain is occupied and what’s its authoritative server.

1. You register a domain name with a domain registrar.
2. The registrar adds your domain to its authoritative name servers, storing the relevant DNS records.
3. The registrar communicates with the TLD registry responsible for the TLD of your domain (e.g., Verisign for .com domains). They update the TLD registry's database to indicate that your domain is registered and provide information about the authoritative name servers for your domain.

### Write all terminology and its explanation:

→ Consistency:   at any given point in time inf to all users should be the same.
→ Throughput:  total no of operations a system can handle a any point in time ( no of requests /s)
→ Latency: It is the time from the request is made and gets a response.
→ Scalability: the ability to handle more load when increasing the resources.
→ Efficiency: the ability to accomplish good results using fewer resources. for e.g. less tc algo
→ Performance: efficiency is more concerned with resource optimization, while performance looks at the overall user experience, including speed and responsiveness
→ Replication: Creating copies of data or components to improve availability or performance. (x,x,x in 3 servers)
→ Distributed:  in replication we have multiple copies, in this, we divide them into parts and put them separately. (x/3 in 3 servers) 
→ Availability: The proportion of time a system is operational and accessible. high available → 99.99% available to handle requests.
→ Single point of failure: A component whose failure can lead to the entire system's failure.
→ Reliability: Ensures a system consistently operates correctly and is available, even when components fail.
→ Partition Tolerance: Allows a system to keep working during network disruptions, even if it means temporary data inconsistencies.
→ fault tolerance: Ensures a system can continue operating despite component failures within the system.

### **what is** **Distributed Consensus?**

[https://www.youtube.com/watch?v=uS19mAa_tFA&t=274s](https://www.youtube.com/watch?v=uS19mAa_tFA&t=274s)

### Deep dive public and private IP and VPN.

### what is distributed consensus.

[Databases](https://www.notion.so/Databases-d6bd866cf1cf49f79f783b7332fbb152?pvs=21)