package builder

import "fmt"

type Meal struct {
	items []Item
}

func (m *Meal) AddItem(item Item) {
	if m.items == nil {
		m.items = []Item{}
	}
	m.items = append(m.items, item)
}

func (m *Meal) ShowItems() {
	for _, item := range m.items {
		fmt.Println(fmt.Sprintf("This is a %s packed in a %s", item.name(), item.packing().pack()))
	}
}

func (m *Meal) getCost() float64 {
	money := float64(0)
	for _, item := range m.items {
		money += item.price()
	}
	return money
}
