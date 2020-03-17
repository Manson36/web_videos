package task_runner

/*
runner.go逻辑：
1.首先我们要有一个runner的对象，在runner中我们会跑一个常驻的任务(startDispatcher),整个任务会长时间的去等待这的一个runner的channel，
这个channel分为两部分(control channel/data channel),control channel是用来dispatcher和executor来相互交换信息，来提醒对方；data channel 是真正的数据部分
*/
type Runner struct {
	Controller controlChan
	Error      controlChan
	Data       dataChan
	dataSize   int
	longLived  bool
	Dispatcher fn
	Executor   fn
}

func NewRunner(size int, longLived bool, d, e fn) *Runner {
	return &Runner{
		Controller: make(controlChan, 1),
		Error:      make(controlChan, 1),
		Data:       make(dataChan, size),
		dataSize:   size,
		longLived:  longLived,
		Dispatcher: d,
		Executor:   e,
	}
}

func (r *Runner) startDispatch() {
	defer func() {
		if !r.longLived {
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}()
	for {

		select {
		case c := <-r.Controller:
			if c == READY_TO_DISPATCH {
				err := r.Dispatcher(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_EXECUTE
				}
			}

			if c == READY_TO_EXECUTE {
				err := r.Executor(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_DISPATCH
				}
			}
		case e := <-r.Error:
			if e == CLOSE {
				return
			}
		default:

		}
	}
}

func (r *Runner) StartAll() {
	r.Controller <- READY_TO_DISPATCH
	r.startDispatch()
}
