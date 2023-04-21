package main

/*
流量限制的手段有很多，最常见的：漏桶、令牌桶两种：

漏桶是指我们有一个一直装满了水的桶，每过固定的一段时间即向外漏一滴水。如果你接到了这滴水，那么你就可以继续服务请求，如果没有接到，那么就需要等待下一滴水。
令牌桶则是指匀速向桶中添加令牌，服务请求时需要从桶中获取令牌，令牌的数目可以按照需要消耗的资源进行相应的调整。如果没有令牌，可以选择等待，或者放弃。
*/
import (
	"fmt"
	"time"
)

var (
	tokenBucket  = make(chan struct{}, capacity)
	fillInterval = time.Millisecond * 10
	capacity = 100
)

func main() {



	fillToken := func ()  {
		ticker := time.NewTicker(fillInterval)

		for {
			select {
			case <-ticker.C:
				select {
				case tokenBucket <- struct{}{}:
				default:
				}
				fmt.Println("current token cnt:", len(tokenBucket), time.Now())
			}
		}
	}

	go fillToken()
	time.Sleep(time.Hour)
}

func TakeAvailable(block bool) bool {
	var takenResult bool
	if block {
		select {
		case<-tokenBucket:
			takenResult = true
		}
	} else {
		select {
		case <-tokenBucket:
			takenResult = true
		default:
			takenResult = false
		}
	}

	return takenResult
}

// 令牌桶每隔一段固定的时间向桶中放令牌，如果我们记下上一次放令牌的时间为 t1，和当时的令牌数 k1，放令牌的时间间隔为 ti，每次向令牌桶中放 x 个令牌，令牌桶容量为 cap。现在如果有人来调用 TakeAvailable 来取 n 个令牌，我们将这个时刻记为 t2。在 t2 时刻，令牌桶中理论上应该有多少令牌呢？


/*
cur = k1 + ((t2 - t1)/ti) * x
cur = cur > cap ? cap : cur
*/