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