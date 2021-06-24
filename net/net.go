/*
Encoding: utf-8
Author: legend
Time: 下午9:50
*/
package net

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"gitee.com/legendqi/utils-go/file"
	"gitee.com/legendqi/utils-go/logger"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path"
)

//下载文件
func DownloadFile(savePath string, url string) (bool, error) {
	fileName := path.Base(url)
	res, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()
	reader := bufio.NewReaderSize(res.Body, 32*1024)
	file, err := os.Create(savePath + fileName)
	if err != nil {
		return false, err
	}
	writer := bufio.NewWriter(file)
	written, _ := io.Copy(writer, reader)
	fmt.Println("Total length is", written)
	return true, nil
}

//上传文件
func UploadFile(filePath string, server string) (bool, error) {
	if !file.CheckFileIsExist(filePath) {
		logger.Error("file not exist")
		return false, errors.New("file not exist")
	}
	//获取文件信息
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		logger.Error("get file info error")
		return false, err
	}
	bodyBuffer := &bytes.Buffer{}
	writer := multipart.NewWriter(bodyBuffer)
	formFile, err := writer.CreateFormFile("uploadfile", fileInfo.Name())
	if err != nil {
		logger.Error("create form failed")
		return false, err
	}
	srcFile, err := os.Open(filePath)
	if err != nil {
		logger.Error("open file error")
		return false, err
	}
	defer srcFile.Close()
	_, err = io.Copy(formFile, srcFile)
	if err != nil {
		logger.Error("copy file failed")
		return false, err
	}
	contentType := writer.FormDataContentType()
	writer.Close()
	_, err = http.Post(server, contentType, bodyBuffer)
	if err != nil {
		logger.Error("upload file failed")
		return false, err
	}
	return true, nil
}

//判断服务器是否可用
func CheckServerStatus(server string) bool {
	cmd := exec.Command("ping", server, "-c", "1", "-W", "5")
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}
