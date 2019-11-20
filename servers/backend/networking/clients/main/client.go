package main

import (
	clients "MODE/servers/backend/networking/clients/clientTypes"
	"fmt"
	"time"
)

func main() {
	clientStressTest()
	rpcStressTest()
}

func eCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func rpcStressTest() {
	cert := "/home/arline/go/src/MODE/servers/backend/certs/ModeCertificate.crt"
	client, err := clients.NewTLSClient("localhost", "3218", cert)
	eCheck(err)
	err = client.Connect()
	eCheck(err)
	for !client.IsConnected() {
		time.Sleep(15 * time.Millisecond)
	}
	now := time.Now()
	for i := 0; i < 1000; i++ {
		client.RequestRefreshToken("chasearline", "mypassword")
	}
	fmt.Printf("\n 1000 RPC in %v", time.Since(now))
}

func clientStressTest() {
	numOfClients := 1000
	c := make(chan int)
	now := time.Now()
	for i := 0; i < numOfClients; i++ {
		go func(c chan int, i int) {
			cert := "/home/arline/go/src/MODE/servers/backend/certs/ModeCertificate.crt"
			client, err := clients.NewTLSClient("localhost", "3218", cert)
			eCheck(err)
			now := time.Now()
			err = client.Connect()
			eCheck(err)
			for !client.IsConnected() {
				time.Sleep(15 * time.Millisecond)
			}
			connect := time.Since(now)
			now = time.Now()
			token, err := client.RequestRefreshToken("chasearline", "mypassword")
			eCheck(err)
			refresh := time.Since(now)
			client.RefreshToken = token
			client.ApplyTokenToMetadata(client.RefreshToken)
			now = time.Now()
			token, err = client.RequestAccessToken()
			eCheck(err)
			access := time.Since(now)
			client.AccessToken = token
			client.ApplyTokenToMetadata(client.AccessToken)
			now = time.Now()
			client.TestCall()
			test := time.Since(now)
			now = time.Now()
			token, err = client.RequestAccessToken()
			eCheck(err)
			fail := time.Since(now)
			fmt.Printf("")
			fmt.Printf("Client %v: \tConnect: %v\tRefresh: %v\tAccess: %v\tTest: %v\tFailAccess: %v\n",
				i, connect, refresh, access, test, fail)
			i = <-c
		}(c, i)
		time.Sleep(5 * time.Millisecond)
	}
	for i := 0; i < numOfClients; i++ {
		c <- i
	}
	fmt.Printf("Finished in %v\n", time.Since(now))

}
