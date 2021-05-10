package main
// #cgo CFLAGS: -I${SRCDIR}/callClib
// #cgo LDFLAGS: ${SRCDIR}/callC.a
// #include <stdlib.h>
// #include <callC.h>
import "C"

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

type Person struct {
	Name string
	Age int8
	Id int16
}

func createNewPStruct(name string, age int8, id int16) *Person {

	if age > 100  {
		age = 100
	}

	return &Person{
		Name: name,
		Age:  age,
		Id:   id,
	}
}

func main()  {

	/*C.cHello()
	myMessage := C.CString("This is Mikhalis!")
	defer C.free(unsafe.Pointer(myMessage))
	C.printMessage(myMessage)*/

	fmt.Println("Hello Go :)")
	d1()
	d2()
	d3()

	fmt.Println("\nMachine environment variables : ")
	machineEnv()

	var arr [4]int = [4]int{}
	for i:=0;i<len(arr);i++ {
		arr[i] = i*5 + 83
	}

	fmt.Println("The arr elements : ", arr)

	timeFormat()

	person := createNewPStruct("Dapper", 20, 9832)
	fmt.Println("Alireza : ", *person)

	testPerson := new(Person)
	fmt.Println("Test Person : ", testPerson)

	var n1, n2 int = 10, 23
	fmt.Println("num1, num2 : ", n1, n2)
	n1, n2 = n2, n1
	fmt.Println("num1, num2 : ", n1, n2)

	return
}

func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
	return
}

func d2()  {

	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
	fmt.Println()
	return
}

func d3()  {

	for i := 3; i>0;i-- {
		defer func(i int) {
			fmt.Print(i, " ")
		}(i)
	}
	fmt.Println()
	return
}

func machineEnv()  {

	fmt.Println("The Kernel name : ", runtime.GOOS)
	fmt.Println(runtime.Version())
	fmt.Println("Number of CPU : ", runtime.NumCPU())

	verstr := strings.Split(runtime.Version(), ".")
	fmt.Println("Actual Version : ", uint(verstr[0][2]), verstr[1], verstr[2])
}

func timeFormat()  {

	t := time.Now()
	fmt.Println("Epoch Time: ", t.Unix())
	fmt.Println("RFC Format : ", t.Format(time.RFC822))

	time.Sleep(time.Second*3)

	fmt.Println("GoodBye")
}