package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/avdmsajaykumar/exercise3/handlers"
	"github.com/gorilla/mux"
)

func main() {

	logger := log.New(os.Stdout, "Exercise 3 :", log.LstdFlags)
	dbHandler := handlers.NewDBHandler(logger)
	Router := mux.NewRouter()

	Router.HandleFunc("/get", dbHandler.Get)
	Router.HandleFunc("/update", dbHandler.Update)
	Router.HandleFunc("/delete", dbHandler.Delete)
	Router.HandleFunc("/create", dbHandler.Create)

	Server := &http.Server{
		ReadTimeout:  5 * time.Second,  //Read timeout value set to 120 sec for any request
		Handler:      Router,           //Binds the mux router to http server
		WriteTimeout: 10 * time.Second, //Write timeout value set to 120 sec for any request
		IdleTimeout:  60 * time.Second, //Idle timeout value set to 120 sec for any request
		Addr:         ":80",            // Addr on which service is listened
	}

	err := Server.ListenAndServe()

	if err != nil {
		log.Printf("Error %v\n", err)
	}
}
