package main

import (
	clienttests "MODE/servers/backend/networking/diagnostics/clientTests"
	"bufio"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"strings"
	"time"
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
	fmt.Println("Enter port")
	port, _ := reader.ReadString('\n')
	port = strings.TrimSuffix(port, "\n")
	time.Sleep(2 * time.Second)
	make, finish, _ := clienttests.ManyClientsManyRequests(1000, 100, port)
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
