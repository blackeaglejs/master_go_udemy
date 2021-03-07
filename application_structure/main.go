// Package declarations are required in all go programs.
// They let you name where the main code compilation should happen.
package main

import "fmt"

// Constants and variables can be declared anywhere before the main function.
// They can also be set to a value using basic equal syntax as shown here.
const secondsInHour = 3600

// The main function is present in all main.go files - it's the main body for code execution.
func main() {
	fmt.Println("Hello Go World!")
	// The := is used to implicitly declare variables.
	distance := 60.8

	fmt.Printf("The distance in miles is %f \n", distance*0.62137)
}

// go run vs go build vs go install
// go run runs any executable file that you give it without generating a binary.
// go build will take the files in your existing directory, then compile it into a binary for use.
// go install will take the files in your directory, compile them, and then place the binary in the /bin directory.
// you can build for whichever operating system you desire.
// adding GOOS="windows,linux,darwin,etc" will build for whichever of these three operating systmes.
// adding GOARCH="amd64,arm,etc" will build for whichever architecture you prefer.
// adding the -o flag plus a filename will let you name the executable.

// gofmt uses go language rules to clean up your files.
// your editor can be configured to automatically run gofmt on a file.
// the -w operator will overwrite your file with the formatted version.
// the -l operator will let you do it to an entire directory.
