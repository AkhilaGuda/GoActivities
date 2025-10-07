# File reader activity
- This program prompts the user to input a file path, then reads and prints the contents of that file to the console.
- Buffered I/O - it's a standard Go package that helps to read or write data
- bufio.Scanner is a type provided by bufio that lets to read input line by line, word by word.

# Features
- Reads file line by line
- Handles errors like file not found or permission denied
- Works with any text file

# How to Run locally
1. Clone the repository:
    - git clone https://github.com/AkhilaGuda/GoActivities.git
    - cd GoActivities/module2/module2_files
    - go run file-reader.go

# Error handling
- Uses os.Open with error checks
- Handles scanner.Err() for read errors
- Prints appropriate messages for failures