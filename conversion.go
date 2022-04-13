package main

import "fmt"

func main() {
	var value32 int32 = 100000
	var value64 int64 = int64(value32)

	fmt.Println(value32)
	fmt.Println(value64)

	var name string = "Fakhrul";
	var f byte = name[0];
	var fString = string(f);
	fmt.Println(fString)

}