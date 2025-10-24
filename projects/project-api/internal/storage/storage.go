package storage

import "github.com/MayurakshaSikdar/golang-tutorial/project-api/internal/types"

type Storage interface {
	CreateStudent(name string, age int, email string) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetAllStudents() ([]types.Student, error)
}
