package filter

type Criteria interface {
	meetCriteria([]*Person) []*Person
}
