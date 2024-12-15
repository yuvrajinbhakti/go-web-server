// package main


// import (
//     "fmt"
//     "log"
//     "net/http"
// )

// func formHandler(w http.ResponseWriter, r *http.Request) {
//     if err := r.ParseForm(); err != nil {
//         fmt.Fprintf(w, "ParseForm() err: %v", err)
//         return
//     }
//     fmt.Fprintf(w, "POST request successful")
//     name := r.FormValue("name")
//     address := r.FormValue("address")
//     fmt.Fprintf(w, "Name = %s\n", name)
//     fmt.Fprintf(w, "Address = %s\n", address)
// }

// func helloHandler(w http.ResponseWriter, r *http.Request) {
//     if r.URL.Path != "/hello" {
//         http.Error(w, "404 not found.", http.StatusNotFound)
//         return
//     }

//     if r.Method != "GET" {
//         http.Error(w, "Method is not supported.", http.StatusNotFound)
//         return
//     }


//     fmt.Fprintf(w, "Hello!")
// }


// func main() {
//     fileServer := http.FileServer(http.Dir("./static"))
//     http.Handle("/", fileServer)
//     http.HandleFunc("/form", formHandler)
//     http.HandleFunc("/hello", helloHandler)


//     fmt.Printf("Starting server at port 8080\n")
//     if err := http.ListenAndServe(":8080", nil); err != nil {
//         log.Fatal(err)
//     }
// }






// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"sync"

// 	"github.com/gorilla/mux"
// )

// // User represents a user structure
// type User struct {
// 	ID    string `json:"id"`
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// }

// var (
// 	users  = make(map[string]User) // in-memory user storage
// 	mu     sync.Mutex             // to handle concurrent access
// )

// func main() {
// 	r := mux.NewRouter()

// 	// Routes for CRUD operations
// 	r.HandleFunc("/users", createUser).Methods("POST")
// 	r.HandleFunc("/users", getAllUsers).Methods("GET")
// 	r.HandleFunc("/users/{id}", getUser).Methods("GET")
// 	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
// 	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

// 	// Start the server
// 	fmt.Println("Server is running on http://localhost:8080")
// 	http.ListenAndServe(":8080", r)
// }

// // createUser creates a new user
// func createUser(w http.ResponseWriter, r *http.Request) {
// 	var user User
// 	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
// 		http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 		return
// 	}

// 	mu.Lock()
// 	defer mu.Unlock()

// 	if _, exists := users[user.ID]; exists {
// 		http.Error(w, "User already exists", http.StatusConflict)
// 		return
// 	}

// 	users[user.ID] = user
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(user)
// }

// // getAllUsers returns all users
// func getAllUsers(w http.ResponseWriter, r *http.Request) {
// 	mu.Lock()
// 	defer mu.Unlock()

// 	var userList []User
// 	for _, user := range users {
// 		userList = append(userList, user)
// 	}

// 	json.NewEncoder(w).Encode(userList)
// }

// // getUser retrieves a user by ID
// func getUser(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	mu.Lock()
// 	defer mu.Unlock()

// 	user, exists := users[id]
// 	if !exists {
// 		http.Error(w, "User not found", http.StatusNotFound)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(user)
// }

// // updateUser updates an existing user
// func updateUser(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	mu.Lock()
// 	defer mu.Unlock()

// 	if _, exists := users[id]; !exists {
// 		http.Error(w, "User not found", http.StatusNotFound)
// 		return
// 	}

// 	var updatedUser User
// 	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
// 		http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 		return
// 	}

// 	updatedUser.ID = id // Ensure the ID remains the same
// 	users[id] = updatedUser
// 	json.NewEncoder(w).Encode(updatedUser)
// }

// // deleteUser removes a user by ID
// func deleteUser(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	mu.Lock()
// 	defer mu.Unlock()

// 	if _, exists := users[id]; !exists {
// 		http.Error(w, "User not found", http.StatusNotFound)
// 		return
// 	}

// delete(users, id)
// 	w.WriteHeader(http.StatusNoContent)
// }






package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// User represents a simple user
// Each user has an ID, a Name, and an Email
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Store for users and a lock for safe concurrent access
var (
	users = make(map[string]User) // Stores users in memory
	lock  sync.Mutex             // Ensures safe access to the user store
)

func main() {
	// Define the API endpoints
	http.HandleFunc("/users", handleUsers)
	http.HandleFunc("/users/", handleUserByID)

	// Start the server
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// handleUsers manages requests for all users
func handleUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getAllUsers(w)
	} else if r.Method == "POST" {
		createUser(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleUserByID manages requests for a single user by ID
func handleUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/users/"):] // Extract ID from URL

	if r.Method == "GET" {
		getUser(w, id)
	} else if r.Method == "PUT" {
		updateUser(w, r, id)
	} else if r.Method == "DELETE" {
		deleteUser(w, id)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// createUser adds a new user
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	lock.Lock()
	defer lock.Unlock()

	if _, exists := users[user.ID]; exists {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	users[user.ID] = user
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// getAllUsers retrieves all users
func getAllUsers(w http.ResponseWriter) {
	lock.Lock()
	defer lock.Unlock()

	var userList []User
	for _, user := range users {
		userList = append(userList, user)
	}

	json.NewEncoder(w).Encode(userList)
}

// getUser retrieves a user by ID
func getUser(w http.ResponseWriter, id string) {
	lock.Lock()
	defer lock.Unlock()

	user, exists := users[id]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// updateUser modifies an existing user
func updateUser(w http.ResponseWriter, r *http.Request, id string) {
	var updatedUser User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	lock.Lock()
	defer lock.Unlock()

	if _, exists := users[id]; !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	updatedUser.ID = id // Keep the same ID
	users[id] = updatedUser
	json.NewEncoder(w).Encode(updatedUser)
}

// deleteUser removes a user by ID
func deleteUser(w http.ResponseWriter, id string) {
	lock.Lock()
	defer lock.Unlock()

	if _, exists := users[id]; !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	delete(users, id)
	w.WriteHeader(http.StatusNoContent)
}
