package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var no int = 100
	fmt.Println(reflect.TypeOf(no))

	var intStr string = "100"
	fourBaseEightBitInt, _ := strconv.ParseInt(intStr, 4, 8) // becomes no 16 and in\

	tenBaseSixteenBitInt, _ := strconv.ParseInt(intStr, 10, 16) // no 100, and int64
	fmt.Println(reflect.TypeOf(fourBaseEightBitInt))
	fmt.Println(reflect.TypeOf(tenBaseSixteenBitInt))
	fmt.Println(tenBaseSixteenBitInt)
	fmt.Println(tenBaseSixteenBitInt)

	name := "unch"

	fmt.Println(name)
	// hehe
	// tambah

}
