package query

import (
	"testing"
)

func TestNestedConditionsWorks(t *testing.T) {
	query := Builder()
	query.Where("A = ?", 1)
	sub := query.AndNested()
	sub.Where("B=?", 2)
	sub.Or("C=?", 3)

	cond, _ := query.Render()
	if cond != " WHERE (A = ? AND (B=? OR C=?))" {
		t.Fatalf("Incorrect cond; got %v", cond)
	}
}