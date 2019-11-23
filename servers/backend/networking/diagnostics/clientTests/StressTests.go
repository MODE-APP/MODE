package clienttests

import (
	clients "MODE/servers/backend/networking/clients/clientTypes"
	generalservices "MODE/servers/backend/networking/proto/generated/generalservices"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"google.golang.org/grpc/connectivity"
)

//SingleClientManyRequests runs a request on a single client for reqs number of times and returns how long it takes
func SendManyRequestsConcurrently(client clients.TLSClient, reqs int) (time.Duration, error) {
	c := make(chan int)
	if !client.IsConnected() {
		return 0, errors.New("client not connected")
	}
	now := time.Now()
	for i := 0; i < reqs; i++ {
		go func(c chan int) {
			client.RequestRefreshToken("chase", "mypassword")
			c <- 0
		}(c)
	}
	for i := 0; i < reqs; i++ {
		<-c
	}
	return time.Since(now), nil
}

//SingleClientManyRequests runs a request on a single client for reqs number of times and returns how long it takes
func SendManyRequestsNonConcurrently(client clients.TLSClient, reqs int) (time.Duration, error) {
	now := time.Now()
	creds := &generalservices.Credentials{Password: "mypassword", Username: "admin"}
	for i := 0; i < reqs; i++ {
		client.ApplyCredentialsToMetadata(creds)
		client.TestCall()
	}
	return time.Since(now), nil
}

//ManyClientManyRequests runs reqs number of requests on clientsNum of clients concurrently and returns time taken for making clients and making the calls
func ManyClientsManyRequests(clientsNum, reqs int, port, addr string) (time.Duration, time.Duration, error) {
	mClients, makeTime, err := CreateManyTLSClients(clientsNum, port, addr)
	if err != nil {
		return 0, 0, err
	}
	c := make(chan int)
	now := time.Now()
	printMemUsage()
	for i := 0; i < clientsNum; i++ {
		go func(c chan int, i int, client clients.TLSClient) {
			_, err := SendManyRequestsConcurrently(client, reqs)
			if err != nil {
				panic(err)
			}
			client.Close()
			i = <-c
		}(c, i, mClients[i])
	}
	fmt.Printf("time to start requests: %v\n", time.Since(now))
	for i := 0; i < clientsNum; i++ {
		c <- i
	}
	for _, client := range mClients {

		if client.GetState() != connectivity.Shutdown {
			client.Close()
			fmt.Println("shutting down client")
		}
	}
	return makeTime, time.Since(now), nil

}

func CreateManyTLSClients(numOf int, port, addr string) ([]clients.TLSClient, time.Duration, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, 0, err
	}
	clientDurs := map[int]time.Duration{}
	cert := filepath.Join(wd, "../../../", "certs/off-host-crt.pem")
	mClients := make([]clients.TLSClient, numOf)
	clientC := make(chan clients.TLSClient)
	intC := make(chan int)
	durC := make(chan time.Duration)
	now := time.Now()
	for i := 0; i < numOf; i++ {
		go func(clientC chan clients.TLSClient, intC chan int, durC chan time.Duration, i int) {
			now := time.Now()
			client, err := clients.NewTLSClient(addr, port, cert)
			if err != nil {
				panic(err)
			}
			err = client.Connect()
			if err != nil {
				fmt.Println("error was not nil")
			}
			for !client.IsConnected() {
				time.Sleep(time.Millisecond * 200)
			}
			if !client.IsConnected() {
				fmt.Println("returning unconnected")
			}
			intC <- i
			clientC <- client
			durC <- time.Since(now)
		}(clientC, intC, durC, i)
	}
	fmt.Printf("time to start con: %v\n", time.Since(now))
	for i := 0; i < numOf; i++ {
		x := <-intC
		mClients[x] = <-clientC
		clientDurs[i] = <-durC
	}
	return mClients, time.Since(now), nil
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

func CreateManyTLSClientsNonC(numOf int, port, addr string) ([]clients.TLSClient, time.Duration, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, 0, err
	}
	cert := filepath.Join(wd, "../../../", "certs/off-host-crt.pem")
	mClients := make([]clients.TLSClient, numOf)
	now := time.Now()
	for i := 0; i < numOf; i++ {
		client, err := clients.NewTLSClient(addr, port, cert)
		if err != nil {
			panic(err)
		}
		err = client.Connect()
		if err != nil {
			panic(err)
		}
		for !client.IsConnected() {
			time.Sleep(1 * time.Millisecond)
		}
		fmt.Println(i)
		mClients[i] = client
	}

	return mClients, time.Since(now), nil
}

func CreateManyEssentialClients(numOf int) ([]clients.EssentialClient, time.Duration, error) {
	clientDurs := map[int]time.Duration{}
	mClients := make([]clients.EssentialClient, numOf)
	clientC := make(chan clients.EssentialClient)
	intC := make(chan int)
	durC := make(chan time.Duration)
	now := time.Now()
	go func() {
		for i := 0; i < numOf; i++ {
			go func(clientC chan clients.EssentialClient, intC chan int, durC chan time.Duration, i int) {
				now := time.Now()
				client := clients.NewEssentialClient("localhost", "3218")
				err := client.Connect()
				if err != nil {
					panic(err)
				}
				for !client.IsConnected() {
					time.Sleep(5 * time.Millisecond)
				}
				intC <- i
				clientC <- client
				durC <- time.Since(now)
			}(clientC, intC, durC, i)
		}
	}()
	for i := 0; i < numOf; i++ {
		x := <-intC
		mClients[x] = <-clientC
		clientDurs[i] = <-durC
	}
	for i := 0; i < numOf; i++ {
		fmt.Printf("%v client done took %v\n", i, clientDurs[i])
	}
	return mClients, time.Since(now), nil
}
