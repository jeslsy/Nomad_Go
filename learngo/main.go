package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [5]string{"nico", "flynn", "dal", "japanguy", "larry"}
	for _, person := range people {
		//두번의 함수는 5초뒤 ture라는 2개의 메시지를 main에 보내줌
		go isSexy(person, c)
	}

	for i := 0; i < len(people); i++ {
		fmt.Print("waiting for", i)
		fmt.Println(<-c)
	}

}

// url 체크하고 결과를 main에 알려주는 기능(채널로) 만들어 보겠음
func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 2)
	c <- person + "is sexy"
}
