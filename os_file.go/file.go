package main

import (
	"bufio"
	"fmt"
	"io"

	//"io"
	"os"
)

func main() {
	createFile, err := os.Open("./File1.txt")
	secondFile, secFile_err := os.OpenFile("/home/abhiral/Desktop/go/os_file.go/File2.txt", os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		panic("Err in creating file")
	}
	if secFile_err != nil {
		panic("Error in opening the file")
	}
	defer secondFile.Close()
	file1Reader := bufio.NewReader(createFile)
	secondWriter := bufio.NewWriter(secondFile)

	// b, err := secondWriter.WriteString("Hello iam writing to file2")
	//fmt.Println(b)
	secondWriter.Flush()
	for {
		data, _, err := file1Reader.ReadLine()
		if err == io.EOF {
			fmt.Println("Completed Reading file....")
			break
		} else if err != nil {
			fmt.Println("An unexpected err occured while reading the file")

			break
		}
		secondWriter.WriteString(string(data) + " \n")
	}
	secondWriter.Flush()
	defer createFile.Close()

	// _, writeFileErr := fileWrite.WriteString("Hello iam writing something idk what is that but iam writing something , dammn write write write")
	// if writeFileErr != nil {
	// 	panic("Error while writing in file")
	// }
	// fileWrite.Flush()

}
