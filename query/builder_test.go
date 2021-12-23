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

func TestGroupBy(t *testing.T) {
	query := Builder()
	query.Join("JOIN Asset ON Project.ID = Asset.ProjectID").Group("Asset.ProjectID")

	cond, _ := query.Render()
	if cond != "JOIN Asset ON Project.ID = Asset.ProjectID  GROUP BY Asset.ProjectID" {
		t.Fatalf("Incorrect condition, got %v", cond)
	}
}