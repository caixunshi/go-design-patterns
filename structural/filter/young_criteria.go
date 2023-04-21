package filter

type YoungCriteria struct {
}

func (c *YoungCriteria) meetCriteria(persons []*Person) []*Person {
	result := make([]*Person, 0, len(persons))
	for _, person := range persons {
		if person.age <= 18 {
			result = append(result, person)
		}
	}
	return result
}
