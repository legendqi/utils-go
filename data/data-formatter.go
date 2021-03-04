/* coding: utf-8
@Time :   2021/3/4 下午1:56
@Author : legend
@File :   data.go
*/
package data

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
)

func JsonToMap(str string) (map[string]interface{}, error) {
	var tempMap map[string]interface{}
	err := json.Unmarshal([]byte(str), &tempMap)
	if err != nil {
		return nil, err
	}
	return tempMap, nil
}

func JsonToMaps(str string) ([]map[string]interface{}, error) {

	var tempMap []map[string]interface{}

	err := json.Unmarshal([]byte(str), &tempMap)

	if err != nil {
		return nil, err
	}

	return tempMap, nil
}

func MapToJson(jsonMap map[string]interface{}) (string, error) {
	jsonStr, err := json.Marshal(jsonMap)
	if err != nil {
		return "", err
	}
	return string(jsonStr), nil
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

/**
16进制字符串转二进制
*/
func DecConvertToBin(input string, formatterBit int) (string, error) {
	outputStr, err := strconv.ParseInt(input, 0, 0)
	if err != nil {
		return "", err
	}
	output := int(outputStr)
	if output < 0 {
		return strconv.Itoa(output), errors.New("Only supports positive integers")
	}
	if formatterBit != 2 && formatterBit != 8 && formatterBit != 16 {
		return strconv.Itoa(formatterBit), errors.New("Only supports two, eight, hexadecimal conversion")
	}
	result := ""
	h := map[int]string{
		0:  "0",
		1:  "1",
		2:  "2",
		3:  "3",
		4:  "4",
		5:  "5",
		6:  "6",
		7:  "7",
		8:  "8",
		9:  "9",
		10: "A",
		11: "B",
		12: "C",
		13: "D",
		14: "E",
		15: "F",
	}
	for ; output > 0; output /= formatterBit {
		lsb := h[output%formatterBit]
		result = lsb + result
	}
	return result, nil
}

/*
数组反转
*/
func SliceReverse(listData []interface{}) []interface{} {

	for i, j := 0, len(listData)-1; i < j; i, j = i+1, j-1 {
		listData[i], listData[j] = listData[j], listData[i]
	}
	return listData
}

/*
字符串反转
*/
func ReverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
