package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	basicFileOperations()
	writingBytesToFiles()
	usingBufferedWriter()
	usingReader()
	usingScanner()
	usingUserInput()
}

func basicFileOperations() {
	// We use the os package - it's unix-like on operating systems, but golike when we're coding with it.

	// Creating a new file.
	// os.Create() creates a file with the supplied file. You can store this information in
	newFile, err := os.Create("a.txt")

	// Note: Go does not have exceptions like other programmming languages.
	// The idiomatic way to deal with errors is to use the log package. log.Fatal() will log the error and kill the process.
	// error handling
	if err != nil {
		// log the error and exit the program
		log.Fatal(err) // the idiomatic way to handle errors

	}

	// The truncate method will remove all contents from the file and make it zero bytes.
	err = os.Truncate("a.txt", 0) //0 means completely empty the file.

	// error handling
	if err != nil {
		log.Fatal(err)
	}
	// The Open() function will open a file for editing.
	file, err := os.Open("a.txt") // open in read-only mode

	// error handling
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	//OPENING a FILE WITH MORE OPTIONS
	file, err = os.OpenFile("a.txt", os.O_APPEND, 0644)
	// We can Use opening attributes individually or combined
	// using an OR between them
	// e.g. os.O_CREATE|os.O_APPEND
	// or os.O_CREATE|os.O_TRUNC|os.O_WRONLY
	// os.O_RDONLY // Read only
	// os.O_WRONLY // Write only
	// os.O_RDWR // Read and write
	// os.O_APPEND // Append to end of file
	// os.O_CREATE // Create is none exist
	// os.O_TRUNC // Truncate file when opening

	// error handling
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	// The Close() function will close a file for use.
	newFile.Close()

	// The FileInfo type stores file information.
	var fileInfo os.FileInfo
	// The Stat() function can store info in FileInfo type and will get stats about the file.
	fileInfo, err = os.Stat("a.txt")

	p := fmt.Println
	// The Name() function returns the name of the file.
	p("File Name:", fileInfo.Name()) // => File Name: a.txt
	// The Size() function will return file size in bytes.
	p("Size in bytes:", fileInfo.Size()) // => Size in bytes: 0
	// The ModTime() function will return the time the file was last changed.
	p("Last modified:", fileInfo.ModTime()) // => Last modified: 2019-10-21 16:16:00.325037748 +0300 EEST
	// The IsDir() function will return whether or not the file is a directory.
	p("Is Directory? ", fileInfo.IsDir()) // => Is Directory?  false
	// The Mode() function will return the permissions of the file.
	p("Pemissions:", fileInfo.Mode()) // => Pemissions: -rw-r-----
	// The IsNotExist(err) function will check if the error is "File does not exist."
	fileInfo, err = os.Stat("b.txt")
	// error handling
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("The file does not exist")
		}
	}
	// The Rename() function takes two paths to a file, and changes the old path to the new path.
	oldPath := "a.txt"
	newPath := "aaa.txt"
	err = os.Rename(oldPath, newPath)
	// error handling
	if err != nil {
		log.Fatal(err)
	}
	// The Remove() function will delete the file from the filesystem.
	err = os.Remove("aa.txt")
	// error handling
	if err != nil {
		log.Fatal(err)
	}
}

func writingBytesToFiles() {
	// Reader and Writer functions simply read and write bytes to a file. They don't care where they come from, they just read/write.
	// Open File.
	file, err := os.OpenFile(
		"b.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // will defer closing to the end of hte function execution.

	byteSlice := []byte("I learn Golang!")
	byteWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes Written: %d\n", byteWritten)

	// Faster is the ioutil package
	bs := []byte("Go Programming is cool!")
	err = ioutil.WriteFile("c.txt", bs, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Depending on how complex your needs are, you may choose to just use the os package for writing files because more packages means a bigger executable.
}

func usingBufferedWriter() {
	// You write with a buffer in memory before storing in a file.
	// Disks are slow and RAM is much faster - so it's faster to write a file in memoery.
	// You can also do a lot of manipulation before putting it in a file.

	// Opening the file for writing
	file, err := os.OpenFile("my_file.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Creating a buffered writer from the file variable using bufio.NewWriter()
	bufferedWriter := bufio.NewWriter(file)

	// Writing a byte slice to a file using buffered writer.
	bs := []byte{97, 98, 99}
	bytesWritten, err := bufferedWriter.Write(bs)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written to buffer (not file): %d\n", bytesWritten)
	// Writing a string to a file using buffered writer.

	// You can check how many bytes are buffered and available
	bytesAvailable := bufferedWriter.Available()
	log.Printf("Bytes available in buffer: %d\n", bytesAvailable)
	if err != nil {
		log.Fatal(err)
	}
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)

	// You use the Flush() method to dump into a file.
	bufferedWriter.Flush()
}

func usingReader() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// make a byteslice specifying the number of bytes.
	byteSlice := make([]byte, 2)

	// Use the io.ReadFull method to read a file, specifying the byteSlice.
	numberBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", numberBytesRead)
	log.Printf("Data read: %s\n", byteSlice)

	fmt.Println(strings.Repeat("#", 20))
	// You can use the ioutil.ReadAll() function with a file to read an entire file.
	file, err = os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data as string: %s\n", data)
	fmt.Println("Number of bytes read:", len(data))
	// If you need to read the entire file into a byteslice you can use the ioutil.ReadFile() method to open, read, and close the entire file.
	data, err = ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Data read: %s\n", data)
}

func usingScanner() {
	// Use a delimiter to read a file.
	// Use the bufio.Scanner() method to do it.

	file, err := os.Open("my_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Initialize a scanner
	scanner := bufio.NewScanner(file) // reads line by line since delimiter not provided.

	success := scanner.Scan()
	if !success {
		err = scanner.Err()
		if err == nil {
			log.Println("Scan was completed and it returned EOF")
		} else {
			log.Fatal(err)
		}
	}

	fmt.Println("First line found:", scanner.Text())

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// You can also scan by word or rune.
}

func usingUserInput() {
	// Create new Scanner from bufio.
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%T\n", scanner)

	scanner.Scan()

	text := scanner.Text()
	bytes := scanner.Bytes()

	fmt.Println("Input text:", text)
	fmt.Println("Input bytes:", bytes)

	for scanner.Scan() {
		text = scanner.Text()
		fmt.Println("You entered:", text)
		if text == "exit" {
			fmt.Println("Existing the scanning ...")
			break
		}
	}
}
