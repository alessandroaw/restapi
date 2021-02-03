package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/therealsandro/restapi/controllers"
)

func main() {
	l := log.New(os.Stdout, "book-api", log.LstdFlags)

	// Init Router
	router := mux.NewRouter()
	booksController := controllers.NewBooksController(l)

	// Route Handlers / Endpoints
	router.HandleFunc("/api/books", booksController.GetBooks).Methods("GET")
	router.HandleFunc("/api/books/{id:[0-9]+}", booksController.GetBook).Methods("GET")
	router.HandleFunc("/api/books", booksController.CreateBook).Methods("POST")
	router.HandleFunc("/api/books/{id:[0-9]+}", booksController.UpdateBook).Methods("PUT")
	// router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	s := &http.Server{
		Addr:         ":3000",
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}

	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	// http.ListenAndServe(":3000", router)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
