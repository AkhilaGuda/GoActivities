# Blog Application with Go & GORM
- Created a new Go module using go mod init gormActivity
- Installed required dependencies (gorm.io/gorm, gorm.io/driver/mysql, github.com/joho/godotenv).
- Created a .env file for database credentials: host, port, user, password, database name.
- Used godotenv to load environment variables in Go.
- Created a database service struct to manage connection.
- Connected to MySQL using GORM’s MySQL driver.
- Create: Method to insert new blog posts into the database.
- Read: Methods to get all posts or a post by ID.
- Update: Method to modify an existing post by ID.
- Delete: Method to remove a post by ID.
- Used db.Create() for inserting records.
- Used db.Find() and db.First() for querying.
- Used db.Save()  for updating records.
- Used db.Delete() for deleting records.
# Run locally
- Clone the repository: git clone https://github.com/AkhilaGuda/GoActivities.git
- cd gorm-activity
- set up environment variables
- go run . -table="tableName"
