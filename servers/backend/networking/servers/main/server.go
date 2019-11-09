package main

import (
	servers "MODE/servers/backend/networking/servers/serverTypes"
	"log"
	"os"
)

func main() {
	log.Println("Starting main")
	pub, err := os.Open("/home/arline/go/src/MODE/servers/backend/certs/ModeCertificate.crt")
	if err != nil {
		panic(err)
	}
	priv, err := os.Open("/home/arline/go/src/MODE/servers/backend/certs/ModeKey.key")
	if err != nil {
		panic(err)
	}
	serv := servers.NewTLSserver("localhost", "3218", pub, priv)
	serv.Serve()
}
