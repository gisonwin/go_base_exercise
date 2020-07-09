package main

import (
	"log"
)

/**
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/6/28 11:35
 */

func main() {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dial error :",err)
	}
	var reply string
	err = client.Hello("safe client", &reply)
	if err != nil {
		log.Fatal("Call error",err)
	}
	println(reply)
}