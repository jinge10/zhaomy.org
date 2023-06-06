package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"zhaomy.org/router"
)

type page struct {
}

func isAjaxRequest(r *http.Request) bool {
	return strings.EqualFold("XMLHttpRequest", r.Header.Get("X-Requested-With"))
}
func (p *page) Render(c *gin.Context, data interface{}) {
	fmt.Println(c.Request.Method)
	if isAjaxRequest(c.Request) {
		c.JSON(200, data)
		return
	}
	config := make(map[string]interface{})
	config["uid"] = 23455
	c.HTML(200, "index.html", config)
}

var Page page

func init() {
	r := router.GetRouter()
	r.Group("/")
	r.GET("/", Page.Home)
}
