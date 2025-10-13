package main

import "gorm.io/gorm"

// BlogPost struct with fields Title, Content & gorm.Model
type GenericTable struct {
	gorm.Model
	Title   string
	Content string
}

// BlogService provides CRUD operations for blog posts
type BlogService struct {
	db *gorm.DB
}

// BlogRepository interface defines methods for CRUD operations
type BlogRepository interface {
	CreatePost(post *GenericTable) error
	GetAllPosts() ([]GenericTable, error)
	GetPostByID(id uint) (*GenericTable, error)
	UpdatePost(id uint, newData GenericTable) error
	DeletePost(id uint) error
}
