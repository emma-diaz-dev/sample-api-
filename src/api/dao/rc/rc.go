package rc

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Rest interface {
	Get(url string) ([]byte, error)
}

func CreateBaseRestClient(t time.Duration) Rest {
	return &rest{
		client: &http.Client{Timeout: t},
	}
}

type rest struct {
	client *http.Client
}

func (r *rest) Get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))
	if res.StatusCode >= 300 {
		fmt.Printf("[Status: %s] [Code: %d]", res.Status, res.StatusCode)
		return nil, errors.New(string(bytes))
	}
	return bytes, nil
}
