package main

import (
	"net"
	"testing"
	"time"
)

func TestTelnetClient(t *testing.T) {
	// Запускаем тестовый сервер
	go startTestServer()

	// Ждем, чтобы сервер успел запуститься
	time.Sleep(1 * time.Second)
}

func startTestServer() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go handleTestConnection(conn)
	}
}

func handleTestConnection(conn net.Conn) {
	defer conn.Close()

	// Просто отправляем "Hello, World!" в ответ
	conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Length: 13\r\n\r\nHello, World!"))
}
