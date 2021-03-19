package main

import (
	"fmt"

	"github.com/jeslsy0507/learngo/mydict"
)

func main() {
	// map 구조체 선언은 해야되나봐
	dictionary := mydict.Dictionary{}

	baseWord := "hello"
	dictionary.Add(baseWord, "First")

	dictionary.Search(baseWord) // 단어 찾고
	dictionary.Delete(baseWord) // 단어 삭제

	word, err := dictionary.Search(baseWord)

	if err != nil { // 삭제 못했으면
		fmt.Println(err) // 삭제 실패 메시지 출력
	} else {
		fmt.Println(word) // 단어 없다고 출력
	}
}
