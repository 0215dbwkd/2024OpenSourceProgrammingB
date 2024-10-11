package main

import (
	"fmt"
	"strings"
)

func main() {
	// var now time.Time = time.Now()
	// // year := now.Year()
	// // month := now.Month()
	// // day := now.Day()
	// // fmt.Println(int(year), "년", int(month), "월", int(day), "일 입니다")
	// fmt.Printf("오늘은 %d년 %d월 %d일 입니다\n", now.Year(), int(now.Month()), now.Day())
	// fmt.Printf("지금 시각은 %d시 %d분 %d초 입니다\n", now.Hour(), now.Minute(), now.Second())

	var army string = "우리는 !가와 !민에 충성을 다하는 대한민! 육!이다."
	armyFixed := strings.NewReplacer("!", "국")
	fmt.Println(army)
}
