package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

const (
	Saturday = iota+1
	Sunday
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
)

func main()  {

	var  myname string = "Dapper"
	var myage int8 = 20

	pathModule()
	fmt.Println("The result of the Multiple : ", Multiple())
	fmt.Println("The path of the execution : ", cmdArgs())

	//getArgs()
	var name string = "Dapper"
	var expr string = "Hey"

	repeatedGreeting(&name, &expr)
	fmt.Printf("My Name is %s, I am %T variable", myname, myage)

	var result float64 = 1.0/10.0
	fmt.Printf("The result of the 1.0/10.0 is %.3f \n", result)

	var input byte = 98
	byteRepresentation(&input)

	fmt.Println("The Monday code is : ", Monday)

	var color1, color2 string
	color2 = "red"
	color1 = "green"
	fmt.Println("The result : ", color1 == "green" || color2 == "red")

	arg := os.Args[1]
	argf, err := strconv.ParseFloat(arg, 32)

	if err != nil {
		fmt.Println("We have some problems with conversion arg to Float32 ...")
		os.Exit(1)
	}else {
		fmt.Println("the result * 10.01 : ", argf*10.01)
	}

}


func aliasTypes()  {

	var id rune = 8921 // is an alias for the int32
	var myByte byte = 78
	fmt.Println("The rune type is equal to int32 : ", id, " ,", myByte)

}

func byteRepresentation(input* byte) {

	fmt.Printf("The representation of a byte in GO lang : %b = %d\n", *input, *input)

}

func repeatedGreeting(name* string, expr* string) string {

	fmt.Println("The address of the expr : ", expr, " : ", &expr)
	fmt.Println("The address of the name : ", name, " : ", &name)
	return strings.Repeat(*expr, 2) + " " + *name
}

func Multiple() float32 {

	number1 := 367
	var number2 float32 = 89.81

	return number2 * float32(number1)

}

func pathModule()  {

	dir, file := path.Split("/alirezaFiles/file.txt")
	fmt.Println("Dir : ", dir, "\nFile : ", file)

}

func cmdArgs() string {
	return os.Args[0]
}

func getArgs()  {

	if len(os.Args[1]) != 0 {
		fmt.Println("Hello, ", os.Args[1])
	}else {
		fmt.Println("you entered nothing.")
	}

}