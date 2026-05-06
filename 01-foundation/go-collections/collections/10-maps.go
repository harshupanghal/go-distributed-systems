package collections

import "fmt"

func Maps() {
	// Built in data type (hash table), stores key value pairs

	m := make(map[string]int) // map with string keys and int values
	m["harsh"] = 22           // map_name[key] = value
	m["panghal"] = 23

	fmt.Println(m)

	// keys can be of any type that is comparable , can be strings, integers, structs, can't be slices, maps or functions

	// key must be unique in a map, if assign a values to an existing key, it will overwrite the previous value

	age := m["harsh"] // access - map_name[key]
	fmt.Println(age)

	// if the key doesn't exist , we get zero value

	// Go computes the hash of the key  to find which bucket in the hash table to look in.
	// A bucket is a small container within the hash table that holds on or more key-value pairs. when multiple keys hash to same bucket, they are stored together inside it.

	// if in search a key exists, Go returns the associated value, if not exists :
	// - 0 for int
	// "" for string
	// nil forpointer or slice

	// to distinguish between a key that doesn't exist and
	// a key whose value happens to be the zero value of the map's type
	// Go provides a second return value when we access a map.

	// value, ok := m[key]

	// ok is a boolean that is true if the key exists in the map, false if it doesn't

	m["golang"] = 0
	age2, ok2 := m["golang"]
	if !ok2 {
		fmt.Println("key not found for go lang", age2)
	}

	age1, ok1 := m["rustlang"]
	if !ok1 {
		fmt.Println("key not found for rust lang", age1)
	}

	// Iteration in Map

	for key, value := range m {
		fmt.Printf("%s : %d\n", key, value)
	}

	// range m produces each key in the map one by one
	// the loop assigns the current key to key and the corresponsing value to value

	// Iteration order in Go is randomized for Maps, each loop may produce keys in different order.
	// if need a deterministic order, collect keys in a slice and sort them and iterate over them.
	// keys are hashed to decide which bucket they go into
	// each bucket holds multiple key-value pairs
	// when a bucket gets too full, Go splits it into two (similar to dynamic resizing)
	// usaually O(1) but not guaranteed constant time.

	// Not safe for concurrent writes. If multiple goroutines write to a map at the same time, we will get runtime panuc.

	// use : 1. sync.Mutex   2. sync.RWMutex

	// to return number of key.value pairs from a map,
	fmt.Println(len(m))

	// to delete/ remove key/value pairs from a map.
	delete(m, "golang") // delete(map_name, "key")

	// to remove all key/value pairs from a map
	// clear(map_name)

}
