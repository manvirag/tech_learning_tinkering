package main

import (
	"codesignal/models"
	"codesignal/service"
	"codesignal/util"
	"fmt"
	"time"
)

/*
Solutions:

choose and book fitness classes

user management, class scheduling, and booking, including waitlisting and cancellation features.


flow -> admin will create some type of class ( gym ) -> there will be some start time of this + atendee count.
	-> user register and take packages
	-> then user will book for class let say above some
			-> if avaible and have quota -> book
			-> not quota -> reset package
			-> if not avaiable -> waitlist. -> if booked cacnel -> get place.

*/

/*
Objective: Design and implement an application to allow users to choose and book fitness classes. The application should cater to user management, class scheduling, and booking, including waitlisting and cancellation features.

Requirements:

User Management:

Registration and Login: Users should be able to register and log in to the system.

User Tiers: Users can be categorized into three tiers: Platinum, Gold, and Silver. Each tier has a different booking limit:

Platinum: 10 classes
Gold: 5 classes
Silver: 3 classes

==> is it life time or per day ?? or per month -> in per Packages
==> what about once this package is done ??

Packages:
The user tiers align with the packages, determining the number of classes a user can book.

Classes:

Types: Classes can include yoga, gym, and dance etc etc.

Capacity: Each class has a maximum number of attendees.

Scheduling: Classes are scheduled at specific times, and multiple classes can run in a single day.

Booking:
Users can book a class if it has not reached capacity.

An admin can create, schedule, and cancel classes.

Waitlisting: If a class is full, users can join a waitlist. When a booked user cancels, the first user on the waitlist is allocated the slot.

Cancellation: Users can cancel their booking up to 30 minutes before the class starts. Admins can also cancel classes.

Methods: This is just to help you understand. Actual implementation may differ.

User Management: Register/Login, Select Package

Classes (Admin): Create a class, Schedule a class, Cancel a class
Booking (User): Book a class, Cancel a class

Additional Considerations:

- Users should be able to book classes at different times.
- Concurrency management for multiple users trying to book the same class simultaneously.
- Handling package strategies effectively.
- Ensuring that user cancellations restore their booking quota.

Evaluation Criteria:

Design:

- Use of object-oriented design principles (classes, interfaces, inheritance, polymorphism).
- Efficiency and scalability of data structures.
- Flexibility and modularity of system components.
- Consideration of security and data privacy principles.

Implementation:

-Functionality and accuracy of system features.
-Performance and resource optimization of algorithms.
-Code readability, maintainability, and adherence to coding standards.
-Error handling and exception management for smooth operation.

Expectations:

-Provide clean, professional-level code.
-Model core entities and relationships effectively.
-Demonstrate the solution with a method-based approach.
-Use memory-based object storage; a backend database is not required.
[execution time limit] 0.5 seconds (cpp)

[memory limit] 2g

In interview
was asked to implement a method such that user shouldn't be able to book a class it has already booked a class in same timeslot. (was given 5 mins for that)
*/

import "codesignal/dao"

func main() {
	ac := util.NewAtomicCounter()
	ud := dao.NewUserDao()
	us := service.NewUserService(ud, ac)
	err := us.RegisterUser("normaluser")
	fmt.Println(err)
	err = us.AddAdminUser("anubhav")
	fmt.Println(err)
	err = us.LogInUser("normaluser")
	fmt.Println(err)
	err = us.BuyPackage("normaluser", models.Premium)
	fmt.Println(err)
	fmt.Println(".........")
	cd := dao.NewClassDao()
	cs := service.NewClassService(cd, ud, ac)
	cs.CreateAdminClassCategory("anubhav", models.Yoga)
	cs.CreateClasses("anubhav", models.Yoga, time.Now(), time.Now().Add(2*time.Hour), 2)
	fmt.Println(".........")
	bd := dao.NewBookingDao()
	bs := service.NewBookingService(bd, ud, cd, ac)
	err = bs.BookingClassSchedule("normaluser", 3) // booking user, class scheduler id.
	err = bs.BookingClassSchedule("normaluser", 3) // booking user, class scheduler id.
	bs.CancelBooking(4, "normaluser")
	err = bs.BookingClassSchedule("normaluser", 3) // booking user, class scheduler id.
	fmt.Println(err)
}
