package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// var realScore int32
	fmt.Print("점수 입력 : ")
	in := bufio.NewReader(os.Stdin)
	score, err := in.ReadString('\n') //err 대신 _를 쓰면 print에 err안써도 됨

	if err != nil {
		log.Fatal(err)
	}
	// fmt.Print(strings.TrimSpace(score))

	score = strings.TrimSpace(score)                 //줄 바꿈, 띄어쓰기, 탭 등 제거 (python strip과 유사)
	realScore, _ := strconv.ParseInt(score, 16., 32) //정수형 32비트 타입으로 형변환

	if realScore >= 60 {

		fmt.Println("A")
		fmt.Printf("%d\n", realScore)
	} else {
		fmt.Println("BCDF")
		fmt.Printf("%d\n", realScore)
	}

}
