package mydict

import (
	"errors"
)

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errCantUpdate = errors.New("Can't Update non-existing word")
	errWordExists = errors.New("That word already exists")
	errCantDelete = errors.New("Can't Delete non-exisiting word")
)

//Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word] //word라는 키에 해당하는 값을 반환
	if exists {
		return value, nil // 단어 있음 == 에러가 없다는 말임
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound: //단어 없으면
		d[word] = def // 단어 추가
	case nil: // 단어 있으면
		return errWordExists //단어 존재한다고 에러메시지 반환
	}
	return nil // 단어 추가하고 nil 반환
}

//Update a word to the dictionary
func (d Dictionary) Update(word, definition string) error {
	// 먼저 수정할 단어가 있나 확인
	_, err := d.Search(word)
	switch err {
	case nil: // 단어 있으면
		d[word] = definition // Update
	case errNotFound: //단어 없으면
		return errCantUpdate // Update할 수 없다고 err메시지
	}
	return nil
}

//Delete a word from the dictionary
func (d Dictionary) Delete(word string) error {
	// 먼저 삭제할 단어가 있나 확인
	_, err := d.Search(word)
	switch err {
	case nil: //단어 있으면
		delete(d, word)
	case errNotFound: //단어 없으면
		return errCantDelete //삭제할 수 없옹~
	}
	// 단어 삭제하고 nil 반환
	return nil
}

// 에러가 nil이면 뭔가 실행된다는 걸 테니까,..!
