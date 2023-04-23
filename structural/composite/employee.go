package composite

type Employee struct {
	name         string
	dept         string
	salary       float32
	subordinates []Employee
}
