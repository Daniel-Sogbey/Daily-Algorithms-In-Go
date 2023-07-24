package main

import "fmt"

func main() {

	a := []string{"a", "b", "c", "d"}

	b := []string{"c", "d", "e", "f"}

	fmt.Println(CommonItems(a, b))
}

func CommonItems(a []string, b []string) []string {
	hashMap := make(map[string]bool)
	var commons []string

	for _, itemInA := range a {
		hashMap[itemInA] = true
	}

	for _, itemInB := range b {
		if hashMap[itemInB] {
			commons = append(commons, []string{itemInB}...)
		}
	}

	return commons
}
