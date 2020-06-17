package persistent

import (
	"database/sql"
	"github.com/rajaanova/cleancode/internal/entity"
	"context"
)

type dataProvider  struct {
	db *sql.DB
}
func NewDataprovider(db *sql.DB) *dataProvider {
	return &dataProvider{db : db }
}

func (a *dataProvider)GetStudent(ctx context.Context, id string) entity.Student {
	//TODO : need to databae implemetation
	return entity.Student{}

}

func (a *dataProvider) AddStudent(ctx context.Context,st  entity.Student) error {
	sqlQuery := "INSERT student SET name = ?,student_id = ?,age = ?,gender = ?"
	stmt, err := a.db.Prepare(sqlQuery)
	if err != nil {
		return  err

	}
	defer stmt.Close()
	_ , err = stmt.Exec(st.Name, st.Student_Id,st.Age,string(st.Gender[0]))
	return err
}