package api

import (
	"SilicomAPPv0.3/global"
	"SilicomAPPv0.3/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var permission = 0

// WebFileUpload 图片文件上传
func WebFileUpload(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		response.Failed("上传图片出错", c)
	}
	image := global.Config.Upload
	err = c.SaveUploadedFile(file, image.Path+file.Filename)
	if err != nil {
		return
	}
	imageURL := image.AccessUrl + file.Filename
	response.Success("上传图片成功", imageURL, c)
}

// 测试用
func WebGetPicture(c *gin.Context) {
	username := fmt.Sprint(c.Keys["username"])
	fmt.Println("请求的用户：", c.Keys["username"])
	fmt.Println("用户权限：", u.GetUserPermission(username))
	if u.GetUserPermission(username) < permission {
		response.Failed("权限不足", c)
		return
	}
	c.Redirect(http.StatusFound, "/image/test.jpg")
}
