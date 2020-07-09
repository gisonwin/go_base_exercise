package main

import (
	"net/http"
	"time"
)

/**
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/6/29 10:42
 */
func helloHandler(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("hello"))
}

func showInfoHandler(wr http.ResponseWriter, r *http.Request) {

}
func showEmailHandler(wr http.ResponseWriter, r *http.Request) {

}
func showFriendsHandler(wr http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()
	wr.Write([]byte("your friends is tom and alex"))
	timeElapsed := time.Since(timeStart)
	logger.Println(timeElapsed)
}
func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/info/show", showInfoHandler)
	http.HandleFunc("/emain/show", showEmailHandler)
	http.HandleFunc("/friends/show", showFriendsHandler)
}
