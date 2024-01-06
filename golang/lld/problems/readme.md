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