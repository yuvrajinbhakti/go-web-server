This Go server demonstrates how to implement a simple CRUD (Create, Read, Update, Delete) REST API. It includes all the essential operations while maintaining clarity and readability.

Features:
Routing: Handled using http.HandleFunc with custom endpoint paths.
Concurrency Safety: A sync.Mutex is used to prevent race conditions when modifying the in-memory users map.
Endpoints:
GET /users: Fetch all users.
POST /users: Add a new user.
GET /users/{id}: Retrieve a specific user by ID.
PUT /users/{id}: Update an existing user by ID.
DELETE /users/{id}: Delete a user by ID.
How It Works:
main Function:

Sets up the server with two handlers: /users for operations on all users and /users/{id} for individual users.
Starts the server on port 8080.
Request Handlers:

handleUsers: Manages requests for all users (GET to fetch, POST to create).
handleUserByID: Handles user-specific operations (GET, PUT, DELETE) based on user ID extracted from the URL.
CRUD Operations:

Create: Decodes JSON input to a User struct and adds it to the users map.
Read: Retrieves data from the users map. Returns all users (GET /users) or a specific user by ID (GET /users/{id}).
Update: Decodes new data and updates the existing user in the map (PUT /users/{id}).
Delete: Removes a user from the map (DELETE /users/{id}).
Error Handling:

Returns appropriate HTTP status codes for errors:
400 Bad Request for invalid JSON.
404 Not Found if the user doesn't exist.
409 Conflict for duplicate user IDs.
405 Method Not Allowed for unsupported HTTP methods.
Example Usage:
Start the server:

bash
Copy code
go run main.go
Server runs on http://localhost:8080.

Test endpoints:

Create a user:
bash
Copy code
curl -X POST -H "Content-Type: application/json" -d '{"id": "1", "name": "John Doe", "email": "john@example.com"}' http://localhost:8080/users
Get all users:
bash
Copy code
curl http://localhost:8080/users
Update a user:
bash
Copy code
curl -X PUT -H "Content-Type: application/json" -d '{"name": "John Smith", "email": "johnsmith@example.com"}' http://localhost:8080/users/1
Delete a user:
bash
Copy code
curl -X DELETE http://localhost:8080/users/1
Why It’s Beginner-Friendly:
Uses Go’s standard library (net/http) for simplicity.
No external dependencies or advanced concepts.
Clear separation of logic for each CRUD operation.



Explanation:
Setup:

gorilla/mux is used for routing.
In-memory storage (map) is used to store users, synchronized with a sync.Mutex for thread safety.
Routes:

POST /users: Create a new user. Validates the payload and ensures the user ID is unique.
GET /users: Fetch all users.
GET /users/{id}: Fetch a specific user by ID.
PUT /users/{id}: Update an existing user by ID.
DELETE /users/{id}: Delete a user by ID.
Concurrency:

A sync.Mutex ensures safe concurrent access to the users map.
HTTP Response Codes:

201 Created: When a user is successfully created.
404 Not Found: If a user with a specified ID is not found.
409 Conflict: If a user with the same ID already exists.
204 No Content: When a user is successfully deleted.
Running the Code:
Install dependencies: go get -u github.com/gorilla/mux.
Run the server with go run main.go.
Use tools like Postman or curl to test the API.
