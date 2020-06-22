package handlers

import (
	"fmt"
	"log"
	"net/http"

	operations "github.com/avdmsajaykumar/exercise3/dboperations"
)

type DBHandler struct {
	logger *log.Logger
}

func NewDBHandler(l *log.Logger) *DBHandler {
	return &DBHandler{l}
}

func (h *DBHandler) Create(rw http.ResponseWriter, r *http.Request) {
	request := new(operations.Data)
	err := request.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to parse Input", http.StatusBadRequest)
		return
	}
	response := request.Create()
	fmt.Fprintf(rw, "%v", response)
}

func (h *DBHandler) Get(rw http.ResponseWriter, r *http.Request) {
	request := new(operations.Data)
	err := request.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to parse Input", http.StatusBadRequest)
		return
	}
	response := request.Get()
	response.ToJSON(rw)

}

func (h *DBHandler) Update(rw http.ResponseWriter, r *http.Request) {
	request := new(operations.Data)
	err := request.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to parse Input", http.StatusBadRequest)
		return
	}

	OldDocument, NewDocument := request.Update()
	fmt.Fprint(rw, "\n Old Document \n")
	OldDocument.ToJSON(rw)
	fmt.Fprint(rw, "\n Updated Document \n")
	NewDocument.ToJSON(rw)

}

func (h *DBHandler) Delete(rw http.ResponseWriter, r *http.Request) {
	request := new(operations.Data)
	err := request.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to parse Input", http.StatusBadRequest)
		return
	}

	result, status := request.Delete()
	if status == true {
		fmt.Fprintf(rw, "%d Document(s) has been deleted sucessfully\n ", result)
	} else {
		fmt.Fprint(rw,
			"Failed to delete the document\n Document may been deleted already or it doesn't exist")
	}

}
