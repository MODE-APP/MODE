package main

import (
	servers "MODE/servers/backend/networking/servers/serverTypes"
	"log"
)

func main() {
	log.Println("Starting main")
	serv := servers.NewEssentialServer("localhost", "3218")
	serv.Serve()
}
