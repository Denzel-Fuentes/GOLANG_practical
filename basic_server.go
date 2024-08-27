package main
import (
    "fmt"
    "log"
    "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello!")
}

func main() {
    http.HandleFunc("/", helloWorld)
    log.Println("Server is running in port 8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("Error to load server: ", err)
    }
}
