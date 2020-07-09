package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

/**
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/6/28 22:43
 */
func echo(wr http.ResponseWriter,r *http.Request){
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wr.Write([]byte("echo error"))
		return
	}
	writeLen, err := wr.Write(msg)
	if err != nil || writeLen != len(msg) {
		log.Println(err,"write len:",writeLen)
	}
}
func main() {
	http.HandleFunc("/",echo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}