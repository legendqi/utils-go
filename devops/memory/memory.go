/*
Encoding: utf-8
Author: legend
Time: 上午1:21
*/
package memory

import (
	"gitee.com/legendqi/utils-go/logger"
	"github.com/shirou/gopsutil/mem"
)

func GetMemInfo() (map[string]int64, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return map[string]int64{
		"total":       int64(memInfo.Total),
		"available":   int64(memInfo.Available),
		"used":        int64(memInfo.Used),
		"usedPercent": int64(memInfo.UsedPercent),
		"free":        int64(memInfo.Free)}, nil
}
