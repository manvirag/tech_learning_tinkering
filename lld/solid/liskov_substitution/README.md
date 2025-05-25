Don't use the direct instance of a type. Always accept an interface so that the It is easy to accept
another implementation of the interface without any code change


Or

Objects of a superclass should be replaceable with objects of a subclass without affecting the correctness of the program.

Car car = new Ford()  ( means ford would be implementing all function of car, to solve this only implement when sure all method will be used. ) 


