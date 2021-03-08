/*
Encoding: utf-8
Author: legend
Time: 上午1:21
*/
package cpu

import (
	"gitee.com/legendqi/utils-go/logger"
	"github.com/shirou/gopsutil/cpu"
	"strconv"
)

func GetCpuInfo() (map[string]string, error) {
	cpuInfo, err := cpu.Info()
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return map[string]string{"name": cpuInfo[0].ModelName, "count": strconv.Itoa(len(cpuInfo))}, nil
}
