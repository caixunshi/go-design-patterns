package filter

// OrCriteria 表示取两个标准的并集
type OrCriteria struct {
	criteria      Criteria
	otherCriteria Criteria
}

func (c *OrCriteria) meetCriteria(persons []*Person) []*Person {
	persons1 := c.criteria.meetCriteria(persons)
	persons2 := c.otherCriteria.meetCriteria(persons)

	resultMap := make(map[string]*Person)
	for _, person := range persons1 {
		resultMap[person.name] = person
	}
	for _, person := range persons2 {
		if resultMap[person.name] == nil {
			resultMap[person.name] = person
		}
	}
	result := make([]*Person, 0, len(resultMap))
	for _, person := range resultMap {
		result = append(result, person)
	}
	return result
}
