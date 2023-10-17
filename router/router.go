package router

import (
	"bizd/metion"
	"bizd/metion/service"
	"github.com/gin-gonic/gin"
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

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token") // 访问令牌
		if metion.CheckToken(token) && token != "" {
			// 验证通过，会继续访问下一个中间件
			c.Next()
		} else {
			// 验证不通过，不再调用后续的函数处理
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    201,
				"message": "账号未登录或已过期",
			})
			// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
			return
		}
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
	service.Login(c)
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

func getUsersByKeyword(context *gin.Context) {
	service.GetUsersByKeyword(context)
}

func getUsersByType(context *gin.Context) {
	service.GetUserByType(context)
}
func getAllUsers(context *gin.Context) {
	service.GetAllUsers(context)
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

func getAllClients(context *gin.Context) {
	service.GetAllClients(context)
}

func getClientsByKeyword(context *gin.Context) {
	service.GetClientsByKeyword(context)
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

func GetPointPositionByKeyword(context *gin.Context) {
	service.GetPointPositionByKeyword(context)
}

func addPointPosition(context *gin.Context) {
	service.AddPointPosition(context)
}

func finishAssignment(context *gin.Context) {
	service.FinishAssignment(context)
}

func loginOut(context *gin.Context) {
	context.JSON(200, gin.H{
		"code":    0,
		"message": "成功！",
		"status":  "success",
	})
}

func startAssignment(context *gin.Context) {
	service.StartAssignment(context)
}

func allocatingAssignment(context *gin.Context) {
	service.AllocatingAssignment(context)
}
func cancelAssignment(context *gin.Context) {
	service.CancelAssignment(context)
}

func getHomePageData(context *gin.Context) {
	service.GetHomePageData(context)
}

func getMenuJson(context *gin.Context) {
	service.GetMenuJson(context)
}

// SetupRouter 配置路由信息
func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.Use(Cors())

	// default
	r.GET("/test", test)

	// 登录获取token
	r.POST("/admin/login", login)
	r.POST("/admin/logout", loginOut)

	// 获取路由信息
	r.GET("/admin/administrator/self/info", menuInfo)

	// 用户相关接口
	userApi := r.Group("/user")
	{
		userApi.Use(Authorize())
		userApi.POST("/getUsers", getUsers)
		userApi.POST("/getUsersByKeyword", getUsersByKeyword)
		userApi.POST("/getUsersByType", getUsersByType)
		userApi.POST("/addUser", addUser)
		userApi.POST("/updateUser", updateUser)
		userApi.POST("/delUser", delUser)
		userApi.GET("/getAllUsers", getAllUsers)
	}

	// 客户相关接口
	clientApi := r.Group("/client")
	{
		clientApi.Use(Authorize())
		clientApi.POST("/addClient", addClient)
		clientApi.POST("/getClients", getClients)
		clientApi.GET("/getAllClients", getAllClients)
		clientApi.POST("/getClientsByKeyword", getClientsByKeyword)
		clientApi.POST("/updateClient", updateClient)
		clientApi.POST("/delClient", delClient)
	}

	// 单位相关接口
	pointPositionApi := r.Group("/pointPosition")
	{
		pointPositionApi.Use(Authorize())
		pointPositionApi.POST("/addPointPosition", addPointPosition)
		pointPositionApi.POST("/getPointPosition", getPointPosition)
		pointPositionApi.POST("/GetPointPositionByKeyword", GetPointPositionByKeyword)
		pointPositionApi.POST("/updatePointPosition", updatePointPosition)
		pointPositionApi.POST("/delPointPosition", delPointPosition)
		pointPositionApi.POST("/startAssignment", startAssignment)
		pointPositionApi.GET("/finishAssignment/:pointPositionId", finishAssignment)
		pointPositionApi.POST("/allocatingAssignment", allocatingAssignment)
		pointPositionApi.POST("/cancelAssignment", cancelAssignment)
		pointPositionApi.POST("/exportExcel", exportExcel)
	}

	commonApi := r.Group("/common")
	{
		//commonApi.Use(Authorize())
		commonApi.GET("/getHomePageData/:clientId", getHomePageData)
		commonApi.GET("/getMenuJson", getMenuJson)
		commonApi.GET("getAllStatus", getAllStatus)
	}

	detailApi := r.Group("/details")
	{
		detailApi.GET("/getAllDetail/:id", getAllDetail)
		detailApi.POST("/saveDetail", saveDetail)
	}

	documentApi := r.Group("/document")
	{
		// 富文本展示页面
		//documentApi.LoadHTMLGlob("templates/*.html")
		documentApi.POST("/saveFile", saveDocumentFile)
		documentApi.POST("/getDoc", getDocument)
	}
	return r
}

func getAllStatus(context *gin.Context) {
	service.GetAllStatus(context)
}

func exportExcel(context *gin.Context) {
	service.ExportExcel(context)
}

func saveDetail(context *gin.Context) {
	service.SaveDetail(context)
}

func getAllDetail(context *gin.Context) {
	service.GetAllDetail(context)
}

func saveDocumentFile(context *gin.Context) {
	service.SaveDocumentFile(context)
}
func getDocument(context *gin.Context) {
	service.GetDoc(context)

}
