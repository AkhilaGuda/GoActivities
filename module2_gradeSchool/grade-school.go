package main

import (
	"fmt"
	"sort"
)

type Grade struct {
	level    int
	students []string
}
type School map[int][]string

func (s School) Add(student string, grade int) {
	s[grade] = append(s[grade], student)
}

func (s School) Grade(level int) []string {
	return s[level]
}

func (s School) Enrollment() []Grade {
	var allGrades []Grade
	for level, students := range s {
		sort.Strings(students)
		allGrades = append(allGrades, Grade{level, students})
	}
	sort.Slice(allGrades, func(i, j int) bool {
		return allGrades[i].level < allGrades[j].level
	})
	return allGrades
}

func main() {
	school := School{}
	school.Add("Alice", 1)
	school.Add("Bob", 1)
	school.Add("Charlie", 2)
	school.Add("Zoe", 2)
	school.Add("Eve", 3)

	fmt.Println("Students in Grade 2:", school.Grade(2))

	fmt.Println("\nFull Enrollment:")
	for _, grade := range school.Enrollment() {
		fmt.Printf("Grade %d: %v\n", grade.level, grade.students)
	}
}
