package main

import (
	_"bufio"
	"log"
	"net"
	"time"
	"context"
	"fmt"
)

func handleConn(conn net.Conn) {
	defer conn.Close()

	channel := make(chan string, 1)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go sendMessage(ctx, conn, channel)
	go readMessage(ctx, conn, channel)
}

func sendMessage(ctx context.Context,conn net.Conn, ch chan string) {
	string := "hello world"
	count := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("writeMessage goroutine stoped")
			return
		default:
			count++
			newstr := fmt.Sprintf("%s for %d times!", string, count)
			ch <- newstr
			fmt.Printf("write:%v\n", ch)
			time.Sleep(time.Second)
		}
	}
}

func readMessage(ctx context.Context,conn net.Conn, ch <-chan string) {
	for {
		line := <-ch
		fmt.Printf("read:%v\n", line)
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("readMessage goroutine stoped")
					return
				default:
					continue
				}
			}
		}()
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatalf("listen error:%v\n", err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept err:%v\n", err)
			continue
		}

		go handleConn(conn)
	}
}