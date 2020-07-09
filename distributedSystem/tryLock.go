package main

import "sync"

/** 逻辑限定每个goroutine只有成功执行了Lock才会继续执行后续逻辑,因此在Unlock时可以保证Lock结构体中channel一定是空,从而
不会阻塞,也不会失败.代码使用大小为1的channel来描述tryLock,理论上还可以使用标准库中的CAS来实现相同的功能且成本更低.
在单机系统中,trylock并不是一个好选择.因为大量goroutine抢锁可能会导致CPU无意义的资源浪费.有一个名词描述这种抢锁场景:活锁
活锁是指程序看起来在正常运行,但实际上CPU周期被浪费在抢锁,而非执行任务上,从而程序整体执行效率低下.活锁问题定位起来要麻烦
很多.所在在单机场景下不建议使用这种锁.
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/7/2 21:44
 */
//Lock try lock
type Lock struct {
	c chan struct{}
}

//NewLock generate a try lock
func NewLock() Lock {
	var l Lock
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}
	return l
}

//Lock try lock,return lock
func (l Lock) Lock() bool {
	select {
	case <-l.c:
		return true
	default:

	}
	return false
}

//Unlock ,unlock the try lock
func (l Lock) UnLock() {
	l.c <- struct{}{}
}

var cnt int

func main() {
	lock := NewLock()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !lock.Lock() {
				println("lock failed")
				return
			}
			cnt++
			println("current counter", cnt)
			lock.UnLock()
		}()
		wg.Wait()
	}
}
