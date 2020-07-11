package main

import (
	"fmt"
	"net"
	"net/rpc"
	"time"
)

type Server struct{}

func (this *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}

func server() {
	rpc.Register(new(Server))
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		connection, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(connection)
	}
}

func client() {
	connection, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var result int64
	err = connection.Call("Server.Negate", int64(999), &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server.Negate(999) =", result)
	}
}

func main() {
	go server()
	time.Sleep(200 * time.Millisecond) // Just to wait server start
	go client()
	var input string
	fmt.Scanln(&input)
}
