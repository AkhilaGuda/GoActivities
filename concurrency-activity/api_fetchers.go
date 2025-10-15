package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func fetchUsers() ([]int, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var list []User
	err = json.Unmarshal(body, &list)
	if err != nil {
		return nil, err
	}
	var ids []int
	for _, user := range list {
		ids = append(ids, user.ID)
	}
	return ids, nil
}
func fetchUser(id int, ch chan<- User, sem chan bool) {
	sem <- true              //acquire
	defer func() { <-sem }() //release
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching user :", id, "error: ", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("Non 200 response ", id)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading user ", id, "error: ", err)
		return
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("Error parsing user ", id, err)
		return
	}

	ch <- user

}
