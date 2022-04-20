package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := initRoutes()

	log.Fatal(http.ListenAndServe(":8000", router))
}

type Book struct {
	Id     string `json:"id"`
	Isbn   string `json:"isbn"`
	Author Author `json:"author"`
}

type Author struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	books := []Book{
		{
			Id:   "1",
			Isbn: "ASD1231S",
			Author: Author{
				Id:   "A1",
				Name: "Coel",
			},
		}, {
			Id:   "2",
			Isbn: "VBDd1231F",
			Author: Author{
				Id:   "A2",
				Name: "Jackson",
			},
		},
	}
	resp, _ := json.Marshal(books)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book Book

	errResponse := []byte("{\"Error\":\"Unable to parse request body\"}")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
		w.Write(errResponse)
		return
	}
	err = json.Unmarshal(body, &book)
	if err != nil {
		fmt.Println(err.Error())
		w.Write(errResponse)
		return
	}
	fmt.Printf("book %+v", book)
	resp, _ := json.Marshal(book)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

func updateBook(w http.ResponseWriter, r *http.Request)  {
	var rbook Book
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &rbook)
	param := mux.Vars(r)
	id := param["id"]
	books := []Book{
		{
			Id:   "1",
			Isbn: "ASD1231S",
			Author: Author{
				Id:   "A1",
				Name: "Coel",
			},
		},
	}

	for idx, book := range books {
		if book.Id == id {
			books[idx] = rbook
			break
		}
	}
	resp,_ := json.Marshal(books)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func initRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/books", getBooks).Methods(http.MethodGet)
	r.HandleFunc("/book", addBook).Methods(http.MethodPost)
	r.HandleFunc("/book/{id}", updateBook).Methods(http.MethodPut)

	return r
}
