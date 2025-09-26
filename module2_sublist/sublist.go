package main

import "fmt"

type Relation int

const (
	RelationEqual Relation = iota
	RelationSubList
	RelationSuperList
	RelationUnequal
)

func isSubList(l1, l2 []int) bool {
	if len(l1) == 0 {
		return true
	}
	for i := 0; i <= len(l2)-len(l1); i++ {
		match := true
		for j := 0; j < len(l1); j++ {
			if l2[j+i] != l1[j] {
				match = false
				break
			}

		}
		if match {
			return true
		}
	}
	return false

}

func sublist(l1, l2 []int) Relation {
	if len(l1) == 0 && len(l2) == 0 {
		return RelationEqual
	}
	if len(l1) > 0 && len(l2) == 0 {
		return RelationSuperList
	}
	if len(l1) == 0 && len(l2) > 0 {
		return RelationSubList
	}
	if len(l1) == len(l2) && isSubList(l1, l2) {
		return RelationEqual
	}
	if isSubList(l1, l2) {
		return RelationSubList
	}
	if isSubList(l2, l1) {
		return RelationSuperList
	}
	return RelationUnequal

}
func main() {
	l1 := []int{1, 2, 3, 4}
	l2 := []int{2, 3, 4}
	result := sublist(l1, l2)
	switch result {
	case RelationEqual:
		fmt.Println("Equal")
	case RelationSubList:
		fmt.Println("Sublist")
	case RelationSuperList:
		fmt.Println("SuperList")
	case RelationUnequal:
		fmt.Println("Unequal")
	}
}
