# Monolith → Microservice

## Monolith:

![Untitled](Monolith%20%E2%86%92%20Microservice%209d17e126ca7c4e3e9cf82fdddcdef1d7/Untitled.png)

### Advantages

The following are some advantages of monoliths:

- Fast and reliable communication.
- Supports ACID transactions.

### Disadvantages

Some common disadvantages of monoliths are:

- Requires commitment to a particular technology stack.
- On each update, the entire application is redeployed.
- Reduced reliability as a single bug can bring down the entire system.
- Difficult to scale or adopt new technologies.

## Modular Monolith:

A Modular Monolith is an approach where we build and deploy a single application (that's the *Monolith* part), but we build it in a way that breaks up the code into independent modules for each of the features needed in our application.

This approach reduces the dependencies of a module in such as way that we can enhance or change a module without affecting other modules. When done right, this can be really beneficial in the long term as it reduces the complexity that comes with maintaining a monolith as the system grows. [ **ui-container** ]

## **Service-oriented architecture (SOA)**:

Not a monolith, also not distributed like microservice its in middle.

## Microservices

Characteristics

The microservices architecture style has the following characteristics:

- **Loosely coupled**:
- **Small but focused**:
- **Highly maintainable**: etc

Demerits:

- Complexity of a distributed system.
- Expensive to maintain (individual servers, databases, etc.).
- Inter-service communication has its own challenges.
- Data consistency.

**Beware of the distributed monolith**

Distributed Monolith is a system that resembles the microservices architecture but is tightly coupled within itself like a monolithic application

Our microservices are just a distributed monolith if any of these apply to it:

- Services don't scale easily.
- Dependency between services.
- Sharing the same resources such as databases.
- Tightly coupled systems.

One of the primary reasons to build an application using microservices architecture is to have scalability. Therefore, microservices should have loosely coupled services which enable every service to be independent. The distributed monolith architecture takes this away and causes most components to depend on one another, increasing design complexit

![Untitled](Monolith%20%E2%86%92%20Microservice%209d17e126ca7c4e3e9cf82fdddcdef1d7/Untitled%201.png)

## Why you don't need microservices

![https://raw.githubusercontent.com/karanpratapsingh/portfolio/master/public/static/courses/system-design/chapter-III/monoliths-microservices/architecture-range.png](https://raw.githubusercontent.com/karanpratapsingh/portfolio/master/public/static/courses/system-design/chapter-III/monoliths-microservices/architecture-range.png)

Before making the decision to move to microservices architecture, you need to ask yourself questions like:

- *"Is the team too large to work effectively on a shared codebase?"*
- *"Are teams blocked on other teams?"*
- *"Does microservices deliver clear business value for us?"*
- *"Is my business mature enough to use microservices?"*
- *"Is our current architecture limiting us with communication overhead?"*

If your application does not require to be broken down into microservices, you don't need this. There is no absolute necessity that all applications should be broken down into microservices.

We frequently draw inspiration from companies such as Netflix and their use of microservices, but we overlook the fact that we are not Netflix. They went through a lot of iterations and models before they had a market-ready solution, and this architecture became acceptable for them when they identified and solved the problem they were trying to tackle.

That's why it's essential to understand in-depth if your business *actually* needs microservices. What I'm trying to say is microservices are solutions to complex concerns and if your business doesn't have complex issues, you don't need them.