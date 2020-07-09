package main

import (
	"io"
	"log"
	"net/http"
)

/**
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/6/30 22:35
 */
func sayHello(wr http.ResponseWriter,r *http.Request)  {
	wr.WriteHeader(200)
	io.WriteString(wr,"hello world")
}
func main() {
	http.HandleFunc("/",sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
}