package main

import (
	"fmt"

	"github.com/jeslsy0507/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)

	// 항상 이렇게 에러처리 코드를 해줘야함
	if err != nil {
		fmt.Println(err)
	}

	// 위에서 추가 됐을 테니 hello 로 받아서 출력!
	hello, _ := dictionary.Search(word)
	fmt.Println("found", word, "definition:", hello)

	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println()
}
