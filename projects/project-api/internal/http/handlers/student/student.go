package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/MayurakshaSikdar/golang-tutorial/project-api/internal/storage"
	"github.com/MayurakshaSikdar/golang-tutorial/project-api/internal/types"
	"github.com/MayurakshaSikdar/golang-tutorial/project-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Creating a student...")
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		if err := validator.New().Struct(student); err != nil {
			validateErrors := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrors))
			return
		}
		lastId, err := storage.CreateStudent(
			student.Name,
			student.Age,
			student.Email,
		)
		slog.Info("User Created Successfully", slog.String("UserId", fmt.Sprint(lastId)))
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}
		response.WriteJson(w, http.StatusCreated, map[string]int64{"lastId": lastId})
	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("Fetching a student by ", slog.String("ID", id))
		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			slog.Error("Error ID not a valid number", slog.String("ID", id))
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		student, err := storage.GetStudentById(intId)

		if err != nil {
			slog.Error("Error Getting User", slog.String("ID", id))
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		slog.Info("Student found.", slog.String("ID", id))
		response.WriteJson(w, http.StatusOK, student)
	}
}

func GetList(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Fetching all students")

		students, err := storage.GetAllStudents()
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}
		response.WriteJson(w, http.StatusOK, students)
	}
}
