package main

import (
	"fmt"
	"sync"
)

func main() {
	ids, err := fetchUsers()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	var wg sync.WaitGroup
	ch := make(chan User, len(ids))
	sem := make(chan bool, 5)
	for _, id := range ids {
		wg.Add(1)
		go func(userId int) {
			defer wg.Done()
			fetchUser(userId, ch, sem)
		}(id)
	}
	wg.Wait()
	close(ch)
	fmt.Println("Fetched users: ")
	for user := range ch {
		fmt.Printf("Id: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}
}
