package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Value interface {
	String() string
	Set(v string) error

}

func main()  {

	dashI := flag.Int("i", 1, "i")
	dashD := flag.Int("d", 1, "d")
	flag.Parse()

	valueI := *dashI
	valueD := *dashD

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Not enough Args :(( Run again")
		os.Exit(1)
	}else {
		number, err := strconv.Atoi(arguments[1])
		if err != nil {
			fmt.Println(valueI, valueD)
			fmt.Println("Bad Argument :(( Run again")
			os.Exit(1)
		}else {

			number = number + valueI - valueD
			fmt.Println(valueI, valueD)
			fmt.Println("Your Entered Number + (-I) - (-D) : ", number)
		}
	}

}
