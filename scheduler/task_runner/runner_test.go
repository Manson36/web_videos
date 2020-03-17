package task_runner

import (
	"errors"
	"log"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	d := func(dc dataChan) error {
		for i := 0; i < 30; i++ {
			dc <- i
			log.Printf("Dispatcher send :%v", i)
		}
		return nil
	}

	e := func(dc dataChan) error {
	forloop:
		for {
			select {
			case d := <-dc:
				log.Printf("executor retrieved :%v", d)
			default:
				break forloop
			}
		}
		return errors.New("execute finish") //nil, 如果这里不返回error，那么会不断执行直到3s时间到
	}

	runner := NewRunner(30, false, d, e)

	/*这里我们没有直接调用startAll，而是使用go routine，为什么？
	在StartAll中包含startDispatch函数，它内部有for的死循环(死循环中加forloop也可以破解)，如果我们不在后台以go routine的形式启动，就会一直blocking(
	不断地写入数据，读取数据)，就不会执行到下面的time.sleep
	在go中我们有时会需要block整个进程，不让他运行结束，那么就有两种方式：一种是for的死循环，另一个是声明一个没有buffer的channel，不给里面写东西，直接去读
	*/
	go runner.StartAll()
	time.Sleep(3 * time.Second)
}
