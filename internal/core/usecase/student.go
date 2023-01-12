package usecase

import (
	"github.com/abdullahaaf/go-clean-arch-crudstudent/internal/core/domain"
	"github.com/abdullahaaf/go-clean-arch-crudstudent/internal/core/ports"
)

type StudentUseCase struct {
	studentRepo ports.StudentRepository
}

func NewStudentUseCase(studrepo ports.StudentRepository) ports.StudentUseCase {
	return &StudentUseCase{studentRepo: studrepo}
}

func (stud *StudentUseCase) GetAll() ([]*domain.Student, error) {
	listStudent, err := stud.studentRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return listStudent, nil
}

func (stud *StudentUseCase) GetByName(name string) *domain.Student {
	detailStudent := stud.studentRepo.GetByName(name)
	return detailStudent
}

func (stud *StudentUseCase) Store(student *domain.Student) (int64, error) {
	studentID, err := stud.studentRepo.Store(student)
	if err != nil {
		return 0, err
	}
	student.ID = studentID
	return studentID, nil
}

func (stud *StudentUseCase) Update(student *domain.Student) (*domain.Student, error) {
	return stud.studentRepo.Update(student)
}

func (stud *StudentUseCase) Delete(name string) (bool, error) {
	return stud.studentRepo.Delete(name)
}
