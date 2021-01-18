package main

import (
	"flag"
	"fmt"
)

func main()  {
	//get port details
	port := getPort()
	fmt.Println(port)
}

func getPort() int {
	port := flag.Int("port", 8084, "TCP addr Port")
	flag.Parse()
	return *port
}