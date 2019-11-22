package main

import (
	clients "MODE/servers/backend/networking/clients/clientTypes"
	clienttests "MODE/servers/backend/networking/diagnostics/clientTests"
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
	/**
	printMemUsage()
	make, finish, err := clienttests.ManyClientsManyRequests(1000, 100)
	printMemUsage()
	fmt.Printf("MakeClients: %v\tFinishReqs: %v\terr: %v\n", make, finish, err)
	clients, time, err := clienttests.CreateManyTLSClientsNonC(1000)
	fmt.Printf("1000 connects: %v\n", time)
	time, err = clienttests.SendManyRequestsConcurrently(clients[0], 100000)
	printMemUsage()
	fmt.Printf("1000 req: %v\n", time)
	*/
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
	cert := filepath.Join(wd, "../../../", "certs/off-host-crt.pem")
	client, err := clients.NewTLSClient(address, port, cert)
	if err != nil {
		panic(err)
	}
	err = client.Connect()
	if err != nil {
		panic(err)
	}
	_, _, err = client.FetchCertificate()
	if err != nil {
		panic(err)
	}
	make, finish, _ := clienttests.ManyClientsManyRequests(100000, 10, port, address)
	fmt.Printf("%v\t%v", make, finish)
	printMemUsage()

}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MB", m.Alloc/1000000)
	fmt.Printf("\tTotalAlloc = %v MB", m.TotalAlloc/1000000)
	fmt.Printf("\tSys = %v MB", m.Sys/1000000)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
