package main

import "log"

/*流控模块：
1.什么是流控？
在我们整个网站过程中，在网站server online的时候，会有用户或者攻击者 有意或无意的向网站(server)发起请求，当请求达到一定数量的时候
会导致网站server的链接数不够，链接数不够还事小，但是如果他不仅将你的链接数消耗完，并且把你的带宽占完之后，那么你的server就处于不可用的状态，
还有一种更可怕的情况是，当你的 消耗完之后，你的系统可能会 crash，为了保护我们的server、保护系统，我们需要做流控。
流控用英语来说叫 run limiter,一般都用一种算法 bucket-token算法。
2.bucket-token算法：
有一个bucket，它包含 n个token，request 拿到token 之后才算是真正的可以进入服务的request；当request拿到它想要的结果，返回response之后，
那么request的生命周期在我们的server中也已经结束了，此时将token还回去，放回bucket，以便给下一个request使用。以此来限制访问我们server的数量
3.bucket-token的使用：
通常大家会想bucket-token直接给数组赋值就可以了，但实际上并没有这么简单。当我们使用数组或者一段普通的变量的话，由于我们的Handler
(在go中，每一个Handler都是新起了一个goroutine，而goroutine它们之间是并发的)，在并发的情况下，很可能会有不同handler访问同一个变量的情况，
如果我们的bucket是一个普通的内存中的数据块的话，同一时间访问它的话肯定会出现问题，我们为了避免竞争，肯定要给它加锁，来保证它的同步；
加了锁之后，它的性能势必会下降，影响性能是小，实际上这并不是go 设计出来用来处理并发问题的理想做法，go 常使用channel。(shared channel
instead of shared memory.)通过共享通道来同步线程之间的信息，而不是使用共享内存。
*/
type ConnLimiter struct {
	concurrentConn int
	bucket         chan int
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc),
	}
}

/*
buffer channel和no_buffer channel 的区别：
no_buffer channel相当于是一个同步channel，当这个channel写进数据的时候，如果没有别的goroutine来读这个channel，这个channel永远都是阻塞的，
就不会有别的东西往里写，对程序来说不是特别好用。在通常的情况下，我们会使用一些buffer channel来做消息之间的同步，有了buffer就可以保证
在一定的缓冲区间之内做一些消息同步的事情。
*/

func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Println("Reached the rate limitation")
		return false
	}

	cl.bucket <- 1
	return true
}

func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.bucket
	log.Println("New connection coming:", c)
}
