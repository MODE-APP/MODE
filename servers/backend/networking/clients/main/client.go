package main

import (
	clienttests "MODE/servers/backend/networking/diagnostics/clientTests"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6061", nil))
	}()

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
	time.Sleep(2 * time.Second)
	clienttests.CreateManyTLSClients(2000)
	makeTime, finishTime, err := clienttests.ManyClientsManyRequests(1000, 300)
	if err != nil {
		panic(err)
	}
	printMemUsage()
	fmt.Printf("MakeTime: %v\tFinishTime: %v\n", makeTime, finishTime)

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
