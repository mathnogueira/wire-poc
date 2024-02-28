package paidhttp

import (
	"net/http"

	"github.com/gorilla/mux"
	ht "github.com/mathnogueira/ioc/http"
	"github.com/mathnogueira/ioc/repository"
)

type controller struct {
	innerController ht.Controller
	userRepository  repository.UserRepository
}

func NewController(userRepository repository.UserRepository) ht.Controller {
	innerController := ht.NewController(userRepository)
	return &controller{innerController: innerController, userRepository: userRepository}
}

func (c *controller) Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", c.Greet).Methods("GET")
	router.HandleFunc("/hey", c.Hey).Methods("GET")

	return router
}

func (c *controller) Greet(resp http.ResponseWriter, req *http.Request) {
	c.innerController.Greet(resp, req)
}

func (c *controller) Hey(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Ho!"))
}
