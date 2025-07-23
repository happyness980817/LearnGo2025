package mydict

import "errors"

type Dictionary map[string]string

var errNotFound = errors.New("value not found")
var errWordExists = errors.New("existing word")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word] // map 의 key 를 호출하면 두 종류의 값을 반환. 
	// 값, ok 두 종류의 값, ok 는 값이 존재하는지 알려주는 boolean
	if exists {
		return value, nil
	}
	return "", errNotFound // can use methods inside the types in Go
} 
// we don't need the * on the receiver
// because maps in Go are automatically using *

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}