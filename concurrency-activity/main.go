package main

import (
	"fmt"
	"sync"
)

func main() {
	// Fetch all user IDs from the API
	ids, err := fetchUsers()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Buffered channel to collect fetched users
	// Buffered size = number of IDs ensures no goroutine blocks while sending
	ch := make(chan User, len(ids))
	// Semaphore channel to limit concurrency (max 5 concurrent requests)
	sem := make(chan bool, 5)
	// Loop over all user IDs and fetch each user concurrently
	for _, id := range ids {
		// Increment WaitGroup counter for each goroutine
		wg.Add(1)
		go func(userId int) {
			defer wg.Done() // Decrement counter when goroutine completes
			// Fetch user and send to channel
			fetchUser(userId, ch, sem)
		}(id)
	}
	// Wait for all fetchUser goroutines to finish
	wg.Wait()
	// Close channel after all users are fetched
	close(ch)
	fmt.Println("Fetched users: ")
	// Print the fetched users from the channel
	for user := range ch {
		fmt.Printf("Id: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}
}
