package main
 
import (
  "log"
  "fmt"
  "net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}

func main() {

	http.HandleFunc("/", Handler)

	err := http.ListenAndServe("0.0.0.0:3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
