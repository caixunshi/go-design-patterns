package composite

import (
	"fmt"
	"testing"
)

func TestComposite(t *testing.T) {
	ceo := Employee{"John", "CEO", 30000, nil}

	// 总裁办公室
	headSales := Employee{"Robert", "Head Sales", 20000, nil}
	headMarketing := Employee{"Michel", "Head Marketing", 20000, nil}
	headSubordinates := []Employee{headSales, headMarketing}
	ceo.subordinates = headSubordinates

	// 文员
	clerk1 := Employee{"Laura", "Marketing", 10000, nil}
	clerk2 := Employee{"Bob", "Marketing", 10000, nil}
	marketingSubordinates := []Employee{clerk1, clerk2}
	headMarketing.subordinates = marketingSubordinates

	// 销售
	salesExecutive1 := Employee{"Richard", "Sales", 10000, nil}
	salesExecutive2 := Employee{"Rob", "Sales", 10000, nil}
	salesSubordinates := []Employee{salesExecutive1, salesExecutive2}
	headSales.subordinates = salesSubordinates

	//打印该组织的所有员工
	fmt.Println(ceo)
	fmt.Println(&ceo.subordinates[1] == &headMarketing)
	fmt.Println(headMarketing.subordinates)
}
