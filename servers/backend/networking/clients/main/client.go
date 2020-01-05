package main

import (
	clients "MODE/servers/backend/networking/clients/clientTypes"
	"bufio"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6061", nil))
	}()
	fmt.Println("starting main:")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter address")
	address, _ := reader.ReadString('\n')
	address = strings.TrimSuffix(address, "\n")
	fmt.Println("Enter port")
	port, _ := reader.ReadString('\n')
	port = strings.TrimSuffix(port, "\n")
	fmt.Printf("starting client")
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cert := filepath.Join(wd, "../../../", "certs/ModeCertificate.pem")
	fmt.Println(cert)
	tls, err := clients.NewTLSClient(address, port, cert)
	if err != nil {
		panic(err)
	}
	err = tls.Connect()
	if err != nil {
		panic(err)
	}
	_, err = tls.RequestRefreshToken("chasearline", "mypassword")
	if err != nil {
		panic(err)
	}
	ess := clients.NewEssentialClient(address, port)
	err = ess.Connect()
	if err != nil {
		panic(err)
	}
	err = ess.TestCall()
	if err != nil {
		panic(err)
	}

}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MB", m.Alloc/1000000)
	fmt.Printf("\tTotalAlloc = %v MB", m.TotalAlloc/1000000)
	fmt.Printf("\tSys = %v MB", m.Sys/1000000)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
