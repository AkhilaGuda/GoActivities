package main

import (
	"flag"
	"fmt"
)

var DynamicTableName string

func main() {
	tableName := flag.String("table", "", "Name of the table to create")
	flag.Parse()
	if *tableName == "" {
		fmt.Println("Please provide a table name using -table")
		return
	}
	DynamicTableName = *tableName
	fmt.Println("Table name: ", *tableName)
	var blogRepo BlogRepository
	service, err := NewBlogService()
	if err != nil {
		panic(err)
	}
	blogRepo = service
	post := GenericTable{Title: "Hello golang", Content: "This blog contains information about go lang"}
	if err := blogRepo.CreatePost(&post); err != nil {
		fmt.Println("Error creating post:", err)
	}
	addedPost, _ := blogRepo.GetPostByID(1)
	fmt.Println("Added post: ", addedPost)
	posts, _ := blogRepo.GetAllPosts()
	fmt.Println("All posts:\n", posts)
	blogRepo.UpdatePost(post.ID, GenericTable{Title: "Golanguage", Content: "Updated content"})
	updatedPost, _ := blogRepo.GetPostByID(post.ID)
	fmt.Println("Updated post: ", updatedPost)
	blogRepo.DeletePost(post.ID)
	fmt.Println("Post deleted")
	posts, _ = blogRepo.GetAllPosts()
	fmt.Println(posts)
}
