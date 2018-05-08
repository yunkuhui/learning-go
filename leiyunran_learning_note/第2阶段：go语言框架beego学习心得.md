# beego 框架学习心得

* 读取 conf 文件配置

  框架默认配置文件为：app.conf，可通过 `beego.AppConfig.String("key")` 来获取配置文件中相应的 value 值。可以配置 json 和 xml 文件，读取文件路径使用相对路径。

      // 在 app.conf 文件中配置数据库连接
      test = root:root@tcp(127.0.0.1:3306)/test?charset=utf8&loc=Asia%2FShanghai
      
      // 在代码中获取配置信息
      var test string
      test = beego.AppConfig.String("test")
      
      // 在代码中读取 json 文件
      var data []byte
      data, _ = ioutil.ReadFile("conf/config.json")
      
  json 解析，可以使用 beego 官方提供的解析方式，转为结构体；也可以使用第 3 方类库动态解析 json。以下为使用 `github.com/bitly/go-simplejsongithub.com/bitly/go-simplejson` 类库动态解析 json

      // 动态解析 json；data 为 []byte 类型
      var json simplejson.Json
      json, _ = simplejson.NewJson(data)      
      arr, _ := json.Get("key1").Get("array").Array()
      i, _ := json.Get("key2").Get("int").Int()
      ms := json.Get("key3").Get("string").MustString()
      f := json.Get("key4").Get("float").MustFloat64()
  
* mysql 数据库

  可通过工具（doc 命令行：`bee api 项目名称 -conn=root:root@tcp(127.0.0.1:3306)/test?charset=utf8`）生成 mysql 数据库对应的 model。  
  初始化连接数据库一般放在 init 函数中。go 文件中，如果有 init 函数，将会在文件加载时自动调用。beego 支持多种查询 mysql 的方式。比如原生 sql 查询，同时可以使用占位符的方式。
  
  配置 mysql 数据库连接参数时，最好设置时区 `loc=Asia%2FShanghai`。beego 默认的时间是国际 UTC 时间。结构体中对应的时间属性，可以使用 `time.Time`

      // mysql连接: main.go 文件中的 init 函数
      func init() {
        orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("test"))    
      }
      
  
* mongodb 数据库
  
  beego 官方不支持 mongodb 数据库，需要使用第 3 方类库。如：`gopkg.in/mgo.v2` 或 `labix.org/v2/mgo`。
  
* 输出 json

  可以定义一个 `data := map[string]interface{}{"resultCode":"00"}`，然后通过上下文模块变量使用 `output.JSON(data, true, false)` 方式响应输出 json

      // 控制层的参数类型 *context.Context 为上下文模块
      func TestController(context *context.Context) {
        data := map[string]interface{}{"resultCode":"00"}
        // 响应输出 json，参数一 data 为响应数据，参数二为是否格式化显示 json，参数三为是否转为 utf-8 encoding 值
        context.output.JSON(data, true, false)
      }
      
  
* 路由配置

  路由一般在 routers 包中的 router.go 文件中的 `init` 函数中配置。支持固定路由，正则路由和 RESTful 路由。

      // 固定路由配置，router.go 文件的 init 函数中配置自定义路由
      func init() {
        beego.Get("/test", controllers.TestController)
      }
      
* 请求参数

  请求参数，一般在上下文结构体 `*context.Context` 的 `input` 对象里。可以通过调用响应的方法获取请求参数。如果是使用 axios 等发送 json 数据，可以在通过 `input` 对象的 `RequestBody` 获取 byte 数据。

      // 获取请求参数
      func TestController(context *context.Context) {
        // 获取表单数据
        var name string
        name = context.Input.Query("name")
        // 获取请求中 body 数据
        var data []byte
        data = context.Input.RequestBody
      }