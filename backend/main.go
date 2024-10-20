package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Allow CORS
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    // Handle OPTIONS request
    if r.Method == http.MethodOptions {
        return
    }

    fmt.Fprintf(w, "Hello from Go!")
}

func main() {
	fs := http.FileServer(http.Dir("../frontend/dist"))
    http.Handle("/dist/", http.StripPrefix("/dist/", fs))
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "../frontend/index.html")
    })
    
    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
