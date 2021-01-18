package main

import (
	"ClientServerComm/server/engine/listener"
	"ClientServerComm/server/engine/server"
	"flag"
)

func main()  {
	//get port details
	port := getPort()
	listen := listener.CreateListener(port)
	server.Server(listen)


}

func getPort() int {
	port := flag.Int("port", 8084, "TCP addr Port")
	flag.Parse()
	return *port
}