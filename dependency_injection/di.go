package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Greet(os.Stdout, "Elodie")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))

}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)

}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
func FmtGreet(name string) {
	fmt.Printf("Hello, %s", name)
}
