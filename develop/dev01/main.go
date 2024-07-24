package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func main() {
	// Получение текущего текущего времени
	localTime := time.Now()
	fmt.Printf("Текущее время: %s\n", localTime)

	// Получение точного времени с NTP-сервера
	exactTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		// Вывод ошибки в STDERR и возврат ненулевого кода выхода
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}

	// Печать точного времени
	fmt.Printf("Точное время: %s\n", exactTime)
}
