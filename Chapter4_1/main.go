package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// We use this program for read/write into a file


func main()  {

	err := lineByLine("D:/AAlireza/PROJECTS/GO/Golang-master-class/Chapter4_1/testfile.txt")
	fmt.Println("This error occurred : ", err)

	dataBuffer := readBytes("D:/AAlireza/PROJECTS/GO/Golang-master-class/Chapter4_1/testfile.txt")
	fmt.Println(string(*dataBuffer))

	sign := writeBufIO("Hi, How are you ?")
	fmt.Println("Write to the testfiletest.txt was", sign)

	return
}

func lineByLine(filename string) error {

	var err error
	file, err := os.Open(filename)

	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for  {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("EOF occurred !")
			break
		}else if err != nil {
			fmt.Println("This error occurred :(( ", err)
			break
		}
		fmt.Println(line)
	}
	return nil

}

func readBytes(filename string) *[]byte {

	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		fmt.Println("We have some problem with opening file : ", filename)
		os.Exit(1)
	}else {
		buffer := make([]byte, 20)
		num, errB := file.Read(buffer)
		if errB != nil {
			fmt.Println("This problem occurred : ", errB)
			fmt.Println("we red ", num, " bytes ")
			fmt.Println(string(buffer))
			os.Exit(1)
		}else {

			return &buffer
		}
	}

	return nil
}

func writeFPrintF(msg string) bool {

	file, err := os.Create("D:/AAlireza/PROJECTS/GO/Golang-master-class/Chapter4_1/testfiletest.txt")
	if err != nil {

		fmt.Println("problem with creating file : ", err)
		os.Exit(1)
	}else {

		num, err := fmt.Fprintf(file, msg)
		if err != nil {
			fmt.Println("problem with writing into file : ", err)
			fmt.Println("We red ", num, "chars")
			return true
		}
	}

	return false
}

func writeBufIO(msg string) bool {

	file, err := os.Create("D:/AAlireza/PROJECTS/GO/Golang-master-class/Chapter4_1/testfiletest.txt")
	defer file.Close()
	if err != nil {

		fmt.Println("problem with creating file occurred : ", err)
		os.Exit(1)
	}else {

		writer := bufio.NewWriter(file)
		defer writer.Flush()
		num, err1 := writer.WriteString(msg)
		if err1 != nil {

			fmt.Println("Problem with writing your msg into the file : ", err1)
			fmt.Println("We write ", num, " chars")
			os.Exit(1)
		}

		return true
	}

	return false
}
