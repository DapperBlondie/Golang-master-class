package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main()  {

	// writing to the stdout & stderr
	_, err := io.WriteString(os.Stdout, "\nHello, Go !")
	if err != nil {
		fmt.Println("an Error occurred !")
		os.Exit(1)
	}

	io.WriteString(os.Stderr, "This is not an error !")

}

func divisionByZeroError(a int, b int) (int, error) {

	if b == 0 {
		diverr := errors.New("the second input b is zero")
		return 0, diverr
	}

	return a/b, nil

}
