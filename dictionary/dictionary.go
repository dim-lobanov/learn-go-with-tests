package dictionary

// Acts like a wrapper for map
type Dictionary map[string]string

// Constant errors - https://dave.cheney.net/2016/04/07/constant-errors
const (
	ErrNotFound = DictionaryErr("could not find the word in dictionary")
	ErrWordAlreadyExists = DictionaryErr("word's definition already in dictionary")
	ErrWordDoesNotExist = DictionaryErr("Cannot update absent word definition")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	result, ok := d[word] // map[key]

	if !ok {
		return "", ErrNotFound
	}

	return result , nil 
}

// An interesting property of maps is that you can modify them without passing as an address to it
// (You may copy reference to 'map', but underlying data still won't be copying)

// Map will not throw an error if the value already exists. 
// Instead, they will go ahead and overwrite the value with the newly provided value.
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordAlreadyExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	// The delete built-in function deletes the element with the specified key
	// (m[key]) from the map. If m is nil or there is no such element, delete
	// is a no-op.
	delete(d, word)
	// setting value to nil won't delete key, because nil is a valid value for map
}

// You shouldn't initialize map like this
// var m map[string]string
//
// because it will be nil and this is error prone. Better:
// var dictionary = map[string]string{}
// OR
// var dictionary = make(map[string]string)
