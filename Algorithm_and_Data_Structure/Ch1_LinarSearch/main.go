package main

import (
	"Algorithm_and_Data_Structure/Ch1_LinarSearch/Student"
	"Algorithm_and_Data_Structure/Ch1_LinarSearch/linearSerach"
	"fmt"
)

func main() {
	var data = []int{2, 4, 6, 7, 8, 9}
	res1 := linearSerach.Search[int](data, 8)
	fmt.Println(res1)

	res2 := linearSerach.Search[int](data, 1)
	fmt.Println(res2)

	var students = []Student.Student{
		Student.Student{Name: "Alice", Age: 18, Id: 1},
		Student.Student{Name: "Bob", Age: 19, Id: 2},
		Student.Student{Name: "Smith", Age: 14, Id: 3},
		Student.Student{Name: "Jack", Age: 20, Id: 4},
	}
	res3 := linearSerach.CustomSearch(students, Student.Student{Name: "Smith", Age: 14, Id: 3})
	fmt.Println(res3)

	res4 := linearSerach.CustomSearch(students, Student.Student{Name: "Matt", Age: 30, Id: 6})
	fmt.Println(res4)

}
