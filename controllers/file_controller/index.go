package filecontroller

import (
	"gin-gonic-gorm/constanta"
	"gin-gonic-gorm/utils"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func HandleUploadFile(ctx *gin.Context)  {
	fileHeader, _ := ctx.FormFile("file")
	if fileHeader == nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "file is required",
		})
		return
	}

	fileType := []string{".png"}
	isFileValidated := utils.FileValidationByExtention(fileHeader, fileType)
	if !isFileValidated {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "file not allowed",
		})
		return
	}

	// fileType := []string{"image/jpg"}
	// isFileValidated := utils.FileValidation(fileHeader, fileType)
	// if !isFileValidated {
	// 	ctx.AbortWithStatusJSON(400, gin.H{
	// 		"message": "file not allowed",
	// 	})
	// 	return
	// }


	extentionFile := filepath.Ext(fileHeader.Filename)
	filename := utils.RandomFileName(extentionFile)

	isSaved := utils.SaveFile(ctx, fileHeader, filename)
	if !isSaved {
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "File uploaded",
	})
}

func HandleRemoveFile(ctx *gin.Context) {
	filename := ctx.Param("filename")
	if filename == "" {
		ctx.JSON(400, gin.H{
			"message": "filename is required",
		})
		return
	}

	err := utils.RemoveFile(constanta.DIR_FILE + filename)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "file deleted",
	})
}
