package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// fetchUsers retrieves the list of all users from the API and extracts their IDs
func fetchUsers() ([]int, error) {
	// Make a GET request to the users endpoint
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Read the full response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var list []User
	// Parse (unmarshal) JSON data into a slice of User structs
	err = json.Unmarshal(body, &list)
	if err != nil {
		return nil, err
	}
	var ids []int
	// Extract user IDs from the list
	for _, user := range list {
		ids = append(ids, user.ID)
	}
	// Return list of user IDs
	return ids, nil
}

// fetchUser fetches the details of a single user by ID concurrently
// It sends the result into channel and uses a semaphore to limit concurrency
func fetchUser(id int, ch chan<- User, sem chan bool) {
	sem <- true              //acquire semaphore slot - blocks if limit is reached
	defer func() { <-sem }() //release semaphore slot once done
	// specific API for specific user
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", id)
	// send http get request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching user :", id, "error: ", err)
		return
	}
	defer resp.Body.Close()
	// check if response status is 200(ok)
	if resp.StatusCode != 200 {
		fmt.Println("Non 200 response ", id)
		return
	}
	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading user ", id, "error: ", err)
		return
	}
	// parse json response into a User struct
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("Error parsing user ", id, err)
		return
	}
	// Send the fetched user back through the channel
	ch <- user

}
