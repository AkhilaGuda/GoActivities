package main

import (
	"flag"
	"fmt"
	"gormActivity/repository"
)

func main() {
	tableName := flag.String("table", "", "Name of the table to create")
	flag.Parse()
	if *tableName == "" {
		fmt.Println("Please provide a table name using -table")
		return
	}
	repository.DynamicTableName = *tableName
	fmt.Println("Table name: ", *tableName)
	service, err := repository.NewBlogService()
	if err != nil {
		panic(err)
	}
	post := repository.GenericTable{Title: "Hello golang", Content: "This blog contains information about go lang"}
	if err := service.CreatePost(&post); err != nil {
		fmt.Println("Error creating post:", err)
	}
	addedPost, _ := service.GetPostByID(1)
	fmt.Println("Added post: ", addedPost)
	posts, _ := service.GetAllPosts()
	fmt.Println("All posts:\n", posts)
	service.UpdatePost(post.ID, repository.GenericTable{Title: "Golanguage", Content: "Updated content"})
	updatedPost, _ := service.GetPostByID(post.ID)
	fmt.Println("Updated post: ", updatedPost)
	service.DeletePost(post.ID)
	fmt.Println("Post deleted")
	posts, _ = service.GetAllPosts()
	fmt.Println(posts)
}
