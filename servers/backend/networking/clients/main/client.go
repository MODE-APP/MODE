package main

import (
	clients "MODE/servers/backend/networking/clients/clientTypes"
	"log"
	"os"
)

func main() {
	log.Println("Starting main")
	client := clients.NewEssentialClient("localhost", "3218")
	client.Connect()
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

}
