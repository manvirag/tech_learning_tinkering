Before going in this, let's first understand what does data modelling and ER diagram means in leyman term. ( Object oriented/class diagram is also a type of data modelling , we can ask with interview -> then can create both -> different between both would be one invole design pattern and different relations like has / is etc. )


- So Data Modelling types are:
    - ER
    - Relation (which we will convert it about in this)
    - object oriented ( check the oops template )


- In leyman term , data modelling mean -> we want to make the database for our system. 
- Assuming single database and mysql for simplicity.
- This include 
    - Defining different entities or table.
    - Define some entity which come after combination of two or more entity, relation like student course
    - What will be fields in it or column.
    - What will be its primary key .
    - What will be its foreign key .
    - How it is connected to other table like what's the relation, 1:many, many:1, many:many etc.
    - Constraints of field and type -> like string non null etc.
    - Index as well.
- This sums up the data modelling part of requirements.
- Now what the heck is ER diagram -> we can also ignore this, data modellign is about discussing real implementation , ER just to show the visually , like showing entity and its field and its relation with other entity by drawing that's it. This is indepent of database type.



### Template for Relational Modelling and ER:


#### Relation Modelling: (Every data base have its differnt modelling , mostly interviewer ask for object oriented data modelling or object oriented design or Relational Data modelling)
- [Prerequisite]Understand the RDMS constaints -> foreign key , primary key, contraints, relationship(1:m ,m:1 ,m:n).
- Figure out the entities ( or table which you want to create)
- Figure out relation between entities and write it. ( in case of 1:m, m:1, m:m) -> we usually keep id of second into first , but what if it can be zero, then that column space will be unused -> in this case we can have different table mapping both.
- Define the table and their fields. 
- Write the type of fields
- Write primary and foreign key.
- Write the constaints of field.
- Normalize if require.
- Create Relevant indexes
```
-- User Table
CREATE TABLE User (
    user_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL
);

-- Book Table
CREATE TABLE Book (
    book_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(100) NOT NULL,
    category VARCHAR(50)
);

-- Borrowing Table (Join Table to handle many-to-many relationship)
CREATE TABLE Borrowing (
    borrowing_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    book_id BIGINT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE,
    FOREIGN KEY (user_id) REFERENCES User(user_id),
    FOREIGN KEY (book_id) REFERENCES Book(book_id)
);

CREATE INDEX idx_user_id ON Borrowing(user_id);
CREATE INDEX idx_book_id ON Borrowing(book_id);

```


#### ER Modelling:

It is nothing but like this, can ignore this

```
+-------------------+      +-------------------+      +------------------------+
|      User         |      |      Book         |      |      Borrowing          |
+-------------------+      +-------------------+      +------------------------+
| user_id (PK)      |      | book_id (PK)      |<-----| borrowing_id (PK)       |
| name              |      | title             |      | start_date              |
| email             |      | author            |      | end_date                |
+-------------------+      | category          |----->| user_id (FK)            |
                            +-------------------+      | book_id (FK)            |
                                                        +------------------------+

```

https://akhileshmj.medium.com/lld-3-schema-design-df7925322982
https://akhileshmj.medium.com/lld-3-schema-design-case-study-20aa6fbd124d



### Remember mostly interview ask of oops design , so ask them first and relational in system design.

#### Next step is to go each lld problem and draw their relational diagram