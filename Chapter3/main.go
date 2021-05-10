package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

type ShapeUtils interface {

	A() int
	P() int

}

type ComplexUtils interface {

	Add(re float32, im float32) *DComplex
	Sub(re float32, im float32) *DComplex
}

type DComplex struct {

	Real float32
	Imag float32

}

type Rectangle struct {

	Width int
	Height int
}

type Sharp struct {

	Width int
	Height int
}

func main()  {

	fmt.Println("random number : ", randomNumberUnixGenerator())

	//channel := make(chan string)

	/*go func() {
		time.Sleep(time.Second*5)
		fmt.Println("Hello there")
		channel <- "Channel Fired up !"
	}()

	go func() {
		time.Sleep(time.Second*2)
		fmt.Println("Hello there")
		channel <- "Channel Fired up ! " + "73882"
	}()*/

	rectangle := Rectangle{
		Width:  12,
		Height: 23,
	}

	sharp := Sharp{
		Width:  78,
		Height: 90,
	}

	fmt.Println(reflect.ValueOf(sharp))
	fmt.Println(reflect.TypeOf(sharp))
	fmt.Println(toString(&sharp))
	fmt.Println(toString(&rectangle))

}

func toString(shape ShapeUtils) string {

	return strconv.Itoa(shape.A())
}

func (rect *Rectangle) A() int {

	return rect.Width * rect.Height
}

func (sharp *Sharp) A() int {

	return sharp.Width * sharp.Height
}

func (rect *Rectangle) P() int {

	return rect.Width*2 + 2*rect.Height

}

func (sharp *Sharp) P() int {

	return sharp.Width*2 + 2*sharp.Height

}

func (dcom *DComplex) Add(re float32, im float32) *DComplex {

	redCom := new(DComplex)
	redCom.Real = dcom.Real + re
	redCom.Imag = dcom.Imag + im

	return redCom
}

func (dcom *DComplex) Sub(re float32, im float32) *DComplex {

	redCom := new(DComplex)
	redCom.Real = dcom.Real - re
	redCom.Imag = dcom.Imag - im

	return redCom
}

func randomNumberUnixGenerator() int {
	SEED := time.Now().UnixNano()
	SRC := rand.NewSource(SEED)
	myRand := rand.New(SRC)

	return myRand.Int()
}