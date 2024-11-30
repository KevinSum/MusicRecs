package main

import (
	"MusicRecs/CLI"
	"MusicRecs/server"
)

func main() {
	// chanel to check if server is set up
	serverReady := make(chan bool)

	go server.SetupServer(serverReady)
	<-serverReady
	CLI.StartCLI()
}
