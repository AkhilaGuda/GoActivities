package db

import "gorm.io/gorm"

// BlogPost struct with fields Title, Content & gorm.Model
type GenericTable struct {
	gorm.Model
	Title   string
	Content string
}

var DynamicTableName string

// TableName method dynamically changes tableName
func (g GenericTable) TableName() string {
	return DynamicTableName
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
