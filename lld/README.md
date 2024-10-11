
#### Design Patterns implementations in golang


#### Before going this world, let's understand few things.

- There are the things
    - Machine Coding.
    - Low Level Design.

- Machine Coding:
    - It is like the DSA problem, where you'll have to write the code in ide and will have to submit it. 
    - But that should be maintain proper format or oops concept and design pattern like when you create real world service.
    - Instead of database , you can use inmemory data.
    - So here the focus is working code with different testcase + it should be proper structure and following solid principles and oops concept.
    - In this you don't get much time to draw diagram etc. just decide the entity at different layer and start coding.
    - Here you don't need to tell too much about you oops class etc -> have to but not that as compare to one with happen with less coding more discussion and class diagram.


    - Template:
        - Write down some thinkable type (golang) and method-> in model -> then its repo and usecase -> optimise it with design patterns and solid principles.
        - controller ( write in main.go and say will go in controller ) -> usecases -> models + modelsrepo in single file. In actual model will contain repo and repo will be different directory .
        - No need to create controller , write all these in main.go
        - Code up and run layer by layer -> like first complete -> user + usermodel +userrepo + userusecase + its main -> then other layers.
        - You can see the code in problem section -> they are mainly focussed in consideration with machine round

- Low Level Design:
    - This is kind of bit undefined, it means we actually want to discuss the implementation details of a , let say a backend service and assuming we have the database connected to it.
    - It can have different part like
    - Database modelling or data modelling : 
        - How we will store the data in database 
        - and its schema and its foreign key, primary key and relation. 
        - Here can assume db as mysql for simplicity. -> 
        - find out entity and their fields and ER diagram. 
        - (ER and Data model is related to db), UML, CLass diagram, actor -> more related to oops.
    - Api designing -> 
        - What will be different types of api , 
        - what will be their input and output and 
        - which type of apis will be etc.
    - Oops design -> 
    - this is nothing but the actual code -> what will be different type of classes you will created in your service , 
    - what will be their attributes and methods , 
    - connection of different classes and their relation (has etc), 
    - UML or sequence diagram (optional), 
    - how will you connect from calling api to fetch data from db and use these code level types and return data and code level type should be optimised with design pattern. 
    - [Understand this basic] How will you map the code level type with database type and how would you structure the code like directory level.
    - Service or repo structure.
        - controller, model , usecase , repository.
        - when we say oops -> its all about the type which are in usecase layer -> 
        - repository have data with respect to database -> convert to for usecase layer
        - that is what we usually discuss in oops ( usecase layer type ).