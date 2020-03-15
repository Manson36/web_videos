package session

import "sync"

/*
session我们的主要操作是：
1.在我们服务起来的时候，从DB中拉取Session，把所有存在的Session全部load到我们Api节点的Cache中
2.当有新用户注册或老用户登录的时候，我们需要费当前的用户分配一个SessionID，里面包含Session的信息，这时候就需要一个产生Session的方法
3.当我们在校验的时候，Session可能会过期，这时需要给用户返回一个过期或没有过期的状态，用来判断当前这个用户是否是合法、已经登录的用户
*/

//这是在Go 1.9版本之后加入的内容，之前的Map不支持并发，读写的时候需要加锁
//sync.Map 在并发读的时候非常稳定，但是在并发写的时候会有一些问题，不会出现它的 key conflicts(冲突)，所以每次写都需要加全局锁，这时非常耗时的
var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func LoadSessionsFromDB() {

}

func GenerateNewSessionID(userName string) string {

}

func IsSessionExpired(sessionId string) (string, bool) {

}
