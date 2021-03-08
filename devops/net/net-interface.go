/*
Encoding: utf-8
Author: legend
Time: 下午9:44
*/
package net

import (
	"gitee.com/legendqi/utils-go/shell"
	"strings"
)

func GetInterfacefGateway(interfaceName string) (string, error) {
	gateway, err := shell.RunCommand("route | grep " + interfaceName + " | grep default | awk '{print $2}'")
	if err != nil {
		return "", err
	}
	gateway = strings.ReplaceAll(gateway, " ", "")
	gateway = strings.ReplaceAll(gateway, "\n", "")
	return gateway, nil
}
