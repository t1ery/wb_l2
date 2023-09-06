package main

import (
    "fmt"
    "sync"
    "time"
)

// Функция sig создает канал, который будет закрыт после указанной задержки.
func sig(after time.Duration) <-chan interface{} {
    c := make(chan interface{})
    go func() {
        defer close(c)
        time.Sleep(after)
    }()
    return c
}

func main() {
    start := time.Now()
    doneChans := []<-chan interface{}{
        sig(2 * time.Hour),
        sig(5 * time.Minute),
        sig(1 * time.Second),
        sig(1 * time.Hour),
        sig(1 * time.Minute),
    }

    // Слияние каналов и ожидание завершения.
    mergedChan := or(doneChans...)
    <-mergedChan
    fmt.Printf("Завершено после %v\n", time.Since(start))
}

// Функция or объединяет несколько каналов в один и возвращает его.
func or(channels ...<-chan interface{}) <-chan interface{} {
    switch len(channels) {
    case 0:
        // Если нет входных каналов, создаем и возвращаем закрытый канал.
        c := make(chan interface{})
        close(c)
        return c
    case 1:
        // Если есть только один входной канал, просто возвращаем его.
        return channels[0]
    }

    // Создаем канал, который будет объединять результаты.
    orCh := make(chan interface{})

    // Создаем sync.Once для закрытия orCh.
    var closeOnce sync.Once

    // Создаем горутину для отслеживания завершения каналов.
    go func() {
        var wg sync.WaitGroup
        for _, ch := range channels {
            wg.Add(1)
            go func(ch <-chan interface{}) {
                defer wg.Done()
                for v := range ch {
                    orCh <- v
                }
                // При завершении работы канала, закрываем orCh только один раз.
                closeOnce.Do(func() {
                    close(orCh)
                })
            }(ch)
        }
        wg.Wait()
    }()

    return orCh
}
