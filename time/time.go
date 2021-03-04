/* coding: utf-8
@Time :   2021/3/4 下午2:38
@Author : legend
@File :   time.go
*/
package time

import "time"

/*
 * 获取当前时间 ”YYYY-MM-DD hh:mm:ss“
 */
func GetDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

/*
 * Unix时间戳
 */
func GetUnix() int64 {
	return time.Now().Unix()
}

/*
 * 毫秒级时间戳
 */
func GetMilliUnix() int64 {
	return time.Now().UnixNano() / 1000000
}

/*
 * 纳秒级时间戳
 */
func GetNanoUnix() int64 {
	return time.Now().UnixNano()
}

func TimestampToDate(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

func DateToTimestamp(date string) (int64, error) {
	timestamp, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		return 0, err
	}
	return timestamp.Unix(), err
}
