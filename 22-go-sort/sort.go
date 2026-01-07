package main

import (
	"fmt"
	"sort"
)

type Lecture struct {
	NIDN, Name string
	Age int
}

type LectureSlice []Lecture

func (lecture LectureSlice) Len() int { return len(lecture) } 
func (lecture LectureSlice) Swap(x, y int) { lecture[x], lecture[y] = lecture[y], lecture[x] }
func (lecture LectureSlice) Less(x, y int) bool { return lecture[x].Name < lecture[y].Name }

func SortLecture() {
	lecturer := []Lecture{
		{"003124", "T. Basaruddin", 58},
		{"042354", "Malik Ibrahim", 50},
		{"012454", "Soesanto Kunto", 59},
	}

	fmt.Println("=======lecturer section=======")
	fmt.Println("before sorting:")
	for _, lecture := range lecturer {
		fmt.Println(lecture.NIDN + ".", lecture.Name, "(" + fmt.Sprint(lecture.Age) + " y.o.)")
	}

	// sort lecturer by name (default)
	sort.Sort(LectureSlice(lecturer))
	fmt.Println("after sorting by name:")
	for _, lecture := range lecturer {
		fmt.Println(lecture.NIDN + ".", lecture.Name, "(" + fmt.Sprint(lecture.Age) + " y.o.)")
	}

	// sort lecturer by nidn (custom sort using closure)
	sort.Slice(lecturer, func (x, y int) bool {
		return lecturer[x].NIDN < lecturer[y].NIDN
	})
	fmt.Println("after sorting by nidn:")
	for _, lecture := range lecturer {
		fmt.Println(lecture.NIDN + ".", lecture.Name, "(" + fmt.Sprint(lecture.Age) + " y.o.)")
	}
}