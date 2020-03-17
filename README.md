#流媒体点播网站
##为什么选择视频网站
* Go是一门网络编程语言
* 视频网站包含Go在实战项目中的绝大部分要点
* 优良的native http库一级模板引擎（无需任何第三方框架）

##前后端解耦
####1.什么是前后端解耦
* 前后端解耦是时下最流行的web网站架构
分为前端、大前端、后端、大后端
* 前端页面和服务通过普通的web引擎渲染
* 后端数据通过渲染后的页面脚本调用后处理和呈现
####2.前后端解耦的优势
* 解放生产力，提高合作效率
* 松耦合的架构更灵活，部署更方便，更符合微服务的设局特征
* 性能的提升，可靠性的提升
>解耦之后前端主要负责页面的渲染以及数据的如何获取和它的如何展现，而后端主要是来设计它的API，设计它的整个架构，
让它的架构更能实现高可用，可扩展，适应业务变化，这样前端和后端更专注于自己本质的工作

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

###API设计
####用户
* 创建(注册)用户：URL:/user Method:POST, SC:201,400,500
*201表示创建成功*
* 用户登录：URL:/user/:username Method:POST, SC:200, 400, 500
* 获取用户基本信息：URL:/user/:username Method:GET, SC:200,400,401,403,500
*401表示并没有验证，403表示虽然通过验证，但是没有操作某项资源的权限*
* 用户注销：URL:/user/:username Method:DELETE, SC:204,400,401,403,500

####用户资源(videos)
* List All videos:URL:/user/:username/videos Method:GET,SC:200,400,500
*这里可能还要加分页，因为在这里一页可能无法完全显示*
* Get One video:URL:/user/:username/videos/:vid-id Method:GET, SC:200,400,500
* Delete video:URL:/user/:username/videos/:vid-id Method:DELETE, SC:200,400,401,403,500

####评论
* Show comments: URL:/videos/:vid-id/comments Method:GET, SC:200,400,500
* Post One comments: URL:/videos/:vid-id/comments Method:POST, SC:201,400,500
* Delete One comments: URL:/videos/:vid-id/comments/:comment-id Method:DELETE, SC:204,400,401,403,500

###数据库设计
####用户
Table:users  
`id UNSIGNED INT, PRIMARY KEY, AUTO_INCREMENT;   
login_name VARCHAR(64), UNIQUE KEY;  
pwd TEXT`

####视频资源
TABLE:video_info  
`id VARVHAR(64),PRIMARY KEY,NOT NULL;  
author_id UNSIGENED INT;  
name TEXT;  
display_ctime TEXT;  
create_tiem DATETIME
`

####评论
TABLE:comments  
`id VARCHAR(64),PRIMARY KEY, NOT NULL;  
video_id VARCHAR(64);  
author_id UNSIGNED INT;  
content TEXT;  
time DATETIME
`

####sessions
TABLE:sessions  
`session_id TINYTEXT, PRIMARY KEY, NOT NULL;  
TTL TINYTEXT;  
login_name VARCHAR(64)`  
*通过session_id在server端检测id是否是有效的*
*重要点:如何完成session过期和session校验*

###main.go
####session
1. 什么事session？  
    顾名思义，session就是会话，是一种中间态。会话的作用是什么？  
    当我们在客户端与服务端交互，使用的是http，严格点叫restful API，它的特点是状态是不会保持的(Stateless)，我们在服务器用户
    的一个状态我们必须有一个东西来保存它，以完成一些功能，这时就需要session。
2. session和cookie的区别：  
    首先是session，session是一种机制，在服务端为用户保存它相应状态的一种机制；而cookie是在客户端为用户保存的一种机制   
    当我们使用session的时候，需要一个session的id，客户端为了方便这个session的访问的时候，会把sessionID放到Cookie中，两者发生关系
    
3. session的流程：  
   1. 当用户进入页面的时候，需要Signin或Register，api对这些信息进行处理，处理完成之后给用户返回一个值(此时用户login)，
   同时返给客户端一个sessionID，客户端会把sessionId记录到它的Cookie中。然后api在处理Signin或Register的过程中，
   会往Cache中写一份产生的SessionID同时也会往DB中写入一份。
   2. 当用户再一次进入这个页面的时候，不用再用户名和密码注册了，直接使用SessionID。api会使用他的用户名和SessionID去Cache
   中找他的状态，他的SessionID是否存在，或者说session存在，若存在会把sessionID返回给他，这个用户就是Login的，
   这样在每个页面用户不需要再登录了，只需要SessionID即可。当然，如果在Cache中找不到sessionID，我们还回去DB中再查找(fetch)
   一次，如果Fetch到，我们给他返回SessionID，他同样用户状态就是Login；如果都没有就需要用户重新Login。
   3. 系统在初始化也就是重启的时候，Cache会自动去DB中Get所有的SessionID
   4. 当我在更新用户的Session的时候，不仅要忘Cache中写入，还要往DB中写，所以需要些两次。为什么要这么麻烦的写两次，原因是：
   首先DB在网页的访问量和变化量比较大的时候，它的读写压力是非常大的，而且DB的操作对io的消耗是非常大的，所以我们要尽可能减少
   DB的操作，所以这时候就需要使用Cache。Cache机制能够保证它在多读少写的情况下以最快的速度返回用户想要的结果，这就是Cache的作用。

#### MiddleWare
1. 为什么使用MiddleWare
首先是Main()函数，这时api的入口；第二步：在这么我们直接把Router放入，光这样做显然是不够的。我们需要在Router之前做一些公用的东西，
比如说校验，鉴权，流控和一些其他处理，这些统一称为HTTP MiddleWare
2. 此时的流程：
main -> middleWare -> defs(message, err) -> handlers -> dbops - response
3. MiddleWare的实现
    * 在main中直接就是RegisterHandlers，它返回的是*httprouter.Router，实际上它是实现了http中的一个接口 http.Handler
    * http.Handler接口 只有一个方法ServeHTTP(http.ResponseWriter, *http.Request)，duck type，只要实现了ServeHTTP 即可实现http.Handler
#### Handler任务：   
   * 从request中获取body，并将其反序列化到创建好的 Handler需要的ReqBody结构
   * 对传入的参数进行验证
   * 执行数据库操作和session操作
   * 将执行的结果返回到response中

##Streaming
###介绍
1. 什么是Streaming？   
我们平时在优酷、爱奇艺等看的视频就是streaming的过程
2. Streaming特点：
 * 静态视频，非RTMP
 > REMP是用在直播上的，也就是我们的客户端不断的有输入(input), 而且在另外的客户端有(output)  
 * 独立的服务，可独立的部署
 * 统一的api格式
###Stream Server
> 主要实现两个功能：1.Streaming(视频播放) 2.Upload files(视频上传)
####流控
1. 使用流控的原因：  
在Stream Server中，没有数据库的操作和session的验证，但是在Stream Server中有一个很重要的东西，和普通api不太一样的。
因为在stream server中我么会做两件事情：Streaming和Uploading files。这两件事情都需要保持长连接，跟之前的api的http的短连接是不太一样的。
当用户发送一个request，我可能会不断的向client端输出数据流，持续的时间会比较长。所以我们在多路长连接同时保持的时候，
就会出现一个问题，如果他不断的发起连接，我们不断地打开视频，总有一天会把我们的server Crash 掉。所以，为了避免这种情况，
我们需要加一个东西 Limit(流控)，我们这一个流控和普通的流控也有些不一样，我们只控制它的 connection 部分
2. 什么是流控？  
在我们整个网站过程中，在网站server online的时候，会有用户或者攻击者 有意或无意的向网站(server)发起请求，当请求达到一定数量的时候
会导致网站server的链接数不够，链接数不够还事小，但是如果他不仅将你的链接数消耗完，并且把你的带宽占完之后，那么你的server就处于不可用的状态，
还有一种更可怕的情况是，当你的 消耗完之后，你的系统可能会 crash，为了保护我们的server、保护系统，我们需要做流控。
流控用英语来说叫 run limiter,一般都用一种算法 bucket-token算法。
3. bucket-token算法：  
有一个bucket，它包含 n个token，request 拿到token 之后才算是真正的可以进入服务的request；当request拿到它想要的结果，返回response之后，
那么request的生命周期在我们的server中也已经结束了，此时将token还回去，放回bucket，以便给下一个request使用。以此来限制访问我们server的数量
4. bucket-token的使用：  
通常大家会想bucket-token直接给数组赋值就可以了，但实际上并没有这么简单。当我们使用数组或者一段普通的变量的话，由于我们的Handler
(在go中，每一个Handler都是新起了一个goroutine，而goroutine它们之间是并发的)，在并发的情况下，很可能会有不同handler访问同一个变量的情况，
如果我们的bucket是一个普通的内存中的数据块的话，同一时间访问它的话肯定会出现问题，我们为了避免竞争，肯定要给它加锁，来保证它的同步；
加了锁之后，它的性能势必会下降，影响性能是小，实际上这并不是go 设计出来用来处理并发问题的理想做法，go 常使用channel。(shared channel
instead of shared memory.)通过共享通道来同步线程之间的信息，而不是使用共享内存。
5. buffer channel和no_buffer channel 的区别：  
no_buffer channel相当于是一个同步channel，当这个channel写进数据的时候，如果没有别的goroutine来读这个channel，这个channel永远都是阻塞的，
就不会有别的东西往里写，对程序来说不是特别好用。在通常的情况下，我们会使用一些buffer channel来做消息之间的同步，有了buffer就可以保证
在一定的缓冲区间之内做一些消息同步的事情。

####Streaming Handler
> 实现是一个难点  

###Scheduler
####Scheduler 内容介绍
1. 什么是Scheduler？  
顾名思义，Scheduler就是调度器，调度一些 我们通过普通的rest api 没有办法马上给他结果的任务，这些任务都会发送到Scheduler中，
 Scheduler会根据它的时间的interval来定时给他触发，或者是延时触发。Scheduler就是用来完成这种异步任务的。
2. Scheduler包含什么？
    * restful 的http server *用于接受任务*
    * Timer
    * 生产者/消费者模式下的task runner
    
#### 项目流程
项目的整个流程：
* user -> api service -> delete video
* api service ->scheduler -> write video deletion record
* timer -> runner -> read wvdr -> exec ->delete video from file

1. api service: 创建路由和handler函数，将要删除的video_id 写入删除的表中
2. scheduler：
   创建Worker的结构，包含timer和runner
3. runner:
   * 首先我们要有一个runner的对象，
   * 在runner中我们会跑一个常驻的任务(startDispatcher),整个任务会长时间的去等待这的一个runner的channel，这个channel分为两部分(
      control channel/data channel),control channel是用来dispatcher和executor来相互交换信息，来提醒对方；data channel 是真正的数据部分
   * runner_test:
    * 执行runner.StartAll的时候，没有直接调用startAll，而是使用go routine，为什么？  
	在StartAll中包含startDispatch函数，它内部有for的死循环(死循环中加forloop也可以破解)，如果我们不在后台以go routine的形式启动，
	就会一直blocking(不断地写入数据，读取数据)，就不会执行到下面的time.sleep   
**在go中我们有时会需要block整个进程，不让他运行结束，那么就有两种方式：一种是for的死循环，另一个是声明一个没有buffer的channel，不给里面写东西，直接去读**
