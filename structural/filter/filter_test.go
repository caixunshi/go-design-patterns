package filter

import (
	"testing"
)

func TestName(t *testing.T) {
	female := &FemaleCriteria{}
	male := &MaleCriteria{}
	young := &YoungCriteria{}
	and := &AndCriteria{
		criteria:      male,
		otherCriteria: young,
	}
	or := OrCriteria{
		criteria:      female,
		otherCriteria: young,
	}
}
