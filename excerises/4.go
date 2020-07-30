/**
* 要求：
* 修改程序，让其能够处理闰年情况
* 二月：闰年29天，非闰年28天
* 使用for循环生成和展示10个日期
 */

package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func leapYear(year int) bool {
	if year%100 == 0 {
		return year%400 == 0
	}
	return year%4 == 0
}

func generateDay(y int) {
	year := rand.Intn(y) + 2000
	month := rand.Intn(12) + 1
	dayInMonth := 31

	leap := year%400 == 0 || (year%4 == 0 && year%100 != 0)

	switch month {
	case 2:
		if leap {
			dayInMonth = 29
		} else {
			dayInMonth = 28
		}
	case 4, 6, 9, 11:
		dayInMonth = 30
	}

	day := rand.Intn(dayInMonth) + 1
	fmt.Println(era, year, month, day)
}

func main() {
	for i := 0; i < 10; i++ {
		r := rand.Intn(20)
		generateDay(r)
	}
}
