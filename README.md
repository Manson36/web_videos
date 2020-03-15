#流媒体点播网站
##为什么选择视频网站
* Go是一门网络编程语言
* 视频网站包含Go在实战项目中的绝大部分要点
* 优良的native http库一级模板引擎（无需任何第三方框架）

##前后端解耦
###1.什么是前后端解耦
* 前后端解耦是时下最流行的web网站架构
分为前端、大前端、后端、大后端
* 前端页面和服务通过普通的web引擎渲染
* 后端数据通过渲染后的页面脚本调用后处理和呈现
###2.前后端解耦的优势
* 解放生产力，提高合作效率
>解耦之后前端主要负责页面的渲染以及数据的如何获取和它的如何展现，而后端主要是来设计它的API，设计它的整个架构，让它的架构更能实现高可用，可扩展，适应业务变化，这样前端和后端更专注于自己本质的工作

*松耦合的架构更灵活，部署更方便，更符合微服务的设局特征
*性能的提升，可靠性的提升
###3.缺点
* 工作量大
* 带来的团队成本和学习成本
* 系统的复杂度加大

##后端服务
###API
1. api是什么？
api是我们后端形成的一个大的服务端的接口，我们一定要让这个接口通用、简单。我们在这里使用restful风格的api
    * REST(Representational Status Transfer) API
    * REST是一种设计风格，不是任何标准的架构
    * 当今的RESTful API通常使用HTTP作为通信协议，JSON作为数据格式

2. api特点：
    * 统一接口(Uniform interface)
    * 无状态(Stateless):
     *无状态就是说无论我什么时候想要这个东西他都是没有变化的*   
    * 可缓存(Cacheable):
     *为了减少后端服务的压力，我们一般会把常用的，或者读远大于写的东西放在缓存中*
    * 分层(Layered System)
    * CS模式(Client-server Atchitecture)

3. API的设计原则
    * 以URL(统一资源定位符)风格设计API
    * 通过不同的METHOD(GET,POST,DELETE,PUT)来区分对资源的CRUD
    * 返回码(Status Code)符合HTTP资源描述的规定

##API设计
###用户
* 创建(注册)用户：URL:/user Method:POST, SC:201,400,500
*201表示创建成功*
* 用户登录：URL:/user/:username Method:POST, SC:200, 400, 500
* 获取用户基本信息：URL:/user/:username Method:GET, SC:200,400,401,403,500
*401表示并没有验证，403表示虽然通过验证，但是没有操作某项资源的权限*
* 用户注销：URL:/user/:username Method:DELETE, SC:204,400,401,403,500

###用户资源(videos)
* List All videos:URL:/user/:username/videos Method:GET,SC:200,400,500
*这里可能还要加分页，因为在这里一页可能无法完全显示*
* Get One video:URL:/user/:username/videos/:vid-id Method:GET, SC:200,400,500
* Delete video:URL:/user/:username/videos/:vid-id Method:DELETE, SC:200,400,401,403,500

###评论
* Show comments: URL:/videos/:vid-id/comments Method:GET, SC:200,400,500
* Post One comments: URL:/videos/:vid-id/comments Method:POST, SC:201,400,500
* Delete One comments: URL:/videos/:vid-id/comments/:comment-id Method:DELETE, SC:204,400,401,403,500

##数据库设计
###用户
Table:users
`id UNSIGNED INT, PRIMARY KEY, AUTO_INCREMENT;
login_name VARCHAR(64), UNIQUE KEY;
pwd TEXT`

###视频资源
TABLE:video_info
`id VARVHAR(64),PRIMARY KEY,NOT NULL;
author_id UNSIGENED INT;
name TEXT;
display_ctime TEXT;
create_tiem DATETIME
`

###评论
TABLE:comments
`id VARCHAR(64),PRIMARY KEY, NOT NULL;
video_id VARCHAR(64);
author_id UNSIGNED INT;
content TEXT;
time DATETIME
`

###sessions
TABLE:sessions
`session_id TINYTEXT, PRIMARY KEY, NOT NULL;
TTL TINYTEXT;
login_name VARCHAR(64)`
*通过session_id在server端检测id是否是有效的*
*重要点:如何完成session过期和session校验*

1. 什么事session？  
    顾名思义，session就是会话，是一种中间态。会话的作用是什么？  
    当我们在客户端与服务端交互，使用的是http，严格点叫restful API，它的特点是状态是不会保持的(Stateless)，我们在服务器用户
    的一个状态我们必须有一个东西来保存它，以完成一些功能，这时就需要session。
2. session和cookie的区别：  
    首先是session，session是一种机制，在服务端为用户保存它相应状态的一种机制；而cookie是在客户端为用户保存的一种机制   
    当我们使用session的时候，需要一个session的id，客户端为了方便这个session的访问的时候，会把sessionID放到Cookie中，两者发生关系
    
