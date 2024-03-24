package bookcontroller

import "github.com/gin-gonic/gin"

func GetAllBook(ctx *gin.Context)  {
	ctx.JSON(200, gin.H{"Hello": "book"})
}
