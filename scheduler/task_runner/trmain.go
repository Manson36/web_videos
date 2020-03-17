package task_runner

import (
	"time"
)

/*
如何写task Runner:
//task_runner这个包是一个具有独立功能的包，可以独自运行的包，而不是纯的方法调用的包
1.首先我们加一个task_runner_main：因为在task_runner中我们既有dispatcher，又有executor，然后我们要将所有的这些任务集中起来。
	我们会再trmain.go中将它们初始化，并将它们跑起来；然后在外面的main中将trmain.go初始化并且启动起来
2.我们需要一个原生的task的包，里面是我们解耦的dispatcher和executor中真正跑的东西
3.我们加一个runner的包，这是比较重要的部分，整个逻辑都在这里面
4.加defs 包
*/

/*timer:
1.首先我们要 setup
2. start 它就会一直跑下去，跑的过程中{Trigger -> task -> runner}
3. 所以整个业务的流程是：timer -> task -> runner(longLived)
*/
/*项目的整个流程：
user -> api service -> delete video
api service ->scheduler -> write video deletion record
timer -> runner -> read wvdr -> exec ->delete video from file
*/

type Worker struct {
	ticker *time.Ticker
	runner *Runner
}

func NewWorker(interval time.Duration, r *Runner) *Worker {
	return &Worker{
		ticker: time.NewTicker(interval * time.Second),
		runner: r,
	}
}

func (w *Worker) startWorker() {
	//这里我们使用for -select的方式，而不是使用for c = range w.ticker.C{} 的方式：
	//因为ticker.C没有Close，它会一直阻塞在这里，一直等待新的case过来，可以达到定时器的目的，但是这里面有一个非常大的问题：
	//range的过程是一个同步的过程，并不向select是一个non_block的过程，它是一个block的模式，所以每次执行的时候会同步这个过程，
	// 会消耗几毫秒的时间，所以当定时器跑几个小时之后，你就会发现误差越来越大，因为每次执行的时候都会消耗for的执行时间，
	// 大家千万不要for_range的方式来取定时器的时间，这是非常错误的
	for {
		select {
		case <-w.ticker.C:
			go w.runner.StartAll()
		}
	}
}

func Start() {
	//Start video cleaning
	r := NewRunner(3, true, VideoClearDispatcher, VideoClearExecutor) //这里调用的函数名，没有加括号，加括号是调用函数的返回值
	w := NewWorker(3, r)
	go w.startWorker()
	//other somethings start
	//r1: = ...
	//	w2:= ...
}
