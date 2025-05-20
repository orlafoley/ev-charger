**Backend Documentation**


Functionality of main.go

Connection uses gin-gonic library to communicate between frontend and backend
Frontend is running on localhost:5173/api/{insert_command}
Backend is running on localhost:8080/api/{insert_command}

models.ConnectDatabase is the initial creation of the DB

Main.go receives url and processes the information that is being given
Consist of all Functionality needed for the website (and some optionals)

Main.go calls go files inside the models directory

**Models Directory**

db.go -> initalization of all tables and connections to the database 

bookings/person.go -> Insertion / Deletion / Creation of new data 

Refactored to be 1 file each for modularization