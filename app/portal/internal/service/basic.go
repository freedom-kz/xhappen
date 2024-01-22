package service

import (
	"io"
	"os"

	"github.com/go-kratos/kratos/v2/transport/http"
)

func UploadFile(ctx http.Context) error {
	req := ctx.Request()

	// formdata := req.MultipartForm
	// files := formdata.File["file"]

	fileName := req.FormValue("name")
	file, header, err := req.FormFile("file")
	if err != nil {
		return err
	}
	defer file.Close()

	f, err := os.OpenFile(header.Filename, os.O_WRONLY|os.O_CREATE, 0o666)
	if err != nil {
		return err
	}
	defer f.Close()
	_, _ = io.Copy(f, file)

	return ctx.String(200, "File "+fileName+" Uploaded successfully")
}
