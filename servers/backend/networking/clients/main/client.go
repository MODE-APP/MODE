package main

import (
	clients "MODE/servers/backend/networking/clients/clientTypes"
	"log"
	"os"
)

func main() {
	log.Println("Starting main")
	cert, err := os.Open("/home/arline/go/src/MODE/servers/backend/certs/ModeCertificate.crt")
	if err != nil {
		panic(err)
	}
	client, err := clients.NewTLSClient("localhost", "3218", cert)
	if err != nil {
		panic(err)
	}
	err = client.Connect()
	if err != nil {
		panic(err)
	}
	file, filename, err := client.FetchCertificate()
	if err != nil {
		panic(err)
	}
	f, err := os.Create("/home/arline/Documents/MODE/" + filename)
	if err != nil {
		panic(err)
	}
	_, err = f.Write(file)
	if err != nil {
		panic(err)
	}

	client.TestCall()
	if err != nil {
		panic(err)
	}
}
