# network protocol/certificates

1. **IP Public-Private:**
    - **Public IP:** Identifies a device on the internet. It's unique and can be accessed directly from the web.
    - **Private IP:** Used within a private network to identify and communicate with devices within the same network. Not directly accessible from the internet
        - How do we connect then?
2. **OSI Model:**
3. **TCP and UDP:**
    - 
4. **HTTP and HTTPS:**
    - **HTTP (Hypertext Transfer Protocol):** It's the foundation of data communication for the World Wide Web.
    - **HTTPS (Hypertext Transfer Protocol Secure):** It's the secure version of HTTP, encrypted with SSL/TLS.
5. **DNS (Domain Name System):**
    - Translates domain names to IP addresses, allowing users to access websites and other internet resources using easy-to-remember names.
6. **SSL , TLS, mTLS:**
    1. [https://github.com/karanpratapsingh/system-design?tab=readme-ov-file#ssl-tls-mtls](https://github.com/karanpratapsingh/system-design?tab=readme-ov-file#ssl-tls-mtls)
    2. 
        
        Let's briefly discuss some important communication security protocols such as SSL, TLS, and mTLS. I would say that from a *"big picture"* system design perspective, this topic is not very important but still good to know about.
        
        ## SSL
        
        SSL stands for Secure Sockets Layer, and it refers to a protocol for encrypting and securing communications that take place on the internet. It was first developed in 1995 but since has been deprecated in favor of TLS (Transport Layer Security).
        
        ### Why is it called an SSL certificate if it is deprecated?
        
        Most major certificate providers still refer to certificates as SSL certificates, which is why the naming convention persists.
        
        ### Why was SSL so important?
        
        Originally, data on the web was transmitted in plaintext that anyone could read if they intercepted the message. SSL was created to correct this problem and protect user privacy. By encrypting any data that goes between the user and a web server, SSL also stops certain kinds of cyber attacks by preventing attackers from tampering with data in transit.
        
        ## TLS
        
        Transport Layer Security, or TLS, is a widely adopted security protocol designed to facilitate privacy and data security for communications over the internet. TLS evolved from a previous encryption protocol called Secure Sockets Layer (SSL). A primary use case of TLS is encrypting the communication between web applications and servers.
        
        There are three main components to what the TLS protocol accomplishes:
        
        - **Encryption**: hides the data being transferred from third parties.
        - **Authentication**: ensures that the parties exchanging information are who they claim to be.
        - **Integrity**: verifies that the data has not been forged or tampered with.
        
        ## mTLS
        
        Mutual TLS, or mTLS, is a method for mutual authentication. mTLS ensures that the parties at each end of a network connection are who they claim to be by verifying that they both have the correct private key. The information within their respective TLS certificates provides additional verification.
        
        ### Why use mTLS?
        
        mTLS helps ensure that the traffic is secure and trusted in both directions between a client and server. This provides an additional layer of security for users who log in to an organization's network or applications. It also verifies connections with client devices that do not follow a login process, such as Internet of Things (IoT) devices.
        
        Nowadays, mTLS is commonly used by microservices or distributed systems in a [zero trust security model](https://en.wikipedia.org/wiki/Zero_trust_security_model) to verify each other.
        
7. **Subnet:**
    - A logical subdivision of an IP network. It's used to divide a single, large network into smaller, more manageable networks.
8. **VPC (Virtual Private Cloud):**
    - A virtual network dedicated to your AWS account. It's logically isolated from other virtual networks in the AWS cloud.
9. **VPN (Virtual Private Network):**
    - It extends a private network across a public network and enables users to send and receive data across shared or public networks as if their computing devices were directly connected to the private network.
10. **SSH (Secure Shell) Authentication: [ ? ]**
    1. SSH is short for ‘secure shell’. It is a protocol for sharing data between two computers over the internet. Instead of sharing password to service , without encrypted
    2. Client generate 2 keys public and private. Public can be shared and it share it to server.
    3. Now client send data to server, service encrypt it with public key and sen