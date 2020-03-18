#流媒体学习总结
##总结：restful 风格的api：
1. 使用http作为通信协议，使用json作为数据格式
2. 通过不同的METHOD来区分对资源的crud
3. 使用同一资源定位符(rul)设计api
4. 统一接口，无状态，可缓存

###总结：handler
1. 从request中获取body信息，然后将它反序列化为我们想要的字段
2. 然后进行字段的验证 还有session信息的验证
3. 验证通过，进行数据库的操作和session的操作
4. 将执行的结果返回response中

###总结：db ops
1.连接数据库：使用init的函数
2.完成handler的需求

###总结：MiddleWare
* 在router之前，进行校验、鉴权、流控等处理统称为中间件
* 实现：http.Handler接口 只有一个方法ServeHTTP(http.ResponseWriter, *http.Request)，duck type，只要实现了ServeHTTP 即可实现http.Handler
* 在ServeHTTP中进行校验(session)的处理，从r.Header中获取sessionId

###session：
* 是在服务端为用户保存它相应状态的一种机制。是校验处理的中间件
* 主要操作：
    1. 从DB中载入所有的Session信息。
    2. 创建SessionId在DB中和Cache(这里使用sync.Map)。
	3. 验证session是否过期，删除expired session在DB和Cache中

##总结：stream_server
1. 添加限流的结构，包含bucket：chan int和最大token数
2. 在middleWare中加入限流的结构，在ServeHTTP中加入获取token和释放token的操作

##总结：scheduler ：完成一些定时触发或延时触发的异步任务
* 流程：api的删除接口将要删除的vid写入video_del_record表中，达到时间间隔之后ticker触发，执行删除操作
* 添加task_runner包，这是一个可以独立运行的包，不是纯的方法调用的包
    1. 创建worker结构：包含要创建的runner结构和*time.Ticker
    2. 创建runner结构：注意要包含fn
