package task_runner

/*
这三个分别代表什么呢：这三个都是在controlChan中的消息
1.READY_TO_DISPATCH 当这个消息被我们的dispatcher收到之后，它就会开始做它的事情，然后给dataChan下发它的数据，
2.READY_TO_EXECUTE 当数据下发之后，就会将这个消息发送给executor，它就会从dataChan中读取数据，然后去做它的事情
3.CLOSE 无论是dispatcher还是executor任何一个出了问题或者说没有任务可做的时候就会发送一个CLOSE，我们就会把我们传入的任务取消掉
*/
const (
	READY_TO_DISPATCH = "d"
	READY_TO_EXECUTE  = "e"
	CLOSE             = "c"
)

type controlChan chan string
type dataChan chan interface{}

//这是方法：dispatcher和executor
type fn func(dc dataChan) error
