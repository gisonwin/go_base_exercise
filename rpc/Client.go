package main

import (
	"log"
	"net/rpc"
)

/**
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/6/27 21:56
 */
func main() {
	client, e := rpc.Dial("tcp", "localhost:1234")
	if e != nil {
		log.Fatal("dialing:", e)
	}
	var reply string
	err := client.Call("HelloService.Hello", "gisonwin", &reply)
	if err != nil {
		log.Fatal(err)
	}
	println(reply)
}
