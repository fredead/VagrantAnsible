package main
import (
        "fmt"
        "net/http"
        "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
        h, _ := os.Hostname()
	if len(r.URL.Query().Get("name")) <= 0  {
       	   fmt.Fprintf(w, "Hi there, I'm served from %s!", h)
        } else {
       	   fmt.Fprintf(w, "Hi there %s, I'm served from %s!", r.URL.Query().Get("name"), h)
        }
}

func main() {
        http.HandleFunc("/", handler)
        http.ListenAndServe(":8484", nil)
}
