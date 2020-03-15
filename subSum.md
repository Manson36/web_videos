#流媒体学习总结
##章节小总结
1. 数据库实现User

    1. listenANDServer流程：  
    listen -> RegisterHandlers -> handlers(CreateUser,Login)自动以goroutine 的方式启动  
    handler -> validation(校验){1.request是不是合法，2.user是不是合法用户} -> business logic(逻辑处理) -> response  
    validation：{1.data model(数据结构) 2.error handling}
    2. 数据库对User表的CRUD操作和Test  
    添加api_test.go
