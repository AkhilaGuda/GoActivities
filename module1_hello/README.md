# Activity 1
* Installed go from go.dev site
* For removing already existing go installation and extract the newly downloaded go
* command :  rm -rf /usr/local/go && tar -C /usr/local -xzf go1.25.1.linux-amd64.tar.gz

* Added /usr/bin/local/go/bin to path environment variable
* export PATH =$PATH :/usr/local/go/bin 
 added this in ~/.zshrc file
* To make it available system level
* source ~/.zshrc
* go version
* Results in downloaded go version: go version go1.25.1 linux/amd64



# Activity 2
* go mod init <module_path>: This command used to initialize a new go module within project, It created go.mod file which intiially contains go version, it acts as manifest to track dependencies.
* module_path : It becomes official name and import path for module as other modules to import packges 
* It is primarly used for setup tasks, such as global variables, opening database connections or registering dependencies.

* Created hello.go file 
* Main fucntion is the entry point for go program
* imported "fmt" as it is a standard package for formatted input and outputs
* fmt offers functions like fmt.Print, fmt.Println, fmt.Scan,fmt.Scanln etc
* Running the program : In terminal command : go run <filename> 
* Go program is executed : Output - Helo, GO!!