package main

import "fmt"

type MyBoxItem struct {
	Name string
}

type MyBox struct {
	Items []MyBoxItem
}

func (box *MyBox) AddItem(item MyBoxItem) {
	box.Items = append(box.Items, item)
}

type student struct {
	id    int
	name  string
	score int
}

type option struct {
	id       int
	name     string
	quota    int
	students []student
}

type optionList struct {
	options []option
}

func (optionList *optionList) addOpt(item option) {
	optionList.options = append(optionList.options, item)
}

func (option *option) addStd(item student) {
	option.students = append(option.students, item)
}
func (option *option) removeStd(i int) {
	option.students = append(option.students[:i], option.students[i+1:]...)
}

func main() {

	item1 := MyBoxItem{Name: "Test Item 1"}
	item2 := MyBoxItem{Name: "Test Item 2"}

	box := MyBox{}

	box.AddItem(item1)
	box.AddItem(item2)

	// checking the output
	fmt.Println(len(box.Items))
	fmt.Println(box.Items)

	optionList := make([]option, 0)
	studentListA := make([]student, 0)
	studentListB := make([]student, 0)
	std1 := student{
		id: 0, name: "aa", score: 10,
	}
	std2 := student{
		id: 1, name: "ab", score: 9,
	}
	std3 := student{
		id: 2, name: "ac", score: 8,
	}
	studentListA = append(studentListA, std1)
	studentListA = append(studentListA, std2)
	studentListA = append(studentListA, std3)

	std4 := student{
		id: 0, name: "ba", score: 10,
	}
	std5 := student{
		id: 1, name: "bb", score: 9,
	}
	std6 := student{
		id: 2, name: "bc", score: 8,
	}
	studentListB = append(studentListB, std4)
	studentListB = append(studentListB, std5)
	studentListB = append(studentListB, std6)

	opt1 := option{id: 0, name: "sma1", quota: 2, students: studentListA}
	opt2 := option{id: 1, name: "sma2", quota: 1, students: studentListB}
	optionList = append(optionList, opt1)
	optionList = append(optionList, opt2)

	fmt.Println(studentListA)
	fmt.Println(studentListB)
	fmt.Println(optionList)

	for i := 0; i < len(optionList); i++ {
		fmt.Println("======", optionList[i].name, "==========")
		for j := 0; j < len(optionList[i].students); j++ {
			fmt.Println(optionList[i].students[j].name)
			if optionList[i].students[j].name == "bc" {
				//var tmpStdList []student
				//_ = append(optionList[0].students, optionList[i].students[j])
				optionList[0].addStd(optionList[i].students[j])
				optionList[1].removeStd(j)
				//fmt.Println(tmpStdList)
			}
		}
	}

	for i := 0; i < len(optionList); i++ {
		fmt.Println("======", optionList[i].name, "==========")
		for j := 0; j < len(optionList[i].students); j++ {
			fmt.Println(optionList[i].students[j].name)
		}
	}

}
