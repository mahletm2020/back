// package main

// import (
//   "database/sql"       // For working with SQL databases
//   "encoding/json"      // For encoding and decoding JSON
//   "fmt"                // For formatting strings
//   "net/http"           // For creating HTTP servers
//   "time"               // For working with time

//   "github.com/dgrijalva/jwt-go" // For creating JWT tokens
//   _ "github.com/lib/pq"         // PostgreSQL database driver
// )

// // Define the JWT key used for signing tokens
// var jwtKey = []byte("your_secret_key")

// // Credentials struct to read the username and password from the request body
// type Credentials struct {
//   Username string `json:"username"`
//   Password string `json:"password"`
// }

// // Claims struct to include the username and role in the JWT token
// type Claims struct {
//   Username string `json:"username"`
//   Role     string `json:"role"`
//   jwt.StandardClaims
// }

// // loginHandler function to handle login requests
// func loginHandler(w http.ResponseWriter, r *http.Request) {
//   var creds Credentials
//   // Decode the JSON request body into the Credentials struct
//   err := json.NewDecoder(r.Body).Decode(&creds)
//   if err != nil {
//     w.WriteHeader(http.StatusBadRequest) // Return 400 if there's a bad request
//     return
//   }

//   // Connect to the PostgreSQL database
//   db, err := sql.Open("postgres", "user=postgres password=yourpassword dbname=techpulsedb sslmode=disable")
//   if err != nil {
//     w.WriteHeader(http.StatusInternalServerError) // Return 500 if there's an internal server error
//     return
//   }
//   defer db.Close() // Ensure the database connection is closed

//   // Query the database to check if the user exists and get their role
//   var username, role string
//   err = db.QueryRow(`
//     SELECT u.username, r.role_name 
//     FROM users u 
//     JOIN roles r ON u.role_id = r.id 
//     WHERE u.username=$1 AND u.password=$2
//   `, creds.Username, creds.Password).Scan(&username, &role)
//   if err != nil {
//     w.WriteHeader(http.StatusUnauthorized) // Return 401 if the user is not authorized
//     return
//   }

//   // Create the JWT token
//   expirationTime := time.Now().Add(24 * time.Hour) // Set the token to expire in 24 hours
//   claims := &Claims{
//     Username: username,
//     Role:     role,
//     StandardClaims: jwt.StandardClaims{
//       ExpiresAt: expirationTime.Unix(), // Set the expiration time
//     },
//   }

//   token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Create a new token with the specified signing method and claims
//   tokenString, err := token.SignedString(jwtKey)             // Sign the token with the JWT key
//   if err != nil {
//     w.WriteHeader(http.StatusInternalServerError) // Return 500 if there's an internal server error
//     return
//   }

//   // Return the token in the response
//   json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
// }

// // main function to start the HTTP server
// func main() {
//   http.HandleFunc("/login", loginHandler) // Register the loginHandler for the /login route
//   http.ListenAndServe(":8080", nil)       // Start the server on port 8080
// }
