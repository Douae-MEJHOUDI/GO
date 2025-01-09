package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Student struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Major   string   `json:"major"`
	Address Address  `json:"address"`
	Courses []Course `json:"courses"`
}

type Address struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
}

type Course struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Credit int    `json:"credit"`
}

type Class struct {
	students []Student
	Count    int
}

type StudentHandler struct {
	class *Class
}

func main() {

	Class := &Class{Count: 0}
	StudentHandler := &StudentHandler{class: Class}

	http.HandleFunc("/students", StudentHandler.Students)
	http.HandleFunc("/students/", StudentHandler.Student)

	fmt.Println("Server is running on http://localhost:8085")
	http.ListenAndServe(":8085", nil)

}

func (c *Class) AddStudent(s Student) {
	c.Count++
	c.students = append(c.students, s)
}

func (c *Class) GetStudents() []Student {
	return c.students
}

func (c *Class) GetStudentById(id int) (Student, error) {
	for _, s := range c.students {
		if s.Id == id {
			return s, nil
		}
	}
	return Student{}, fmt.Errorf("student not found")
}

func (c *Class) UpdateStudent(s Student) {
	for i, student := range c.students {
		if student.Id == s.Id {
			c.students[i] = s
			return
		}
	}
}

func (c *Class) DeleteStudent(id int) error {
	for i, student := range c.students {
		if student.Id == id {
			c.students = append(c.students[:i], c.students[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("student not found")
}

func (sh *StudentHandler) Students(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		//fmt.Println("this is ge request")
		sh.GetStudents(w, r)
	} else if r.Method == http.MethodPost {
		//fmt.Println("this is post request")
		sh.AddStudent(w, r)
	}

}

func (sh *StudentHandler) GetStudents(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Get students")
	students := sh.class.GetStudents()
	// maybe we can create a function that return the answers in json format (it takes as params w, r and the data as interface), it would help to avoid repeating the same code
	// JsonResponse(w, r, interface{})
	// and we can do the same for error handling (a Jsonresponse and the interface would be the error)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(students)

}

func (sh *StudentHandler) AddStudent(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("I'm adding new student")
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		//fmt.Printf("Decoding error: %v\n", err)
		return
	}
	//fmt.Println(student)
	student.Id = sh.class.Count + 1
	sh.class.AddStudent(student)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

func (sh *StudentHandler) Student(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		sh.GetStudent(w, r)
	} else if r.Method == http.MethodPut {
		sh.UpdateStudent(w, r)
	} else if r.Method == http.MethodDelete {
		sh.DeleteStudent(w, r)
	}
}

func (sh *StudentHandler) GetStudent(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) != 3 {
		http.Error(w, "Invalid url", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(paths[2])
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	student, err := sh.class.GetStudentById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}

func (sh *StudentHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) != 3 {
		http.Error(w, "Invalid url", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(paths[2])
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	student, err := sh.class.GetStudentById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var newStudent Student
	err = json.NewDecoder(r.Body).Decode(&newStudent)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	newStudent.Id = student.Id
	sh.class.UpdateStudent(newStudent)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newStudent)
}

func (sh *StudentHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(r.URL.Path, "/")
	if len(paths) != 3 {
		http.Error(w, "Invalid url", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(paths[2])
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	err = sh.class.DeleteStudent(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("student deleted successfully")
}
