package handler

import (
	"fmt"
	"net/http"

	gg "github.com/tbxark/g4vercel"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	server := gg.New()
	server.Use(gg.Recovery(func(err interface{}, c *gg.Context) {
		if httpError, ok := err.(gg.HttpError); ok {
			c.JSON(httpError.Status, gg.H{
				"message": httpError.Error(),
			})
		} else {
			message := fmt.Sprintf("%s", err)
			c.JSON(500, gg.H{
				"message": message,
			})
		}
	}))
	server.GET("/", func(context *gg.Context) {
		context.JSON(200, gg.H{
			"message": "OK",
		})
	})
	server.POST("/hello", func(context *gg.Context) {
		name := context.Query("name")
		if name == "" {
			context.JSON(400, gg.H{
				"message": "name not found",
			})
		} else {
			context.JSON(200, gg.H{
				"data": fmt.Sprintf("Hello %s!", name),
			})
		}
	})
	server.GET("/user/:id", func(context *gg.Context) {
		context.JSON(400, gg.H{
			"data": gg.H{
				"id": context.Param("id"),
			},
		})
	})
	server.Handle(w, r)
}
