package router

import (
	"bizd/metion"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"io/ioutil"
	"log"
	"net/http"
)

//设置跨域头的中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
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

//设置路由方法
func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "test",
	})
}

//login获取token
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
	c.JSON(http.StatusOK, `{"code":0,"data":{"id":2,"username":"demo1","nickname":"demo1","avatar":"http://admin.gumingchen.icu/file/d6cddc7f-9d67-4366-bc8e-ecad19e1bf76.webp","mobile":"13777777777","email":"1240235512@qq.com","sex":1,"status":1,"supervisor":0,"roles":[{"id":1,"name":"Demo","permission":5,"custom":""}],"token":"28263720307e1a0fcfbdbd021b26f062","department_id":2,"enterprise_id":1,"created_at":"2022-07-15 14:18:02","updated_at":"2022-11-29 13:29:33","department_name":"前端开发部门","department_permission":4,"department_custom":"1"},"message":"成功！","status":"success"}`)
}

// SetupRouter 配置路由信息
func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.Use(Cors())

	//defaute
	r.GET("/test", test)

	//登录获取token
	r.POST("/slipper/admin/login", login)

	//获取路由信息
	r.GET("/slipper/admin/administrator/self/info", menuInfo)

	return r
}
