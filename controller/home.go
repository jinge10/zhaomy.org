package controller

import "github.com/gin-gonic/gin"

func (h *page) Home(ctx *gin.Context) {
	res := make(map[string]interface{})
	res["code"] = 2000
	res["msg"] = "success"
	res["data"] = "data"
	h.Render(ctx, res)
}
