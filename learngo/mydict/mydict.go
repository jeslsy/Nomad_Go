package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Now Found")

//Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word] //word라는 키에 해당하는 값을 반환
	if exists {
		return value, nil
	}
	return "", errNotFound
}
