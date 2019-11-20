package main

import (
	servers "MODE/servers/backend/networking/servers/serverTypes"
	"log"
)

func main() {
	log.Println("Starting main")
	pub := "/home/arline/go/src/MODE/servers/backend/certs/ModeCertificate.crt"

	priv := "/home/arline/go/src/MODE/servers/backend/certs/ModeKey.key"

	serv := servers.NewTLSserver("localhost", "3218", pub, priv)
	serv.Serve()
}
