/* coding: utf-8
@Time :   2021/3/4 下午1:58
@Author : legend
@File :   file.go
*/
package file_util

import (
	"bufio"
	"errors"
	"io"
	"os"
)

/*
检查文件是否存在
*/
func CheckFileIsExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

/**
读取文件返回字符串
*/
func ReadFile(path string) (string, error) {
	if !CheckFileIsExist(path) {
		return "", errors.New("file not exist")
	}
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	buffer := make([]byte, 4096)
	bufferReader := bufio.NewReader(f)
	result := ""
	for {
		n, err := bufferReader.Read(buffer)
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				break
			}
			break
		}
		//fmt.Println("current content is ===", string(buffer[:n]))
		result += string(buffer[:n])
	}
	return result, nil
}

/**
写文件，写文件之前清楚原来的内容
*/
func WriteFile(path string, content string) error {
	if !CheckFileIsExist(path) {
		return errors.New("file not exist")
	}
	f, err := os.OpenFile(path, os.O_TRUNC|os.O_WRONLY, 0766)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
