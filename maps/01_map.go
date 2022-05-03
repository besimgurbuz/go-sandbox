package maps

import "fmt"

func main() {
	hashMap := make(map[string]string)

	emptyValue := hashMap["key"]

	hashMap["anotherKey"] = "Some string"

	n := len(hashMap)

	value, ok := hashMap["key"]

	if ok {
		fmt.Printf("hashMap has a value in 'key' value is %v", value)
	} else {
		fmt.Printf("hashMap doesn't has value in 'key'")
	}

	fmt.Printf("Value of emptyValue label is: %v", emptyValue)
	fmt.Printf("Length of map is %d", n)

	hashMap["secondOtherKey"] = "Some other string"

	for k, val := range hashMap {
		fmt.Println("Key: ", k, " Value: ", val)
	}
}
