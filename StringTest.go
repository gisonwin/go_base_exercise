package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "123"

	//str[0] = 'a'
	var da []byte = []byte(str)
	da[0]='a'
	da[1]='b'
	da[2]='c'
	fmt.Println(string(da))
 //string convert to int
	data,err := strconv.Atoi(str)
	println(data)
	if ! (data > 0) {
		fmt.Println(err)
	}

	fmt.Printf("%v",str)

}