>
>
>开新坑了，模仿着用Go写一个HTTP Server



----

## Version_0.1

- 封装Server接口

  ```go
  type Server interface {
     Route(pattern string, handlerFunc fhttp.HandlerFunc)
     Start(address string) error
  }
  ```

  - 封装路由

  - 封装启动服务

  - 使用结构体实现Server接口

    ```go
    type sdkHttpServer struct {
    	Name string
    }
    func (s *sdkHttpServer) Route(pattern string, handlerFunc http.HandlerFunc)
    func (s sdkHttpServer) Start(address string) error
    func NewServer(name string) Server
    ```

- 封装上下文Context

  ```go
  type Context struct {
  	W http.ResponseWriter
  	R *http.Request
  }
  func (c *Context) ReadJson(req interface{}) error
  func (c *Context) WriteJson(code int, res interface{}) error
  func (c *Context) OKJson(res interface{}) error
  func (c *Context) SystemErrorJson(res interface{}) error
  func (c *Context) BadRequestJson(res interface{}) error
  ```

  - 使用结构体而不是接口的原因：？？？
  - 对返回不同HTTP状态码的读取JSON的方法、写JSON方法进行封装

- 主函数测试

  ```go
  server1 := NewServer("server-1")
  server1.Route("/", home)
  server1.Route("/sign", SignUp)
  server1.Start(":8080")
  ```

  

