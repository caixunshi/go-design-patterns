package filter

type FemaleCriteria struct {
}

func (c *FemaleCriteria) meetCriteria(persons []*Person) []*Person {
	result := make([]*Person, 0, len(persons))
	for _, person := range persons {
		if person.sex == "Female" {
			result = append(result, person)
		}
	}
	return result
}
