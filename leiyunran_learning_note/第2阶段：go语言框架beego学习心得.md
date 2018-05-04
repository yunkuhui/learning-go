#beego框架学习心得

* 读取conf文件配置

  框架默认配置文件为：app.conf，可通过 `beego.AppConfig.String("key")` 来获取配置文件中相应的value值。可以配置json和xml文件，读取文件路径使用相对路径。

      // 在app.conf文件中配置数据库连接
      test = root:root@tcp(127.0.0.1:3306)/test?charset=utf8&loc=Asia%2FShanghai
      
      // 在代码中获取配置信息
      var test string
      test = beego.AppConfig.String("test")
      
      // 在代码中读取json文件
      var data []byte
      data, _ = ioutil.ReadFile("conf/config.json")
      
  
* mysql数据库

  可通过工具（doc命令行：`bee api 项目名称 -conn=root:root@tcp(127.0.0.1:3306)/test?charset=utf8`）生成mysql数据库对应的model。  
  初始化连接数据库一般放在init函数中。go文件中，如果有init函数，将会在文件加载时自动调用。beego支持多种查询mysql的方式。比如原生sql查询，同时可以使用占位符的方式。
  
  配置mysql数据库连接参数时，最好设置时区 `loc=Asia%2FShanghai`。beego默认的时间是国际UTC时间。结构体中对应的时间属性，可以使用 `time.Time`

      // mysql连接: main.go文件中的init函数
      func init() {
        orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("test"))    
      }
      
  
* mongodb数据库
  
  beego官方不支持mongodb数据库，需要使用第3方类库。如：`gopkg.in/mgo.v2` 或 `labix.org/v2/mgo`。
  
* 输出json

  可以定义一个 `data := map[string]interface{}{"resultCode":"00"}`，然后通过上下文模块变量使用 `output.JSON(data, true, false)` 方式响应输出json

      // 控制层的参数类型*context.Context为上下文模块
      func TestController(context *context.Context) {
        data := map[string]interface{}{"resultCode":"00"}
        // 响应输出json，参数一data为响应数据，参数二为是否格式化显示json，参数三为是否转为utf-8 encoding值
        context.output.JSON(data, true, false)
      }
      
  
* 路由配置

  路由一般在routers包中的router.go文件中的 `init`函数中配置。支持固定路由，正则路由和RESTful路由。

      // 固定路由配置，router.go文件的init函数中配置自定义路由
      func init() {
        beego.Get("/test", controllers.TestController)
      }
      
* 请求参数

  请求参数，一般在上下文结构体 `*context.Context`的 `input`对象里。可以通过调用响应的方法获取请求参数。如果是使用axios等发送json数据，可以在通过 `input`对象的 `RequestBody`获取byte数据。

      // 获取请求参数
      func TestController(context *context.Context) {
        // 获取表单数据
        var name string
        name = context.Input.Query("name")
        // 获取请求中body数据
        var data []byte
        data = context.Input.RequestBody
      }