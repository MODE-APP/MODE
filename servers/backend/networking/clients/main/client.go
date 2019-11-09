package main

import (
	clients "MODE/servers/backend/networking/clients/clientTypes"
	"log"
	"os"
)

func main() {
	log.Println("Starting main")
	cert, err := os.Open("/home/arline/go/src/MODE/servers/backend/certs/ModeCertificate.crt")
	eCheck(err)
	client, err := clients.NewTLSClient("localhost", "3218", cert)
	eCheck(err)
	err = client.Connect()
	eCheck(err)
	token, err := client.RequestToken("chasearline", "mypassword")
	eCheck(err)
	client.Token = token
	client.ApplyTokenToMetadata()
	client.TestCall()
	if err != nil {
		panic(err)
	}
}

func eCheck(err error) {
	if err != nil {
		panic(err)
	}
}
