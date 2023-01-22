package repository

import (
	"github.com/abdullahaaf/go-clean-arch-crudstudent/student/model"
	"gorm.io/gorm"
)

type mysqlStudentRepository struct {
	Conn *gorm.DB
}

func NewMysqlStudentRepository(Conn *gorm.DB) StudentRepository {
	return &mysqlStudentRepository{Conn: Conn}
}

func (m *mysqlStudentRepository) GetAll() ([]*model.Students, error) {
	var student []*model.Students
	result := m.Conn.Find(&student)
	return student, result.Error
}

func (m *mysqlStudentRepository) GetByName(name string) *model.Students {
	var student *model.Students
	_ = m.Conn.Where("name = ?", name).First(&student)

	return student
}

func (m *mysqlStudentRepository) Store(stud *model.Students) (int64, error) {
	result := m.Conn.Create(&stud)
	return result.RowsAffected, result.Error
}

func (m *mysqlStudentRepository) Update(stud *model.Students) (*model.Students, error) {
	var student model.Students
	err := m.Conn.Model(&student).Where("id = ?", stud.ID).Updates(model.Students{
		Name:           stud.Name,
		Address:        stud.Address,
		RegisteredDate: stud.RegisteredDate,
	})

	if err != nil {
		return nil, err.Error
	}

	return &student, nil
}

func (m *mysqlStudentRepository) Delete(name string) (bool, error) {
	var student model.Students
	err := m.Conn.Where("name = ?", name).Delete(&student)
	if err != nil {
		return false, err.Error
	}

	return true, nil
}
