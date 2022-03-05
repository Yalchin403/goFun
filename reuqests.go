package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Please enter the URL: ")
	var url string
	fmt.Scanln(&url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	var goMap map[string]interface{}

	strBody := string(body)
	err = json.Unmarshal([]byte(strBody), &goMap)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(goMap)

}
