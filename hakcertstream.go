package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/CaliDog/certstream-go"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	// The false flag specifies that we want heartbeat messages.
	stream, errStream := certstream.CertStreamEventStream(false)
	for {
		select {
		case jq := <-stream:
			domains, err := jq.ArrayOfStrings("data", "leaf_cert", "all_domains")
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error decoding json", err)
			} else {
				for _, domain := range domains {
					fmt.Fprintln(w, strings.Replace(domain, "*.", "", 1))
				}
			}
		case err := <-errStream:
			fmt.Fprintln(os.Stderr, "Stream error", err)
		}
	}
}
