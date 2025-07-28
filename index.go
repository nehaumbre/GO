package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

// //Package declaration:
// //this is entry point for the executable programs in GO
// //Packages are named after the directory they are in (preffered practise but not a strict requirement)

// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// //import statements
// //import statements are used to import packages and functions from other packages that your program depends on
// //the packages are imported using the keyword "import" followed by the package name
// //the fmt package is for formatting and printing

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World! %s", time.Now())
// }

// func main() {
// 	http.HandleFunc("/", greet)
// 	http.ListenAndServe(":8080", nil)
// }

// //main function : where execution of program begins(entry point)
