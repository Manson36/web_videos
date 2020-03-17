package task_runner

import (
	"errors"
	"github.com/web_videos/scheduler/dbops"
	"log"
	"os"
	"sync"
)

//我们task部分是定制化比较强的部分，它和runner最大的不同是：runner的东西是可以复用的。
//这个项目的task主要完成延时删除的事情，我们会分为两个部分：
// 1.diapatcher：它回去读数据库里要删除的信息，放到dataChan0中
//2. executor:通过dataChan中的信息真正的把视频删掉

/*流程：
api ->delete_video_id -> sql
dispatcher -> sql -> delete_video_id -> dataChannel
executor -> dataChan -> delete_video_id -> delete video
*/

func deleteVideo(vid string) error {
	err := os.Remove("./videos/" + vid)
	if err != nil && os.IsNotExist(err) {
		log.Println("Deleting video error:", err.Error())
		return err
	}
	return nil
}

func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Println("Video clear dispatcher error :", err)
		return err
	}

	if len(res) == 0 {
		return errors.New("all tasks finished")
	}

	for _, id := range res {
		dc <- id
	}
	return nil
}

func VideoClearExecutor(dc dataChan) error {
	//我们需要把所有的error一起传出来，而不是一个错误传一次：sync.Map是线程安全的，所以多个go routine往里面写的时候都是安全的
	errMap := &sync.Map{}
	var err error
forloop:
	for {
		select {
		case vid := <-dc:
			//这里有一个非常需要注意的一点，在这个case中，我们把vid作为参数传进go func()中，而不是直接使用vid：
			// go func() {(vid)}()是一个闭包，我们在闭包中调用一个go routine的时候，
			//实际上会拿到一个它瞬时的状态，而不会将它的状态保存，只有将参数传给go routine的时候，才构成一个完整的闭包。
			go func(id interface{}) {
				if err := deleteVideo(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
				if err := dbops.DelVideoDeletionRecord(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
			}(vid)
		default:
			break forloop
		}
	}
	errMap.Range(func(k, v interface{}) bool {
		err = v.(error)
		if err != nil {
			return false
		}
		return true
	})

	return err
}
