package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "Server: ", log.Lshortfile)
	logDecorator := LogDecoratorCreator(logger)

	http.Handle("/example", logDecorator(HeaderDecorator(http.HandlerFunc(myHandlerFunc))))

	http.Handle("/example2", logDecorator(HeaderDecorator(http.HandlerFunc(myHandlerFunc))))

	log.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func myHandlerFunc(w http.ResponseWriter, r *http.Request) {
	//Write Status Code
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	fmt.Fprintln(w, "Example-1; Request was successful")
}

func myHandlerFunc2(w http.ResponseWriter, r *http.Request) {
	//Write Status Code
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	fmt.Fprintln(w, "Example-2; Request was successful")
}

// HeaderDecorator takes http.Handler and returns http.Handler
func HeaderDecorator(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Passed-Through-Decorator", "YES")
		h.ServeHTTP(w, r)
	}) //returns an HTTP Handler, http.HandlerFunc( returns an http handler
}

// HTTPDecorator type that is a function that takes http.Handler and returns http.Handler
type HTTPDecorator func(http.Handler) http.Handler

// LogDecoratorCreator takes a *log.logger and returns HTTPDecorator
func LogDecoratorCreator(logger *log.Logger) HTTPDecorator {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Printf(r.RequestURI+": We have a connection from %s", r.RemoteAddr)
			defer logger.Println("Connection Closed")
			h.ServeHTTP(w, r)
		}) //returns an HTTP Handler, http.HandlerFunc( returns an http handler
	} //Returns a function that takes a handler and returns a handler
}
