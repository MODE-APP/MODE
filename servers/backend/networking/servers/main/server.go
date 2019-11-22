package main

import (
	"bufio"
	_ "net/http/pprof"
	"strings"

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
	pub := filepath.Join(wd, "../../../", "certs/off-host-crt.pem")
	priv := filepath.Join(wd, "../../../", "certs/off-host.pem")
	log.Println(pub)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter address")
	address, _ := reader.ReadString('\n')
	address = strings.TrimSuffix(address, "\n")
	fmt.Println("Enter port")
	port, _ := reader.ReadString('\n')
	port = strings.TrimSuffix(port, "\n")
	fmt.Println(port)
	serv := servers.NewTLSserver(address, port, pub, priv)
	log.Println("serve")
	err = serv.Serve()

	log.Fatalf("%v", err)
}

func eCheck(err error) {
	if err != nil {
		log.Fatalf("%v", err)
	}
}
