package mydict

import "errors"

//Dictionary type
type Dictionary map[string] string 

var errNotFound = errors.New("not found")
var errWordExists = errors.New("that word already exists")
var errCantUpdate = errors.New("cant update non-existing word")
//Search for a word 
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists{
		return value, nil
	}
	return "", errNotFound
}

//Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound{
		d[word] = def
	}else if err == nil{
		return errWordExists
	}
	return nil
}

//Update a word
func (d Dictionary) Update(word, def string) error{
	
	_, err := d.Search(word)
	switch err{
	case nil:
		d[word] = def
	case errNotFound:
        return errCantUpdate
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}