package test

import (
	"github.com/foxiswho/shop-go/router/base"
	"os"
	"io"
	"net/http"
	"fmt"
	"github.com/foxiswho/shop-go/conf"
)

type Upload struct {
}

func NewUpload() *Upload {
	return new(Upload)
}
func (x *Upload) UploadIndex(c *base.BaseContext) error {
	c.Set("tmpl", "example/test/upload")
	c.Set("data", map[string]interface{}{
		"title":         "上传",
	})
	return nil
}

func UploadPostIndex(c *base.BaseContext) error {

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	root_path:="."+conf.Conf.Upload.RootPath
	err = os.MkdirAll(root_path, os.ModePerm)
	fmt.Print("Create Directory=========",root_path)
	if err != nil {
		fmt.Printf("Create Directory ERROR %s", err)
	} else {
		fmt.Print("Create Directory OK! ",root_path)
	}
	// Destination
	dst, err := os.Create(root_path+file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully </p>", root_path+file.Filename))
}

func UploadMorePostIndex(c *base.BaseContext) error {
	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]
	root_path:="."+conf.Conf.Upload.RootPath
	err = os.MkdirAll(root_path, os.ModePerm)
	fmt.Print("Create Directory=========",root_path)
	if err != nil {
		fmt.Printf("Create Directory ERROR %s", err)
	} else {
		fmt.Print("Create Directory OK! ",root_path)
	}
	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(root_path+file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>Uploaded successfully %d files .</p>", len(files)))
}