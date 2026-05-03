package webFetcher

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func WebFetcher(urlChan chan string, url string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(rand.Intn(5) * int(time.Second)))
	urlChan <- fmt.Sprintf("Fetched : %s", url)
}

func FetchResult(urlChan chan string) {
	for data := range urlChan {
		log.Println(data)
	}
}
