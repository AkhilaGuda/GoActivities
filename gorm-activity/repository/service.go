package repository

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// // TableName method dynamically changes tableName
// func (g GenericTable) TableName() string {
// 	return DynamicTableName
// }

// NewBlogService initializes a new Blogservice with database connection
// It loads environment varirables for database connection
// Runs AutoMigrate to create the table if it doesn't exists
func NewBlogService() (BlogRepository, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found")
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&GenericTable{}); err != nil {
		return nil, err
	}
	fmt.Println("Data base connected")
	return &BlogService{db: db}, nil
}

// CreatePost inserts a new blog post into the database
// Returns an error if the insertion fails
//
//post:pointer to GenericTable struct to insert
func (s *BlogService) CreatePost(post *GenericTable) error {
	return s.db.Create(post).Error
}

// GetAllPosts retrieves all blog posts from the database
// returns a slice of GenericTable and error if query fails
func (s *BlogService) GetAllPosts() ([]GenericTable, error) {
	var posts []GenericTable
	err := s.db.Find(&posts).Error
	return posts, err
}

// GetPostByID retrieves a single post by its ID
// returns error if post is not found
func (s *BlogService) GetPostByID(id uint) (*GenericTable, error) {
	var post GenericTable
	err := s.db.First(&post, id).Error
	return &post, err
}

// UpdatePost updates exisiting blog post identified with ID with new data
// newData: BlogPost consisting with updated fields
// return an error if the update fails
func (s *BlogService) UpdatePost(id uint, newData GenericTable) error {
	var post GenericTable
	if err := s.db.First(&post, id).Error; err != nil {
		return err
	}
	post.Title = newData.Title
	post.Content = newData.Content
	return s.db.Save(&post).Error
}

// DeletePost method deletes the post with id
// returns error if id is not found
func (s *BlogService) DeletePost(id uint) error {
	return s.db.Delete(&GenericTable{}, id).Error
}
