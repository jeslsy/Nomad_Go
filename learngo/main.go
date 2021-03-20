package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		//두번의 함수는 5초뒤 ture라는 2개의 메시지를 main에 보내줌
		go isSexy(person, c)
	}
	time.Sleep(time.Second * 5)

	fmt.Println(<-c)
	fmt.Println(<-c)
}

// url 체크하고 결과를 main에 알려주는 기능(채널로) 만들어 보겠음
func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- true
}
