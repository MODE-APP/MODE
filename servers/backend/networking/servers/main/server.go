package main

import (
	servers "MODE/servers/backend/networking/servers/serverTypes"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	log.Println("Starting main")
	wd, err := os.Getwd()
	eCheck(err)
	fmt.Println(wd)
	pub := filepath.Join(wd, "../../../", "certs/ModeCertificate.crt")
	priv := filepath.Join(wd, "../../../", "certs/ModeKey.key")
	log.Println(pub)
	serv := servers.NewTLSserver("localhost", "3218", pub, priv)
	log.Println("serve")
	err = serv.Serve()
	log.Fatalf("%v", err)
}

func eCheck(err error) {
	if err != nil {
		log.Fatalf("%v", err)
	}
}
