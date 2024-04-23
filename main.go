package main

import (
	"log"
	"math/rand"
	"net"
	"time"
)

const (
	nw   = "tcp4"
	addr = ":8080"
)

var (
	arr = []string{"Concurrency is not parallelism", "Channels orchestrate; mutexes serialize.", "The bigger the interface, the weaker the abstraction.", "Make the zero value useful.", "interface{} says nothing.", "Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.", "A little copying is better than a little dependency.", "Cgo is not Go."}
)

func main() {
	listener, err := net.Listen(nw, addr)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		go handleconn(conn)
	}
}

func handleconn(conn net.Conn) {
	length := len(arr)
	for {
		time.Sleep(time.Second * 3)
		index := rand.Intn(length - 1)
		verb := arr[index]
		conn.Write([]byte(verb + "\n"))
	}

}
