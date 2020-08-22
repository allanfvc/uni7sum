package controller

import (
	"io/ioutil"
	"net/http"

	"github.com/sony/gobreaker"
)

var cb *gobreaker.CircuitBreaker

func initBreaker() {
	var st gobreaker.Settings
	st.Name = "HTTP GET"
	cb = gobreaker.NewCircuitBreaker(st)
}

// Get wraps http.Get in CircuitBreaker.
func getWithBreaker(url string) ([]byte, error) {
	body, cbError := cb.Execute(func() (interface{}, error) {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return body, err
	})
	cb.State()
	if cbError != nil {
		return nil, cbError
	}

	return body.([]byte), nil
}
