package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type User struct {
	Name     string `json:"username"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	ImageUrl string `json:"image"`
	ImageB64 string `json:"imageBase64,omitempty"`
}

type wrapper struct {
	Users []User `json:"users"`
}

func fetchData() []byte {
	//get request to api
	resp, err := http.Get("https://dummyjson.com/users?limit=100")
	if err != nil {
		log.Fatal("error fetching api :", err)

	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Error reading response: ", err)
	}
	return data

}
func base64Conversion(usersResponse wrapper) {
	for i, user := range usersResponse.Users {
		if user.ImageUrl == "" {
			continue
		}
		// fetch image
		imageResp, err := http.Get(user.ImageUrl)
		if err != nil {
			fmt.Println("Failed to fetch image: ", user.Name, err)
			continue
		}
		//read image bytes
		imageBytes, err := io.ReadAll(imageResp.Body)
		imageResp.Body.Close()
		if err != nil {
			fmt.Println("Failed to read image:", user.Name, err)
			continue
		}
		// convert to base64
		usersResponse.Users[i].ImageB64 = base64.StdEncoding.EncodeToString(imageBytes)
		fmt.Println("Processed image :", user.Name)
	}
}
func storeInFile(usersResponse wrapper) {
	file, err := os.Create("users_data.json")
	if err != nil {
		log.Fatal(err)
	}
	// marshal go structs to json with indentation
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ") //pretty print
	if err := encoder.Encode(usersResponse.Users); err != nil {
		log.Fatal(err)
	}
	log.Println("Saved in file users_data.json")
}

func main() {
	data := fetchData()
	var usersResponse wrapper
	if err := json.Unmarshal(data, &usersResponse); err != nil {
		log.Fatal("Error unmarshalling JSON: ", err)
	}
	base64Conversion(usersResponse)

	fmt.Println("Data: ")
	for i, user := range usersResponse.Users {
		fmt.Println(i, user.Name, user.Age, user.Email, user.ImageUrl)
	}
	storeInFile(usersResponse)
}
