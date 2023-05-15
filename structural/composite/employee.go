package composite

type Employee struct {
	name         string
	dept         string
	salary       float32
	subordinates []Employee
}

//func (e Employee) String() string {
//    fmt.Println(fmt.Sprintf("name: %s, dept: %s, salary: %.2f\n", e.name, e.dept, e.salary))
//    for _, subordinate := range e.subordinates {
//        fmt.Println(subordinate)
//    }
//    return ""
//}
