/* coding: utf-8
@Time :   2021/3/4 下午2:01
@Author : legend
@File :   color-print.go
*/
package utils

import "fmt"

const (
	TextBlack = iota + 30
	TextRed
	TextGreen
	TextYellow
	TextBlue
	TextMagenta
	TextCyan
	TextWhite
)

func ColorPrint(data string, color int) {
	fmt.Printf("\n %c[%d;%d;%dm%s%c[0m\n", 0x1B, 0, 0, color, data, 0x1B)
}

func InfoPrint(data string) {
	ColorPrint(data, TextGreen)
}

func ErrorPrint(data string) {
	ColorPrint(data, TextRed)
}
