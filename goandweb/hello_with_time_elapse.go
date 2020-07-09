package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

/**
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/6/29 10:39
 */
var logger = log.New(os.Stdout,"",0)

func hello(wr http.ResponseWriter,r *http.Request){
	wr.Write([]byte("hello"))
}
func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		next.ServeHTTP(wr, r)
		timeElapsed := time.Since(timeStart)
		logger.Println(timeElapsed)
	})
}
func main() {
	http.Handle("/",timeMiddleware(http.HandlerFunc(hello)))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Fatal(err)
	}
}