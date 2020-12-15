package main
import "fmt"

// implement of hash table

func main() {
	/*
		* contains key/value pairs
		* key must be unique
	*/

	
	// initialize empty map 
	var idMap1 map[string]int // here key type string and value type int
	idMap2 := make(map[string]int)
	fmt.Println(idMap1)  // map[]
	fmt.Println(idMap2)  // map[]

	
	newMap := map[string]int {
		"id": 123,
	}

	fmt.Println(len(newMap))  // 1

	fmt.Println(newMap)  // map[id:123]
	fmt.Println(newMap["id"])  // 123
	
	newMap["id2"] = 456
	fmt.Println(newMap)  // map[id:123 id2:456]

	delete(newMap, "id")
	fmt.Println(newMap)  // map[id2:456]


	// two-value assignment tests for existence of the key
	value1, key1 := newMap["id"]
	fmt.Println(value1, key1)  // 0 false

	value2, key2 := newMap["id2"]
	fmt.Println(value2, key2)  // 456 true


	// iterating trough a map
	for key, val := range newMap {
		fmt.Println(key, val)  // id2 456
	}
}