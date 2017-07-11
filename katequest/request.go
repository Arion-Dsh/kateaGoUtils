package katequest

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Katequest ...
type Katequest struct {
	client http.Client
}

// Get ...
func (k *Katequest) Get(url string) {
	resp, _ := k.client.Get(url)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Print(string(body))

}

func get(url string) {

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Print(string(body))
}
