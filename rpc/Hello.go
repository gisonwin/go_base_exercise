package main

import (
	"log"
	"net"
	"net/rpc"
)

/**
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/6/27 20:59
 */
type HelloService struct{
	conn net.Conn

}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}
func main() {
	rpc.RegisterName("HelloService", new(HelloService)) //rpc.Register函数调用会将对象类型
	//中所有满足RPC规则的对象方法注册为RPC函数,所有的注册方法会放在HelloService服务空间下.然后我们
	//建立一个唯一的TCP链接,并且通过rpc.ServeConn函数在该TCP链接上为对方提供RPC服务.
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}
	rpc.ServeConn(conn)
}
