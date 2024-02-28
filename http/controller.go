package http

import (
	"fmt"
	"net/http"

	"github.com/mathnogueira/ioc/repository"
)

type Controller interface {
	Greet(rw http.ResponseWriter, req *http.Request)
}

type controller struct {
	userRepository repository.UserRepository
}

func NewController(userRepository repository.UserRepository) Controller {
	return &controller{userRepository: userRepository}
}

func (c *controller) Greet(resp http.ResponseWriter, req *http.Request) {
	user, err := c.userRepository.Current()
	if err != nil {
		resp.Write([]byte(err.Error()))
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp.Write([]byte(fmt.Sprintf("hello, %s", user.Name)))
}
