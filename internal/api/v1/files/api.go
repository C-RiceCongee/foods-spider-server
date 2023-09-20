package files

import (
	"foods-spider-server/com"
	"foods-spider-server/internal/api"
	"foods-spider-server/pkg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type IFileApi interface {
	api.IBase
	GetFilesList() func(ctx *gin.Context)
	SingleUploadFile() func(ctx *gin.Context)
}

type FileApi struct {
}

func (f *FileApi) SingleUploadFile() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 单文件
		file, err := ctx.FormFile("file")
		sceneType := ctx.PostForm("sceneType")
		if err != nil {
			zap.L().Error(err.Error())
			com.ResponseData(ctx, com.CodeFailed, err)
			return
		}
		if len(sceneType) == 0 {
			com.ResponseData(ctx, com.CodeFailed, "请填写文件场景sceneType")
			return
		}
		dir, err := pkg.MkAllPathDir(pkg.Config.StaticBaseUrl + pkg.Config.StaticFileDirPath)
		if err != nil {
			zap.L().Error(err.Error())
			com.ResponseData(ctx, com.CodeFailed, err)
			return
		}
		dst := dir + file.Filename
		// 上传文件至指定的完整文件路径
		err = ctx.SaveUploadedFile(file, dst)
		if err != nil {
			com.ResponseData(ctx, com.CodeFailed, err)
			return
		}
		com.ResponseData(ctx, com.CodeSuccess, gin.H{
			"path": "/" + dst,
		})
	}
}

func (f *FileApi) SetPrefix() string {
	return "/api/v1/files"
}
func (f *FileApi) GetFilesList() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "获取文件列表",
		})
	}
}
func NewFilesApi() IFileApi {
	return &FileApi{}
}
