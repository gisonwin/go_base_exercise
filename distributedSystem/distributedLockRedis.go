package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

/** 分布式场景下如果需要抢占逻辑,我们可以使用redis的setnx命令.
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/7/2 22:25
 */
func incr() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	var lockKey="count_lock"
	var counterKey="counter"
	//lock
	resp := client.SetNX(lockKey, 1, time.Second*5)
	lockSucess, err := resp.Result()

	if err != nil || !lockSucess {
		fmt.Println(err,"lock result:",lockSucess)
		return
	}
	//counter ++
	getResp := client.Get(counterKey)
	cntValue, err := getResp.Int64()
	if err == nil {
		cntValue++
		resp := client.Set(counterKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			println("set value error")
		}
	}
	println("current counter is ",cntValue)

	delResp := client.Del(lockKey)
	unlockSuccess, err := delResp.Result()
	if err == nil && unlockSuccess >0 {
		println("unlock success")
	}else {
		println("unlock failed ",err)
	}

}
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				incr()
			}()
	}
	wg.Wait()
}
/**
setnx很适合在高并发场景下用来抢夺一些"唯一"资源.比如交易撮合系统中卖家发起订单,而多个买家会对其进行并发争抢.这种场景
我们没有办法依赖具体时间来判断先后,因为不管是用户设备的时间还是分布式场景下各台机器的时间都没有办法在合并后保证正确
的时序,哪怕是同一个机房的集群,不同机器的系统时间可能也会有细微的差别.所以我们要依赖于这些请求到达redis节点的顺序来做
正确的抢锁操作.如果用户的网络环境比较差就只能自求多福了.
 */
/**
如何选择合适的锁:
业务还在单机就可以搞定的量级时,可以使用任意的单机锁方案
如果发展到了分布式服务阶段,但业务规模不大,qps很小情况下,哪种分布式锁方案都差不多,如果公司内业务已使用了zookeeper,etcd
或redis集群,就尽量在不引入新技术栈情况下满足业务需求
业务发展到一定量级的话,需要从多方面来考虑,首先是你的锁是否在任务恶劣条件下都不允许数据丢失,如果不允许,就不要使用redis的
setnx锁.
对锁数据的可靠性要求极高的话,只能使用etcd或zookeeper这种强一致性的保证数据可靠性的锁方案,但可靠的背面往往是较低的吞吐量和
较高的延迟.需要根据业务的量级对其进行压力测试,以确保分布式锁所使用的etcd或zookeeper集群可以承受得住实际的业务请求压力.
需要注意的是etcd或zookeeper集群是没有办法通过增加节点来提高其性能的.要对其进行横向扩展,只能增加搭建多个集群来支持更多的
请求.这会加大运维和监控的风险.
多个集群可能需要引入proxy,如果没有proxy,就需要业务去根据某个业务id来做分片.如果业务已上线情况下做扩展,还要考虑数据的
动态迁移.这些都是不容易的.在选择具体方案时还需要多加思考,对风险早做预估.
 */