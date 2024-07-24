package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

// Функция для запуска телнет-клиента
func runTelnetClient(host, port, timeout string) error {
	// Парсинг таймаута
	dur, err := time.ParseDuration(timeout)
	if err != nil {
		return fmt.Errorf("ошибка при парсинге таймаута: %w", err)
	}

	address := fmt.Sprintf("%s:%s", host, port)

	// Установка соединения
	conn, err := net.DialTimeout("tcp", address, dur)
	if err != nil {
		return fmt.Errorf("ошибка подключения: %w", err)
	}
	defer conn.Close()

	// Запуск горутин для чтения данных из сокета и записи в STDOUT
	done := make(chan error, 1)
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		done <- err
	}()

	// Запуск горутины для чтения данных из STDIN и записи в сокет
	_, err = io.Copy(conn, os.Stdin)
	if err != nil {
		return fmt.Errorf("ошибка при записи в сокет: %w", err)
	}

	// Ожидание завершения чтения из сокета
	if err := <-done; err != nil {
		return fmt.Errorf("ошибка при чтении из сокета: %w", err)
	}

	return nil
}

func main() {
	// Параметры командной строки
	timeout := flag.String("timeout", "10s", "таймаут подключения (по умолчанию 10s)")
	host := flag.String("host", "", "хост для подключения")
	port := flag.String("port", "", "порт для подключения")
	flag.Parse()

	if *host == "" || *port == "" {
		fmt.Println("Использование: go-telnet --timeout=10s <host> <port>")
		os.Exit(1)
	}

	// Запуск телнет-клиента
	err := runTelnetClient(*host, *port, *timeout)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
