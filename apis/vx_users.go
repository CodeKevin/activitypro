package apis

import (
	"activitypro/models"
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUserApi(c *gin.Context) {
	println(c.PostForm("name"))
	//c.JSON(http.StatusOK, gin.H{"msg": "you are insert success"})
	var form models.VxUsers
	if err := c.Bind(&form); err == nil {
		_, err := form.AddUser()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"msg": err.Error})
		} else {
			c.JSON(http.StatusOK, gin.H{"msg": "you are insert success"})
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{"msg": "失败"})
	}
}
func AddUserApiTest(c *gin.Context) {
	var form models.User
	if err := c.Bind(&form); err == nil {
		if form.Name == "Kevin" {
			c.JSON(http.StatusOK, gin.H{"msg": "hello"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"msg": "失败"})
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{"msg": "失败"})
	}
}
