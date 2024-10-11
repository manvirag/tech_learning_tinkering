Template:

1. Write down the scoped requirements after discussing.
2. Write down actors and make use diagram (what actor's actions).
3. Write down some thinkable entity/class name only.
4. Write down class in top down mapper (no necessary but its easy to write from big to small).
5. Write variables , methods etc and show relationship with other class like interface , composition etc.
6. At the time of 4 and 5 or after them , think where we can optimize with the help of design patterns . Use them. To make design more scalable according to the requirements.
7. Write the code with the help of class diagram (4,5, and 6). More detailed your 4,5 and 6 points would be like all methods , variables, dependency etc. It would be fast and easy to write the  actual codes.
8. Follow MVC architecture -> Controller -> domain -> use-case -> repository -> client/db 
9. Run code.


For 8.

```
class {
    variables
    methods  -> db
}

model -> types only -> {variables}
usecase(business logic) -> methods  -> (repositories)
repository(no business logic get data from db/client return to usecase)  -> db/client 
```
Above would work if no code , but when need running code.




One is Maching Coding Round -> Expectation Code to be running and need fast implementation, we don't get time to draw things right. 
Template ( Atleast Worked for me ):

1. controller ( write in main.go and say will go in controller ) -> usecases -> models + modelsrepo in single file. In actual model will contain repo and repo will be different directory .
2. Write down some thinkable type (golang) -> in model -> then its repo and usecase -> optimise it with design patterns and solid principles.
3. No need to create controller , write all these in main.go
4. Code up and run layer by layer -> like first complete -> user + usermodel +userrepo + userusecase + its main -> then other layers.




Second is Low Level Design -> Expectation -> It mostly depend upon interviewr ( assume you are going to create new service exposed as api and connected with database like useg.[ignore other] )

- Database modelling or data modelling : How we will store the data in database and its schema and its foreign key, primary key and relation. Here can assume db as mysql for simplicity.  -> find out entity and their fields and ER diagram. (ER and Data model is related to db, UML, CLass diagram, actor -> more related to oops.)
- Api designing -> What will be different types of api , what will be their input and output and which type of apis will be etc.
- Oops design  -> this is nothing but the actual code -> what will be different type of classes you will created in your service , what will be their attributes and methods , connection of different classes and their relation (has etc), UML or sequence diagram (optional), how will you connect from calling api to fetch data from db and use these code level types and return data and code level type should be optimised with design pattern. How will you map the code level type with database type and how would you structure the code like directory level.
