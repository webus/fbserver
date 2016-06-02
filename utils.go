package main

import (
	"log"
	"io/ioutil"
	"net/http"
)


func debugRead(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.Status)
	log.Println(string(body))
}
