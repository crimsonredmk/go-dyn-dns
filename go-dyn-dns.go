package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"os"
)

func getPublicIP() ([]byte, error) {
	res, err := http.Get("http://ip.nfriedly.com/text")
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(res.Body)
}

func main() {
	publicIP, err := getPublicIP()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", publicIP)
}
