package filter

import (
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {

	persons := []*Person{
		&Person{
			name: "jack",
			sex:  "Male",
			age:  25,
		},
		&Person{
			name: "marray",
			sex:  "Female",
			age:  21,
		},
		&Person{
			name: "xiaoming",
			sex:  "Male",
			age:  15,
		},
	}
	// 筛选女性
	female := &FemaleCriteria{}
	fmt.Println("开始筛选女性")
	for _, person := range female.meetCriteria(persons) {
		fmt.Println(*person)
	}
	fmt.Println("筛选女性结束")
	// 筛选女性
	male := &MaleCriteria{}
	fmt.Println("\n开始筛选男性")
	for _, person := range male.meetCriteria(persons) {
		fmt.Println(*person)
	}
	fmt.Println("筛选男性结束")
	// 筛选青少年
	young := &YoungCriteria{}
	fmt.Println("\n开始筛选青少年")
	for _, person := range young.meetCriteria(persons) {
		fmt.Println(*person)
	}
	fmt.Println("筛选青少年结束")

	// 筛选男性青少年
	and := &AndCriteria{
		criteria:      male,
		otherCriteria: young,
	}
	fmt.Println("\n开始筛选男性青少年")
	for _, person := range and.meetCriteria(persons) {
		fmt.Println(*person)
	}
	fmt.Println("筛选男性青少年结束")

	// 筛选男性青少年
	or := OrCriteria{
		criteria:      female,
		otherCriteria: young,
	}
	fmt.Println("\n开始筛选女性或青少年")
	for _, person := range or.meetCriteria(persons) {
		fmt.Println(*person)
	}
	fmt.Println("筛选女性或青少年结束")
}
