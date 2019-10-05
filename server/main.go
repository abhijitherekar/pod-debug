package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

// func echoServer(c net.Conn) {
// 	for {
// 		buf := make([]byte, 512)
// 		nr, err := c.Read(buf)
// 		if err != nil {
// 			return
// 		}

// 		data := buf[0:nr]
// 		println("Server got:", string(data))
// 		_, err = c.Write(data)
// 		if err != nil {
// 			log.Fatal("Writing client error: ", err)
// 		}
// 	}
// }

func printhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
func main() {
	log.Println("Starting echo server on : ", runtime.GOOS)
	http.HandleFunc("/", printhello)
	http.ListenAndServe(":80", nil)
}

// 	//addr, err := net.ResolveTCPAddr("TCP", "")
// 	os.Remove("@test")
// 	ln, err := net.Listen("unix", "@test")
// 	if err != nil {
// 		log.Fatal("Listen error: ", err)
// 	}
// 	fmt.Println(ln.Addr())
// 	sigc := make(chan os.Signal, 1)
// 	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
// 	go func(ln net.Listener, c chan os.Signal) {
// 		sig := <-c
// 		log.Printf("Caught signal %s: shutting down.", sig)
// 		ln.Close()
// 		os.Exit(0)
// 	}(ln, sigc)

// 	for {
// 		fd, err := ln.Accept()
// 		if err != nil {
// 			log.Fatal("Accept error: ", err)
// 		}

// 		go echoServer(fd)
// 	}
// }
