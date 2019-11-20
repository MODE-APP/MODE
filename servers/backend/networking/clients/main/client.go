package main

import (
	clients "MODE/servers/backend/networking/clients/clientTypes"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6061", nil))
	}()
	/**
	printMemUsage()
	clientStressTest()
	printMemUsage()
	time.Sleep(10 * time.Second)
	*/
	log.Println("Starting single client stress test")
	printMemUsage()
	rpcStressTest()
	printMemUsage()
	runtime.GC()

	printMemUsage()
	log.Println("10 no connect")
	noConnectStressTest(10)
	printMemUsage()
	runtime.GC()

	printMemUsage()
	log.Println("100 no connect")
	noConnectStressTest(100)
	printMemUsage()
	runtime.GC()

	printMemUsage()
	log.Println("1000 no connect")
	noConnectStressTest(1000)
	printMemUsage()
	runtime.GC()

	printMemUsage()
	log.Println("10000 no connect")
	noConnectStressTest(10000)
	printMemUsage()
	runtime.GC()

}

func eCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func rpcStressTest() {
	const numOfReqs int = 100000
	wd, err := os.Getwd()
	eCheck(err)
	cert := filepath.Join(wd, "../../../", "certs/ModeCertificate.crt")
	client, err := clients.NewTLSClient("localhost", "3218", cert)
	eCheck(err)
	err = client.Connect()
	eCheck(err)
	for !client.IsConnected() {
		time.Sleep(15 * time.Millisecond)
	}
	defer client.Close()
	now := time.Now()
	c := make(chan int)
	for i := 0; i < numOfReqs; i++ {
		go func(c chan int) {
			//client.RequestRefreshToken("chasearline", "mypassword")
			client.RequestRefreshToken("chasearline", "mypassword")
			c <- 0
		}(c)
	}
	for i := 0; i < numOfReqs; i++ {
		<-c
	}
	fmt.Printf("%v RPC in %v\n", numOfReqs, time.Since(now))
}

func clientStressTest() {
	numOfClients := 10000
	c := make(chan int)
	wd, err := os.Getwd()
	eCheck(err)
	cert := filepath.Join(wd, "../../../", "certs/ModeCertificate.crt")
	now := time.Now()
	for i := 0; i < numOfClients; i++ {
		go func(c chan int, i int) {
			client, err := clients.NewTLSClient("localhost", "3218", cert)
			eCheck(err)
			err = client.Connect()
			eCheck(err)
			for !client.IsConnected() {
				time.Sleep(15 * time.Millisecond)
			}
			defer client.Close()
			token, err := client.RequestRefreshToken("chasearline", "mypassword")
			eCheck(err)
			client.RefreshToken = token
			client.ApplyTokenToMetadata(client.RefreshToken)
			token, err = client.RequestAccessToken()
			eCheck(err)
			client.AccessToken = token
			client.ApplyTokenToMetadata(client.AccessToken)
			client.TestCall()
			token, err = client.RequestAccessToken()
			eCheck(err)
			i = <-c
		}(c, i)
	}
	for i := 0; i < numOfClients; i++ {
		c <- i
	}
	fmt.Printf("Finished in %v\n", time.Since(now))
}

func noConnectStressTest(numOfClients int) {
	c := make(chan int)
	wd, err := os.Getwd()
	eCheck(err)
	cert := filepath.Join(wd, "../../../", "certs/ModeCertificate.crt")
	now := time.Now()
	mClients := make([]clients.TLSClient, numOfClients)
	clientC := make(chan clients.TLSClient)
	intC := make(chan int)
	for i := 0; i < numOfClients; i++ {
		go func(clientC chan clients.TLSClient, intC chan int, i int) {
			client, err := clients.NewTLSClient("localhost", "3218", cert)
			eCheck(err)
			err = client.Connect()
			eCheck(err)
			for !client.IsConnected() {
				time.Sleep(5 * time.Millisecond)
			}
			intC <- i
			clientC <- client
		}(clientC, intC, i)
	}
	for i := 0; i < numOfClients; i++ {
		mClients[<-intC] = <-clientC
	}

	log.Printf("Time to connect %v clients: %v", numOfClients, time.Since(now))
	printMemUsage()
	runtime.GC()
	start := time.Now()
	for i := 0; i < numOfClients; i++ {
		go func(c chan int, i int, client clients.TLSClient) {
			defer client.Close()
			token, err := client.RequestRefreshToken("chasearline", "mypassword")
			eCheck(err)
			client.RefreshToken = token
			client.ApplyTokenToMetadata(client.RefreshToken)
			token, err = client.RequestAccessToken()
			eCheck(err)
			client.AccessToken = token
			client.ApplyTokenToMetadata(client.AccessToken)
			client.TestCall()
			token, err = client.RequestAccessToken()
			eCheck(err)
			i = <-c
		}(c, i, mClients[i])
	}
	for i := 0; i < numOfClients; i++ {
		c <- i
	}
	fmt.Printf("Finished in %v\n", time.Since(start))
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
