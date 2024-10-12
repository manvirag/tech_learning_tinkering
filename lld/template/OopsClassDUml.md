Once we have relational modeling. We can take hint with it and draw the object oriented design.

https://ujjwalbhardwaj.me/post/low-level-design-design-an-atm-system/
Template
- All class will map to that relational model.
- Those which we are pointing as foreign key , instead of that we will have its object.
- When ever we have different value for a particular column, we can use the interface. (rule condition).
- Apart from this , there are different design pattern which are not very specific to attribute but can use in constructure and methods.
- if not making relational model 
    - Ask the doubts and clear the requirements.
    - Figure our actors and some functionality with usecase diagram. (Not compulsory). https://akhileshmj.medium.com/lld-2-uml-diagrams-8d80e9b41256
    - ![alt text](image.png) 
    - Extend => different types (interface), includes one functionality calling other for completion.
    - Figure out Entities -> Can use top down approach. (Make this in consideration with relational as well , mean feasible in database and efficient as well)
    - Make class diagram with attribute and function.
    ![alt text](image-2.png)
    ![alt text](image-3.png)
    ![alt text](image-4.png)
    ![alt text](image-5.png)
    ![alt text](image-6.png)
    - Make interface (two heading, once has name interface in <<>>) and abstract class (only ittalic)
    - figure our relation
    - is a relation => inheritance.
    ![alt text](image-7.png)
    - has a relation -> Association 
    ![alt text](image-8.png)
        - Composition => existance depend on other.
        - Aggregation => existance not depend on other.
    - Find the cardinality relation
    - 1:m, m:1, 1:1, m:m
    ![alt text](image-1.png)
    - Combine all and draw class diagram
    ![alt text](image-9.png)
    - put these classes in model directory.
    - create repo directory, usecase, controller.
    - controller -> usecase -> usecase/repo 
    - repo fetch data from db and convert into models and return those model to usecase. Kind of object betweem usecase and repo.
    - usecase do business logic etc -> return that model to controller.
    - controller change that structure to dto and return to client.
    - Yeah that's pretty much
    - Now go to problem sets and do the lld design of them.
    



