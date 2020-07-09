package main

import (
	"net"
	"net/rpc"
	"time"
)

/**

反向RPC内网服务将不再主动提供TCP监听服务,而是首先主动链接到对方的TCP服务器.然后基于每个建立的TCP
链接向对方提供RPC服务.
而RPC客户端需要在一个公共的地址提供一个TCP服务,用于接受RPC服务器的链接请求.
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/6/28 14:46
 */
func main() {
	rpc.Register(new(HelloService))
	for {
		conn, _ := net.Dial("tcp", "localhost:1234")
		if conn == nil {
			time.Sleep(time.Second)
			continue
		}
		rpc.ServeConn(conn)
		conn.Close()
	}
}
