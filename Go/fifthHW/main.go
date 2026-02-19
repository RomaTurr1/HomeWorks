package main

import "fmt"



type Course struct {
	Id    int
	Title string
	Price int
}

type Student struct {
	Name    string
	Courses []Course
}



var catalog = make(map[int]Course)
var students = make(map[string]Student)


func AddCourse(c Course) {
	catalog[c.Id] = c
}

func AddStudent(name string) {
	students[name] = Student{
		Name:    name,
		Courses: []Course{},
	}
}

func Enroll(name string, courseId int) {
	student, ok1 := students[name]
	course, ok2 := catalog[courseId]

	if !ok1 || !ok2 {
		return
	}

	student.Courses = append(student.Courses, course)
	students[name] = student
}

func TotalCost(name string) int {
	student, ok := students[name]
	if !ok {
		return 0
	}

	sum := 0
	for _, c := range student.Courses {
		sum += c.Price
	}
	return sum
}

func PrintStudent(name string) {
	student, ok := students[name]
	if !ok {
		return
	}

	fmt.Println("Student:", student.Name)
	fmt.Println("Courses:")

	for _, c := range student.Courses {
		fmt.Println(c.Title, "-", c.Price)
	}

	fmt.Println("Total:", TotalCost(name))
	fmt.Println()
}


func main() {

	AddCourse(Course{1, "Go Basics", 20000})
	AddCourse(Course{2, "SQL", 15000})
	AddCourse(Course{3, "Web Dev", 25000})

	AddStudent("Ali")
	AddStudent("Ramazan")

	Enroll("Ali", 1)
	Enroll("Ali", 2)

	Enroll("Ramazan", 1)
	Enroll("Ramazan", 3)

	PrintStudent("Ali")
	PrintStudent("Ramazan")
}

// Practice

package main

import "fmt"

// Структуры
type Course struct {
	Id    int
	Title string
	Price int
}

type Student struct {
	Name    string
	Courses []Course
}

// Хранилища
var catalog = make(map[int]Course)
var students = make(map[string]*Student)

func main() {

	// 1️⃣ Добавляем курсы
	AddCourseToCatalog(Course{Id: 1, Title: "Go Basics", Price: 20000})
	AddCourseToCatalog(Course{Id: 2, Title: "SQL", Price: 15000})
	AddCourseToCatalog(Course{Id: 3, Title: "Python", Price: 25000})

	// 2️⃣ Регистрируем студентов
	RegisterStudent("Ali")
	RegisterStudent("Sara")

	// 3️⃣ Записываем на курсы
	Enroll("Ali", 1)
	Enroll("Ali", 2)
	Enroll("Sara", 2)
	Enroll("Sara", 3)

	// 4️⃣ Вывод информации
	PrintStudent("Ali")
	PrintStudent("Sara")
}

// Добавление курса
func AddCourseToCatalog(c Course) {
	catalog[c.Id] = c
}

// Регистрация студента
func RegisterStudent(name string) {
	students[name] = &Student{
		Name:    name,
		Courses: []Course{},
	}
}

// Запись студента на курс
func Enroll(name string, courseId int) bool {
	stud, ok1 := students[name]
	course, ok2 := catalog[courseId]

	if !ok1 || !ok2 {
		return false
	}

	stud.Courses = append(stud.Courses, course)
	return true
}

// Общая стоимость курсов студента
func TotalCost(name string) int {
	stud, ok := students[name]
	if !ok {
		return 0
	}

	total := 0
	for _, c := range stud.Courses {
		total += c.Price
	}
	return total
}

// Вывод информации о студенте
func PrintStudent(name string) {
	stud, ok := students[name]
	if !ok {
		fmt.Println("Student not found:", name)
		return
	}

	fmt.Println("Student:", stud.Name)
	fmt.Println("Courses:")
	for _, c := range stud.Courses {
		fmt.Printf("%s - %d\n", c.Title, c.Price)
	}
	fmt.Println("Total:", TotalCost(name))
	fmt.Println()
}