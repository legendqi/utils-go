/*
Encoding: utf-8
Author: legend
Time: 上午1:19
*/
package disk

import (
	"gitee.com/legendqi/utils-go/logger"
	"gitee.com/legendqi/utils-go/shell"
	"strings"
)

func GetDiskInfo() ([]map[string]string, error) {
	diskss, err := shell.RunCommand(`df -h | awk '{print $6"##"$2"##"$3"##"$4"##"$5}'`)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	allDisk := make([]map[string]string, 0)
	for index, diskItem := range strings.Split(diskss, "\n") {
		if index > 0 && index < len(strings.Split(diskss, "\n"))-1 {
			diskItemTemp := strings.Split(diskItem, "##")
			diskTemp := map[string]string{"mount": diskItemTemp[0]}
			diskTemp["total"] = diskItemTemp[1]
			diskTemp["used"] = diskItemTemp[2]
			diskTemp["available"] = diskItemTemp[3]
			diskTemp["percent"] = diskItemTemp[4]
			allDisk = append(allDisk, diskTemp)
		}
	}
	return allDisk, nil
}
