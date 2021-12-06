package query

import (
	"fmt"
)

type Fragment struct {
	subconditions []*Fragment

	Condition string
	Value string
	Conjunction string
}

type Query struct {
	conditions Fragment
	joins []Fragment

	sortField string
	limitCount int
}

func Builder() *Query {
	return &Query{}
}

func (q *Query) Order(field string) *Query {
	q.sortField = field
	return q
}

func (q *Query) Limit(limit int) *Query {
	q.limitCount = limit
	return q
}

func (f *Fragment) condition(conjunction string, condition string, value interface{}) *Fragment {
	f.subconditions = append(f.subconditions, &Fragment{Condition:condition, Conjunction: conjunction, Value: fmt.Sprintf("%v", value)})
	return f
}

func (q *Query) Join(condition string) *Query {
	q.joins = append(q.joins, Fragment{Condition: condition})
	return q
}

func (f *Fragment) Or(sqlCondition string, value interface{}) *Fragment {
	return f.condition("OR", sqlCondition, value)
}

func (f *Fragment) And(sqlCondition string, value interface{}) *Fragment {
	return f.condition("AND", sqlCondition, value)
}

func (f *Fragment) Where(sqlCondition string, value interface{}) *Fragment {
	return f.condition("AND", sqlCondition, value)
}

func (q *Query) Or(sqlCondition string, value interface{}) *Query {
	q.conditions.Or(sqlCondition, value)
	return q
}

func (q *Query) And(sqlCondition string, value interface{}) *Query {
	q.conditions.And(sqlCondition, value)
	return q
}

func (q *Query) Where(sqlCondition string, value interface{}) *Query {
	q.conditions.And(sqlCondition, value)
	return q
}

func (f *Fragment) Render(conjoin bool) (condition string, fields []interface{}) {
	if conjoin {
		condition = " " + f.Conjunction + " "
	}

	if f.subconditions != nil {
		condition += "("
		for index, v := range f.subconditions {
			subConditions, subFields := v.Render(index != 0)
			condition += subConditions
			for _, f := range subFields {
				fields = append(fields, f)
			}
		}
		condition = condition + ")"
	} else {
		condition += f.Condition
		fields = append(fields, f.Value)
	}

	return
}

func (q *Query) Render() (sql string, fields []interface{}) {
	sql = ""
	for _, join := range q.joins {
		sql += join.Condition + " "
	}

	if len(q.conditions.subconditions) > 0 {
		sql += " WHERE "
		cond, condFields := q.conditions.Render(false)
		sql += cond
		for _, v := range condFields {
			fields = append(fields, v)
		}
	}

	if q.sortField != "" {
		sql += fmt.Sprintf(" ORDER BY %s", q.sortField)
	}

	if q.limitCount != 0 {
		sql += " LIMIT ?"
		fields = append(fields, q.limitCount)
	}

	return
}

func (q *Query) OrNested() (frag *Fragment) {
	frag = &Fragment{}
	q.conditions.subconditions = append(q.conditions.subconditions, frag)
	frag.Conjunction = "OR"
	return
}

func (q *Query) AndNested() (frag *Fragment) {
	frag = &Fragment{}

	q.conditions.subconditions = append(q.conditions.subconditions, frag)
	frag.Conjunction = "AND"
	return
}