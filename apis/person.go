package apis

import (
	. "activitypro/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}
func AddPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastname := c.Request.FormValue("last_name")

	p := Person{FirstName: firstName, LastName: lastname}
	ra, err := p.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("Insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
func GetAllPersonApi(c *gin.Context) {
	p := Person{}
	persons, err := p.GetPersons()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": persons,
	})
}
