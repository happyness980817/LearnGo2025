package mydict

import "errors"

type Dictionary map[string]string

var (
	errNotFound = errors.New("value not found")
 	errCannotUpdate = errors.New("cannot update")
 	errWordExists = errors.New("existing word")
 ) 

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word] // map 의 key 를 호출하면 두 종류의 값을 반환. 
	// 값, ok 두 종류의 값, ok 는 값이 존재하는지 알려주는 boolean
	if exists {
		return value, nil
	}
	return "", errNotFound 
}

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

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCannotUpdate
	}
	return nil // 에러가 없음을 의미, nil 의 타입도 error
}

func (d Dictionary) Delete(word string) error {
_, err := d.Search(word)
if err != nil {
return errNotFound
}
delete(d, word)
return nil
}