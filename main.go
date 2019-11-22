package main

import (
	"bytes"
	"log"

	"github.com/imroc/req"
)

func main() {
	// use Req object to initiate requests.
	r := req.New()
	err := r.SetProxyUrl("http://209.90.63.108:80")
	if err != nil {
		log.Printf("%v", err)
	}
	resp, err := r.Get("https://api.ipify.org/")
	if err != nil {
		log.Printf("%v", err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Response().Body)
	newStr := buf.String()
	log.Printf("%s", newStr)
}
