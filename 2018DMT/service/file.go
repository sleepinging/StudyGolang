package service

import (
	"net/http"
	"../models"
	"../tools"
	"../global"
	"os"
	"fmt"
	"io"
	"strings"
)

//上传文件之后获取URL
func GetFileUrl(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(2048)
	fs, ok := r.MultipartForm.File["Filename"]
	if !ok {
		models.SendRetJson(0, "缺少Filename参数", "手动滑稽", w)
		return
	}
	file, err := fs[0].Open()
	if err != nil {
		fmt.Println("读取文件", fs[0].Filename, "失败", err)
		models.SendRetJson(0, "上传失败", "服务器出错", w)
		return
	}
	savename := global.FileDirPath + tools.GenFileName(fs[0].Filename)
	f, err := os.OpenFile(savename, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		fmt.Println("打开文件", savename, "失败", err)
		models.SendRetJson(0, "上传失败", "服务器出错", w)
		return
	}
	io.Copy(f, file)
	fmt.Println(tools.FmtTime(), r.RemoteAddr, "上传文件:", f.Name())
	furl := strings.TrimPrefix(savename, global.Config.Wwwroot)
	furl = strings.Replace(furl, `\`, `/`, -1)
	models.SendRetJson(0, "上传成功", `/`+furl, w)
	return
}
