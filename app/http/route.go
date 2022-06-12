package http

import (
	"project/app/http/module/demo"
	"project/framework/gin"
)

func Routes(r *gin.Engine) {
	r.Static("/dist/", "./dist/")
	demo.Register(r)
}
