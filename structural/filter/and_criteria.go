package filter

// AndCriteria 表示需要同时满足两个条件
type AndCriteria struct {
	criteria      Criteria
	otherCriteria Criteria
}

func (c *AndCriteria) meetCriteria(persons []*Person) []*Person {
	persons = c.criteria.meetCriteria(persons)
	persons = c.otherCriteria.meetCriteria(persons)
	return persons
}
