package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/MayurakshaSikdar/golang-tutorial/project-api/internal/config"
	"github.com/MayurakshaSikdar/golang-tutorial/project-api/internal/types"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INTEGER,
		email TEXT
	)`)
	if err != nil {
		return nil, err
	}
	return &Sqlite{
		Db: db,
	}, nil
}

func (s *Sqlite) CreateStudent(name string, age int, email string) (int64, error) {
	smt, err := s.Db.Prepare("INSERT INTO students (name, age, email) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer smt.Close()
	result, err := smt.Exec(name, age, email)
	if err != nil {
		return 0, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastId, nil
}

func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {
	smt, err := s.Db.Prepare("SELECT * FROM students WHERE id=? LIMIT 1")
	if err != nil {
		return types.Student{}, err
	}
	defer smt.Close()
	var student types.Student
	err = smt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Age, &student.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no student found. id : %d", id)
		}
		return types.Student{}, fmt.Errorf("query error : %w", err)
	}
	return student, nil
}

func (s *Sqlite) GetAllStudents() ([]types.Student, error) {
	smt, err := s.Db.Prepare("SELECT * FROM students")
	if err != nil {
		return []types.Student{}, err
	}
	defer smt.Close()
	var students []types.Student
	rows, err := smt.Query()
	if err != nil {
		return nil, fmt.Errorf("query error : %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var student types.Student
		err := rows.Scan(&student.Id, &student.Name, &student.Age, &student.Email)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}
