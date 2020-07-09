package main

import (
	"log"
	"net"
	"net/rpc"
)

/**
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/6/28 11:56
 */
//type Helloservice struct {
//
//}
//func (p *HelloService) Hello(request string,reply *string) error{
//	*reply = "hello:" +request
//	return nil
//}
func main() {
	RegisterHelloService(new(HelloService))
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("Listen error ", e)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
