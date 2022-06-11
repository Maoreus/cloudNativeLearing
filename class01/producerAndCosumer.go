package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func producer(threadID int, wg *sync.WaitGroup, ch chan string) {
	count := 0
	for {
		time.Sleep(time.Second * 1)
		count++
		data := strconv.Itoa(threadID) + "---" + strconv.Itoa(count)
		fmt.Printf("producer, %s\n", data)
		ch <- data
	}
	wg.Done()
}

func consumer(wg *sync.WaitGroup, ch chan string) {
	for data := range ch {
		time.Sleep(time.Second * 1)
		fmt.Printf("consumer, %s\n", data)
	}
	wg.Done()
}

func main() {
	steam := make(chan string, 100)
	producerWg := new(sync.WaitGroup)
	consumerWg := new(sync.WaitGroup)
	for i := 10; i <= 10; i++ {
		producerWg.Add(i)
		go producer(i, producerWg, steam)
	}

	for i := 10; i <= 10; i++ {
		consumerWg.Add(1)
		go consumer(i, consumerWg, steam)
	}

}
