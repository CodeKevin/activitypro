package apis

import (
	. "activitypro/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func AddActivityApi(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	uid, err := strconv.Atoi(c.PostForm("uid"))
	createat := int(time.Now().Unix())
	act := VxActivity{Title: title, Content: content, Uid: uid, CreatedAt: createat}
	ra, err := act.AddActivity()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("Insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
