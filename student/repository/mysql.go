package repository

import (
	"database/sql"
	"fmt"

	"github.com/abdullahaaf/go-clean-arch-crudstudent/student/model"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlStudentRepository struct {
	Conn *sql.DB
}

func NewMysqlStudentRepository(Conn *sql.DB) StudentRepository {
	return &mysqlStudentRepository{Conn: Conn}
}

func (m *mysqlStudentRepository) GetAll() ([]*model.Student, error) {
	query := `SELECT name,registered_date,address FROM student_data`
	rows, err := m.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]*model.Student, 0)
	for rows.Next() {
		stud := new(model.Student)
		err = rows.Scan(
			&stud.Name,
			&stud.RegisteredDate,
			&stud.Address,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, stud)
	}

	return result, nil
}

func (m *mysqlStudentRepository) GetByName(name string) *model.Student {
	query := `SELECT name,registered_date,address FROM student_data WHERE name = ?`
	student := new(model.Student)
	row := m.Conn.QueryRow(query, name)
	row.Scan(
		&student.Name,
		&student.RegisteredDate,
		&student.Address,
	)

	return student
}

func (m *mysqlStudentRepository) Store(stud *model.Student) (int64, error) {
	query := `INSERT student_data SET name = ?, registered_date = ?, address = ?`
	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(stud.Name, stud.RegisteredDate, stud.Address)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (m *mysqlStudentRepository) Update(stud *model.Student) (*model.Student, error) {
	query := `UPDATE student_data SET registered_date = ?, address = ? WHERE name = ?`

	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(stud.RegisteredDate, stud.Address, stud.Name)
	if err != nil {
		return nil, err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affect != 1 {
		err := fmt.Errorf("Weird  Behaviour. Total Affected: %d", affect)
		return nil, err
	}

	return stud, nil
}

func (m *mysqlStudentRepository) Delete(name string) (bool, error) {
	query := `DELETE FROM student_data WHERE name = ?`
	stmt, err := m.Conn.Prepare(query)
	if err != nil {
		return false, nil
	}
	res, err := stmt.Exec(name)
	if err != nil {
		return false, nil
	}
	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return false, nil
	}

	if rowsAfected != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", rowsAfected)
		return false, err
	}

	return true, nil
}
