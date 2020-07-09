package main

import "sync"

/**
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/7/2 21:40
 */
var counter int

func main() {
	var wg sync.WaitGroup
	var l sync.Mutex
	for i := 0; i<1000;i++  {
		wg.Add(1)
		go func() {
			defer wg.Done()
			l.Lock()
			counter++
			l.Unlock()
		}()
	}
	wg.Wait()
	println(counter)
}