package main

import (
	"fmt"
	"time"
)

// or объединяет несколько done-каналов в один. Результирующий канал
// закроется, когда любой из входных каналов закроется.
func or(channels ...<-chan interface{}) <-chan interface{} {
	result := make(chan interface{})

	go func() {
		defer close(result)

		// Создаем слайс для хранения каналов и их состояния
		chs := make([]<-chan interface{}, len(channels))
		copy(chs, channels)

		for {
			select {
			case <-result:
				// Если результат закрылся, завершить
				return
			default:
				// Используем для прослушивания каналов
				for _, ch := range chs {
					select {
					case <-ch:
						// Если любой канал закроется, закрываем результат
						return
					default:
						// В противном случае продолжаем прослушивать
					}
				}
			}
		}
	}()

	return result
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("Done after %v\n", time.Since(start))
}
