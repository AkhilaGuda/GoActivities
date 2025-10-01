# Command-line Flags
* I used flag.String("nums", "", "comma-separated numbers") to define a command-line flag -nums.
* flag.Parse() reads the value from the command line.
* If the user does not provide -nums, the program prints a message and exits using os.Exit(1).
#  Defer, Panic, Recover
* I used defer to set up a recovery function at the start of main().
* If the program encounters an invalid number in the input (e.g., abc), it triggers a panic.
* The recover() inside the deferred function catches the panic, prints a error message (Recovered from panic: ...), and prevents the program from crashing.
# Conversion of Types
* Input from -nums is a string (e.g., "1,2,3").
* splitted the string using strings.Split(*nums, ",") to get individual number strings.
* Each string is converted to an integer using strconv.Atoi(p).
* If conversion fails (non-numeric input), the program panics, which is then handled by recover().
