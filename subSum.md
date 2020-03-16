#流媒体学习总结
##章节小总结
1. 数据库实现User

    * listenANDServer流程：  
    listen -> RegisterHandlers -> handlers(CreateUser,Login)自动以goroutine 的方式启动  
    handler -> validation(校验){1.request是不是合法，2.user是不是合法用户} -> business logic(逻辑处理) -> response  
    validation：{1.data model(数据结构) 2.error handling}
    * 数据库对User表的CRUD操作和Test  
    添加api_test.go
2. Session实现以及完成数据库实现Vido_info和Comments
> session我们的主要操作是：
  1.在我们服务起来的时候，从DB中拉取Session，把所有存在的Session全部load到我们Api节点的Cache中
  2.当有新用户注册或老用户登录的时候，我们需要费当前的用户分配一个SessionID，里面包含Session的信息，这时候就需要一个产生Session的方法
  3.当我们在校验的时候，Session可能会过期，这时需要给用户返回一个过期或没有过期的状态，用来判断当前这个用户是否是合法、已经登录的用户  
  
   * session作为中间件记录用户的登录信息，减少对数据库的操作
   * session的实现包括 数据库增删查 和 缓存(这里使用Map)的增改查
