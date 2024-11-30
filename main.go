package main

import (
	"MusicRecs/server"
)

func main() {
	// chanel to check if server is set up
	serverReady := make(chan bool)

	go server.SetupServer(serverReady)
	<-serverReady
}
