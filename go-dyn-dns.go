package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"os"
	// "github.com/awslabs/aws-sdk-go/aws"
	// "github.com/awslabs/aws-sdk-go/service/route53"
)

func getPublicIP() ([]byte, error) {
	res, err := http.Get("http://ip.nfriedly.com/text")
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(res.Body)
}

func logErrorThenExit(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func main() {
	publicIP, err := getPublicIP()
	logErrorThenExit(err)

	fmt.Printf("%s\n", publicIP)
}
