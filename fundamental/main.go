package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	userID := 10
	resChan := make(chan string, 3)
	wg := &sync.WaitGroup{}

	go fetchUserData(userID, resChan, wg)
	go fetchUserRecommendation(userID, resChan, wg)
	go fetchUserLikes(userID, resChan, wg)
	wg.Add(3)
	wg.Wait()
	close(resChan)

	for resp := range resChan {
		fmt.Println(resp)
	}

	fmt.Println(time.Since(now))
}

func fetchUserData(userID int, resChan chan string, wg *sync.WaitGroup) {
	time.Sleep(80 * time.Millisecond)
	resChan <- "user data"
	wg.Done()
}

func fetchUserRecommendation(userID int, resChan chan string, wg *sync.WaitGroup) {
	time.Sleep(120 * time.Millisecond)
	resChan <- "user recommendation"
	wg.Done()
}

func fetchUserLikes(userID int, resChan chan string, wg *sync.WaitGroup) {
	time.Sleep(50 * time.Millisecond)
	resChan <- "user likes"
	wg.Done()
}
