package main

import (
	"github.com/CaliDog/certstream-go"
	logging "github.com/op/go-logging"
	"fmt"
	"strings"
)

var log = logging.MustGetLogger("example")

func main() {
	// The false flag specifies that we want heartbeat messages.
	stream, errStream := certstream.CertStreamEventStream(false)
	for {
		select {
			case jq := <-stream:
				domains, err := jq.ArrayOfStrings("data", "leaf_cert", "all_domains")
				if err != nil{
					log.Fatal("Error decoding jq string")
				}
				for _, domain := range domains {
					fmt.Println(strings.Replace(domain, "*.", "", 1))
				}	
			case err := <-errStream:
				log.Error(err)
		}
	}
}
