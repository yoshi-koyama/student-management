package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"student-management/model"
)

func (q *Queries) FilterStudents(ctx context.Context, gender string, age int32, nrc string) ([]model.Student, error) {
	genderOperator := "="
	if len(gender) == 0 {
		genderOperator = "!="
	}
	ageOperator := "="
	if age < 1 {
		ageOperator = "!="
	}
	nrcOperator := "="
	if len(nrc) == 0 {
		nrcOperator = "!="
	}
	var sb strings.Builder
	sb.WriteString("select * from students where gender ")
	sb.WriteString(genderOperator)
	sb.WriteString("$1 AND ")
	sb.WriteString("age ")
	sb.WriteString(ageOperator)
	sb.WriteString("$2 AND ")
	sb.WriteString("nrc ")
	sb.WriteString(nrcOperator)
	sb.WriteString("$3 ")
	sb.WriteString("ORDER BY id")
	sqlStatement := sb.String()
	fmt.Printf("query is %s\n", sqlStatement)
	var items []model.Student
	rows, err := q.db.QueryContext(ctx, sqlStatement, gender, age, nrc)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Printf("row close error %v\n", err)
		}
	}(rows)
	for rows.Next() {
		var i model.Student
		if err := rows.Scan(
			&i.Id,
			&i.Name,
			&i.Nrc,
			&i.Age,
			&i.Gender,
			&i.Password,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	
	return items, nil
}
