package main

import (
	"fmt"
	"net/http"
	"time"
	"flag"
)
func PostLocalTime(w http.ResponseWriter, r *http.Request){
	//Prints the local time in the "DD MMM YY HH(24):MM ZON" format
	fmt.Fprintf(w,"<b>The current local time is: %s",time.Now().Local().Format(time.RFC822))
}
func main() {
	// redirect all urls to the PostLocalTime function
	http.HandleFunc("/",PostLocalTime)
	//Gets the flag value for port or sets to 8080 by default
	port := flag.Int("port", 8080, "port to serve on")
	// executes the command-line parsing
	flag.Parse()
	fmt.Printf("Running on port %d\n", *port)
	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	//Listen for connections at port *port on the local machine
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}
