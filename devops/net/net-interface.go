/*
Encoding: utf-8
Author: legend
Time: 下午9:44
*/
package net

import (
	"gitee.com/legendqi/utils-go/data"
	"gitee.com/legendqi/utils-go/logger"
	"gitee.com/legendqi/utils-go/shell"
	"net"
	"strings"
)

func GetNetworkInfo() ([]map[string]string, error) {
	inter, err := net.Interfaces()
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	interfaces := make([]map[string]string, 0)
	for _, item := range inter {
		if !strings.HasPrefix(item.Name, "docker") &&
			!strings.HasPrefix(item.Name, "veth") &&
			item.Name != "lo" {
			netTmp := map[string]string{"interface": item.Name}
			nmclis, err := shell.RunCommand("nmcli device show " + item.Name + " | grep IP4 | awk '{print $2}'")
			if err != nil {
				logger.Error(err)
			}
			netInfos := strings.Split(nmclis, "\n")
			for index, netItem := range netInfos {
				netItem = data.ClearStringUnusual(netItem)
				if index == 0 {
					if strings.Contains(netItem, "/") {
						netItem = strings.Split(netItem, "/")[0]
					}
					netTmp["ip"] = netItem
				}
				if index == 1 {
					netTmp["gateway"] = netItem
				}
				if index == len(netInfos)-2 {
					netTmp["dns"] = netItem
				}
			}
			maskCmdZh := "ifconfig " + item.Name + " | grep 掩码 | awk '{print $4}'"
			maskCmdUS := "ifconfig " + item.Name + " | grep netmask | awk '{print $4}'"
			mask, err := shell.RunCommand(maskCmdUS)
			if err != nil {
				logger.Error(err)
			}
			mask = data.ClearStringUnusual(mask)
			if mask == "" {
				mask, err = shell.RunCommand(maskCmdZh)
				if err != nil {
					logger.Error(err)
				}
				mask = data.ClearStringUnusual(mask)
				if mask == "" {
					netTmp["mask"] = "255.255.255.0"
				} else {
					maskes := strings.Split(mask, ":")
					netTmp["mask"] = maskes[1]
				}
			}
			interfaces = append(interfaces, netTmp)
		}
	}
	return interfaces, nil
}

func GetAllInterface() ([]string, error) {
	nets, err := net.Interfaces()
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	result := make([]string, 0)
	for _, item := range nets {
		if !strings.HasPrefix(item.Name, "docker") &&
			!strings.HasPrefix(item.Name, "veth") &&
			item.Name != "lo" {
			result = append(result, item.Name)
		}
	}
	return result, nil
}
