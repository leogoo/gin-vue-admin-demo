package routers

import (
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func SetMediaRouter(r *gin.Engine) {
	r.POST("/api/upload", func(ctx *gin.Context) {
		file, err := ctx.FormFile("file")
		if err != nil {
			ctx.JSON(200, gin.H{
				"data":    1,
				"message": err,
			})
		}

		//获取文件的后缀名
		extStr := path.Ext(file.Filename)
		key := uuid.New().String()
		fileName := key + extStr

		//保存上传文件
		os.Mkdir("upload", os.ModePerm)
		filePath := filepath.Join("upload", "/", fileName)
		ctx.SaveUploadedFile(file, filePath)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "success",
			"data": gin.H{
				"uuid": fileName,
			},
		})
	})

	r.GET("/api/downloadByBlob", func(ctx *gin.Context) {
		fileName := ctx.Query("fileName")

		//打开文件
		_, err := os.Open("upload/" + fileName)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":    -3,
				"message": err,
			})
		}

		ctx.Header("Content-Type", "application/octet-stream")
		ctx.Header("Content-Disposition", "attachment; filename="+fileName)
		ctx.Header("Content-Transfer-Encoding", "binary")
		ctx.File("upload/" + fileName)
	})
}
