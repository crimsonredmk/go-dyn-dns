package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"os"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/route53"
	"encoding/json"
)

type config struct {
	AWS_ACCESS_KEY string
	AWS_SECRET_ACCESS_KEY string
	ZONE_NAME string
	SUBDOMAIN string
}

func readConfigFile() (*config, error) {
	fileContent, err := ioutil.ReadFile("/etc/go-dyn-dns-conf.json")
	logErrorThenExit(err)

	res := &config{}

	if err := json.Unmarshal(fileContent, res); err != nil {
		return nil, err
	}

	return res, nil
}

func getPublicIP() ([]byte, error) {
	res, err := http.Get("http://ip.nfriedly.com/text")
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(res.Body)
}

func connectToRoute53() (*route53.Route53, error) {
	auth := aws.Auth{
		AccessKey: "ASDFASDFASDFASDK",
		SecretKey: "DSFSDFDWESDADSFASDFADFDSFASDF",
	}
	return route53.NewRoute53(auth)
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

	route53Client, err := connectToRoute53()
	logErrorThenExit(err)
	fmt.Printf("%s\n", route53Client.Endpoint)

	configMap, err := readConfigFile()
	logErrorThenExit(err)
	fmt.Printf("%s\n", configMap.ZONE_NAME)
}
