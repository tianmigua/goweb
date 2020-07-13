package web

import (
	"errors"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func UploadHandler(c *Context) error {
	w := c.Response()
	w.Header().Set("Content-Type", "text/html")
	req := c.Request()
	// 接收图片
	uploadFile, handle, err := req.FormFile("file")
	if err != nil {
		panic(err)
	}

	// 检查图片后缀
	ext := strings.ToLower(path.Ext(handle.Filename))
	if ext != ".jpg" && ext != ".png" {
		panic(errors.New("只支持jpg/png图片上传"))
		//defer os.Exit(2)
	}

	// 保存图片
	os.Mkdir("./uploaded/", 0777)
	timenow := time.Now().Unix()
	uploadedurl := "./uploaded/" + strconv.FormatInt(timenow, 10) + ".jpg"
	saveFile, err := os.OpenFile(uploadedurl, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	io.Copy(saveFile, uploadFile)

	defer uploadFile.Close()
	defer saveFile.Close()

	uploadedurl2 := uploadedurl[1:len(uploadedurl)]
	// 上传图片成功
	uploaddata := struct {
		Path1 string `json:"path1"`
		Path2 string `json:"path2"`
	}{uploadedurl2, uploadedurl2}
	//fmt.Println(uploaddata)
	return c.JSON(200, uploaddata)
}
