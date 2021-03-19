package main

import (
	"fmt"

	"github.com/jeslsy0507/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil { //어떤 에러메시지 있으면
		fmt.Println(err) // 에러메시지 출력
	} else { // 에러메시지 없으면
		fmt.Println(definition) //받아온 값 출력
	}

}
