package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// type Products []Product

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type productsHandler struct {
	sync.Mutex
	products []Product
}

func NewProductHandler() *productsHandler {
	return &productsHandler{}
}

func main1() {
	port := ":8080"

	ph := NewProductHandler()
	http.HandleFunc("/", home)
	http.Handle("/products", ph)
	// http.HandleFunc("/products/", NewProductHandler().ServeHTTP)
	http.Handle("/products/", ph)

	log.Fatal(http.ListenAndServe(port, nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome To Home Page")
}

func (p *productsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		p.handleGetProducts(w, r)

	case http.MethodPost:
		p.handleAddProduct(w, r)

	default:
		e := getHttpError(http.StatusMethodNotAllowed, "Method not allowed")
		respondWithJson(w, http.StatusMethodNotAllowed, e)
	}
}

func (p *productsHandler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	id, err := getProductIdFromURL(r)

	if err != nil || id < 0 || id > len(p.products) {
		p.getAllProducts(w, r)
	} else {
		respondWithJson(w, http.StatusOK, p.products[id])
	}
}

func (p *productsHandler) handleAddProduct(w http.ResponseWriter, r *http.Request) {
	data, err := readDataFromRequestBody(r)
	if err != nil {
		respondWithJson(w, http.StatusBadRequest,
			getHttpError(http.StatusBadRequest, "Unable to parse the request body"))
	}

	prod := Product{}

	json.Unmarshal(data, &prod)

	addProduct(p, w, prod)
}
func addProduct(p *productsHandler, w http.ResponseWriter, product Product) {
	defer p.Unlock()

	p.Lock()
	p.products = append(p.products, product)
	respondWithJson(w, http.StatusCreated, product)
}

func readDataFromRequestBody(r *http.Request) (data []byte, err error) {
	defer r.Body.Close()

	data, err = ioutil.ReadAll(r.Body)
	return
}

func getHttpError(code int, message string) (err Error) {
	err = Error{
		Code:    code,
		Message: message,
	}
	return
}

type ASlice []string

func (a ASlice) Pop() (s string, err error) {
	if len(a) <= 0 {
		return "", errors.New("empty slice")
	}
	s = a[len(a)-1]
	return
}

func getProductIdFromURL(r *http.Request) (productId int, err error) {
	var params ASlice = strings.Split(r.URL.String(), "/")
	id, err := params.Pop()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	productId, err = strconv.Atoi(id)
	return
}

func (p *productsHandler) getAllProducts(w http.ResponseWriter, r *http.Request) {
	defer p.Unlock()
	product := Product{
		Id:    1,
		Name:  "product1",
		Price: 154.9,
	}
	p.Lock()
	p.products = append(p.products, product)
	respondWithJson(w, http.StatusOK, p.products)
}

func respondWithJson(w http.ResponseWriter, statusCode int, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		response, _ = json.Marshal(getHttpError(http.StatusInternalServerError, "Error encoding response response"))
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}
