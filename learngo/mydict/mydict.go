package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("That word already exists")

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
