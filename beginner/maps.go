package main

import "fmt"

func main() {
	firstMap := make(map[string]string)

	firstMap["key1"] = "value1"

	firstMap["key2"] = "value2"

	fmt.Println("firstMap: ", firstMap)

	fmt.Println("firstMap[key1]: ", firstMap["key1"])

	anotherMap := map[string]string {"key1": "value1", "key2": "value2"}

	fmt.Println("anotherMap: ", anotherMap)
}
