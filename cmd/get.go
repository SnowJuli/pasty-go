package gopasty

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetRequest(baseUrl) {

	resp, err := http.Get(baseUrl)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	res := string(body)
	log.Printf(res)
}
