package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("이름 입력 : ")
	r := bufio.NewReader(os.Stdin)
	name, err := r.ReadString('\n') //err 대신 _를 쓰면 print에 err안써도 됨

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(name)
	}

}
