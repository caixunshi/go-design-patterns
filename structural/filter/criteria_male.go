package filter

type MaleCriteria struct {
}

func (c *MaleCriteria) meetCriteria(persons []*Person) []*Person {
	result := make([]*Person, 0, len(persons))
	for _, person := range persons {
		if person.sex == "Male" {
			result = append(result, person)
		}
	}
	return result
}
