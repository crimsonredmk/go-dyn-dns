package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"os"
	"flag"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/route53"
	"encoding/json"
)

type config struct {
	AWS_ACCESS_KEY string
	AWS_SECRET_ACCESS_KEY string
	HOSTED_ZONE_ID string
	SUBDOMAIN string
}

func readConfigFile(configFileLocation string) (*config, error) {
	fileContent, err := ioutil.ReadFile(configFileLocation)
	logErrorThenExit(err)

	res := &config{}

	if err := json.Unmarshal(fileContent, res); err != nil {
		return nil, err
	}

	return res, nil
}

func getPublicIP() ([]byte, error) {
	res, err := http.Get("http://checkip.amazonaws.com")
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(res.Body)
}

func connectToRoute53(AWS_ACCESS_KEY string, AWS_SECRET_ACCESS_KEY string) (*route53.Route53, error) {
	auth := aws.Auth{
		AccessKey: AWS_ACCESS_KEY,
		SecretKey: AWS_SECRET_ACCESS_KEY,
	}
	return route53.NewRoute53(auth)
}

func updateRoute53Record(publicIP string, configMap *config, route53Client *route53.Route53) (*route53.ChangeResourceRecordSetsResponse, error) {
	rrv := []route53.ResourceRecordValue{route53.ResourceRecordValue{Value: publicIP}}
	basicRRS := route53.BasicResourceRecordSet{Action: "UPSERT", Name: configMap.SUBDOMAIN, Type: "A", TTL: 60, Values: rrv}
	RRSSlice := []route53.ResourceRecordSet{route53.ResourceRecordSet(basicRRS)}
	createRequest := route53.ChangeResourceRecordSetsRequest{Changes: RRSSlice, Xmlns: "https://route53.amazonaws.com/doc/2013-04-01/"}

	return route53Client.ChangeResourceRecordSet(&createRequest, configMap.HOSTED_ZONE_ID)
}

func logErrorThenExit(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func main() {
	configFileLocation := flag.String("c", "/etc/go-route53-dyn-dns-conf.json", "configuration file location")
	flag.Parse()

	publicIP, err := getPublicIP()
	logErrorThenExit(err)

	configMap, err := readConfigFile(*configFileLocation)
	logErrorThenExit(err)

	route53Client, err := connectToRoute53(configMap.AWS_ACCESS_KEY, configMap.AWS_SECRET_ACCESS_KEY)
	logErrorThenExit(err)

	_, err = updateRoute53Record(string(publicIP), configMap, route53Client)
	logErrorThenExit(err)
	fmt.Printf("Successfully mapped %s to IP %s.\n", configMap.SUBDOMAIN, publicIP)
}
