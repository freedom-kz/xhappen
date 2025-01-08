package service

import (
	"io"

	"github.com/go-kratos/kratos/v2/transport/http"
)

// 多个文件上传
func UploadFile(ctx http.Context) error {
	req := ctx.Request()

	//1.用户验证
	//2.文件处理
	//3.关系数据保存
	_, err := GetUserID(ctx)
	if err != nil {
		return err
	}

	formdata := req.MultipartForm
	files := formdata.File["file"]

	for _, v := range files {
		// fileName := v.Filename
		file, err := v.Open()
		if err != nil {
			return err

		}
		defer file.Close()
		_, err = io.ReadAll(file)
		if err != nil {
			return err
		}
		//TODO,minio op
	}
	return ctx.String(200, "File Uploaded successfully")
}
