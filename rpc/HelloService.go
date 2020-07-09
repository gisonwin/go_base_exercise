package main

import (
	"log"
	"net/rpc"
)

/**
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/6/28 11:20
 */
//服务名字,RPC服务抽象的包路径,并不完全等价Go语言的包路径.
const HelloServiceName = "path/to/pkg.HelloService"

//服务要实现的详细方法列表.
type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}
type HelloServiceClient struct {
	*rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	client, err := rpc.Dial(network, address)
	if err != nil {
		log.Fatal("Dial error ", err)
	}
	return &HelloServiceClient{Client: client}, nil
}
func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}

//注册该类型服务的函数.
func RegisterHelloService(svc HelloServiceInterface) error {
	println("service name: ",HelloServiceName)
	return rpc.RegisterName(HelloServiceName, svc)
}
