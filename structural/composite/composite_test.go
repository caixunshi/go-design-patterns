package composite

import (
	"fmt"
	"testing"
)

func TestComposite(t *testing.T) {

	// 销售
	headSales := Employee{"Robert", "Head Sales", 20000, nil}
	salesExecutive1 := Employee{"Richard", "Sales", 10000, nil}
	salesExecutive2 := Employee{"Rob", "Sales", 10000, nil}
	headSales.subordinates = []Employee{salesExecutive1, salesExecutive2}

	// 文员
	headMarketing := Employee{"Michel", "Head Marketing", 20000, nil}
	clerk1 := Employee{"Laura", "Marketing", 10000, nil}
	clerk2 := Employee{"Bob", "Marketing", 10000, nil}
	headMarketing.subordinates = []Employee{clerk1, clerk2}

	// 总裁
	ceo := Employee{"John", "CEO", 30000, nil}
	ceo.subordinates = []Employee{headSales, headMarketing}
	//打印该组织的所有员工
	fmt.Println(ceo)
}
