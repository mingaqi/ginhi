package handle

import (
	"ginhi/service/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

// API parameter
// @Summary 保存用户
// @Description 保存用户
// @Tags 保存用户
// @Produce json
// @Param name path string true "name"
// @Success 200 {object} interface{} {"code":200,"data":result,"msg":"success"}
// @Success 400 {object} interface{} {"code":400,"data":,"msg":"failed"}
// @Router /v1/user/{name} [get]
func UserSave(c *gin.Context) {
	logrus.Infoln("请求接口:", c.Request.RequestURI)
	userName := c.Param("name")
	c.String(http.StatusOK, "the ["+userName+"] user get")
}

// path parameter
func UserSaveByQuery(context *gin.Context) {
	// 使用默认值
	userName := context.DefaultQuery("name", "Domi")
	age := context.Query("age")
	context.JSON(http.StatusOK, gin.H{"message": "用户:" + userName + ",年龄:" + age + "已经保存"})
}

// post json  query form data json都可以使用shouldBind接收
func UserRegister(c *gin.Context) {
	user := model.UserModel{}
	if err := c.ShouldBind(&user); err != nil {
		log.Print("err-->", err)
		return
	}
	println("email", user.Email, "password", user.Password, "password again", user.PasswordAgain)
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "success",
	})
}

func JsonTest(c *gin.Context) {
	var params map[string]string
	if err := c.ShouldBind(&params); err != nil {
		log.Print("err-->", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "success",
		"data": params,
	})
}

func Redirect(c *gin.Context) {
	c.Redirect(http.StatusFound, "http://www.baidu.com")
}
