package main

import (
	_ "net/http/pprof"

	servers "MODE/servers/backend/networking/servers/serverTypes"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	log.Println("Starting main")
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	wd, err := os.Getwd()
	eCheck(err)
	fmt.Println(wd)
	pub := filepath.Join(wd, "../../../", "certs/ModeCertificate.pem")
	priv := filepath.Join(wd, "../../../", "certs/ModeKey.pem")
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
