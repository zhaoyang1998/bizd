package router

import (
	"bizd/metion"
	"bizd/metion/service"
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"log"
	"net/http"
)

// 设置跨域头的中间件

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			// 接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			// 服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			// 允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			// 设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			// 允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()
		c.Next()
	}
}

// 设置路由方法
func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "test",
	})
}

// login获取token
func login(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(data))
	username := jsoniter.Get(data, "username")
	//passWord := jsoniter.Get(data, "passWord")

	token, err := metion.CreateToken(1, username)
	fmt.Println(err)
	msg := metion.ResponseMsg{
		Code: 0,
		Data: struct {
			Token           string `json:"token"`
			AdministratorID int    `json:"administrator_id"`
			ExpiredAt       string `json:"expired_at"`
			CreatedAt       string `json:"created_at"`
		}{
			token,
			2,
			"2022-12-05 19:10:33",
			"2022-12-05 17:10:33"},
		Message: "成功！",
		Status:  "success",
	}

	c.JSON(http.StatusOK, msg)

}

func menuInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "成功！",
		"status":  "success",
		"data":    `{"id":2,"username":"demo1","nickname":"demo1","avatar":"http://admin.gumingchen.icu/file/d6cddc7f-9d67-4366-bc8e-ecad19e1bf76.webp","mobile":"13777777777","email":"1240235512@qq.com","sex":1,"status":1,"supervisor":0,"roles":[{"id":1,"name":"Demo","permission":5,"custom":""}],"token":"28263720307e1a0fcfbdbd021b26f062","department_id":2,"enterprise_id":1,"created_at":"2022-07-15 14:18:02","updated_at":"2022-11-29 13:29:33","department_name":"前端开发部门","department_permission":4,"department_custom":"1"}`,
	})
}

func getUsers(context *gin.Context) {
	service.GetUsers(context)
}

func getUsersByType(context *gin.Context) {
	service.GetUserByType(context)
}

func addUser(context *gin.Context) {
	service.AddUser(context)
}
func updateUser(context *gin.Context) {
	service.UpdateUser(context)
}

func delUser(context *gin.Context) {
	service.DelUser(context)
}

func addClient(c *gin.Context) {
	service.AddClient(c)
}

func getClients(c *gin.Context) {
	service.GetClients(c)
}

func delClient(context *gin.Context) {
	service.DelClient(context)
}

func updateClient(context *gin.Context) {
	service.UpdateClient(context)
}

func delPointPosition(context *gin.Context) {
	service.DelPointPosition(context)
}

func updatePointPosition(context *gin.Context) {
	service.UpdatePointPosition(context)
}

func getPointPosition(context *gin.Context) {
	service.GetPointPosition(context)
}

func addPointPosition(context *gin.Context) {
	service.AddPointPosition(context)
}

func finishAssignment(context *gin.Context) {
	service.FinishAssignment(context)
}

func startAssignment(context *gin.Context) {
	service.StartAssignment(context)
}

func allocatingAssignment(context *gin.Context) {
	service.AllocatingAssignment(context)
}

// SetupRouter 配置路由信息
func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.Use(Cors())

	// default
	r.GET("/test", test)

	// 登录获取token
	r.POST("/admin/login", login)

	// 获取路由信息
	r.GET("/admin/administrator/self/info", menuInfo)

	// 用户相关接口
	r.POST("/user/getUsers", getUsers)
	r.POST("/user/getUsersByType", getUsersByType)
	r.POST("/user/addUser", addUser)
	r.POST("/user/updateUser", updateUser)
	r.POST("/user/delUser", delUser)

	// 客户相关接口
	r.POST("/client/addClient", addClient)
	r.POST("/client/getClients", getClients)
	r.POST("/client/updateClient", updateClient)
	r.POST("/client/delClient", delClient)

	// 单位相关接口
	r.POST("/pointPosition/addPointPosition", addPointPosition)
	r.POST("/pointPosition/getPointPosition", getPointPosition)
	r.POST("/pointPosition/updatePointPosition", updatePointPosition)
	r.POST("/pointPosition/delPointPosition", delPointPosition)
	r.POST("/pointPosition/startAssignment", startAssignment)
	r.POST("/pointPosition/finishAssignment", finishAssignment)
	r.POST("/pointPosition/allocatingAssignment", allocatingAssignment)
	return r
}
