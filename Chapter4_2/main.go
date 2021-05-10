package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main()  {

	stringsReader("Hello")
	bytesReader("Hello")

	reader := strings.NewReader("Operation Complete !")
	num, _ := reader.WriteTo(os.Stderr)
	fmt.Println("The number of written chars : ", num)

	return
}

func stringsReader(msg string)  {

	reader := strings.NewReader(msg)

	b := make([]byte, 1)

	for  {
		n, err := reader.Read(b)
		if err == io.EOF {

			fmt.Println("EOF occurred !")
			break
		}else if err != nil {

			fmt.Println("this error occurred : ", err.Error())
			continue
		}else {

			fmt.Printf("char is : %s, number : %d\n", b, n)
		}
	}

	return
}

func bytesReader(msg string)  {

	reader := bytes.NewReader([]byte(msg))

	b := make([]byte, 1)

	for  {
		n, err := reader.Read(b)
		if err == io.EOF {

			fmt.Println("EOF occurred !")
			break
		}else if err != nil {

			fmt.Println("This error occurred : ", err)
			continue
		}else {

			fmt.Printf("This char %s, num is %d\n", b, n)
		}
	}

	return
}

func bytesBuffer(msg string)  {

	var buffer bytes.Buffer

	_, err := buffer.Write([]byte(msg))
	if err != nil {
		fmt.Println("This error occurred : ", err)
		return
	}else {

		num, err1 := buffer.WriteTo(os.Stdout)
		defer buffer.Reset()
		if err1 != nil {

			fmt.Println("This error occurred during writing into STDOUT : ", err1)
			return
		}else {

			fmt.Printf("we write %d chars into os.Stdout by bytes pkg.", num)
			return
		}
	}
}