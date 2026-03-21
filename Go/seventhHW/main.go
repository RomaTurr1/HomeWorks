package seventhhw

import "fmt"

type Participant interface {
	Info()
}

type User struct {
	name string
	age  int
}

func (u *User) SetName(name string) {
	if name == "" {
		fmt.Println("Имя не может быть пустым")
		return
	}
	u.name = name
}

func (u *User) SetAge(age int) {
	if age < 0 {
		fmt.Println("Возраст не может быть отрицательным")
		return
	}
	u.age = age
}

func (u User) GetName() string {
	return u.name
}

func (u User) GetAge() int {
	return u.age
}

type Student struct {
	User
	score    int
	MaxScore int
}

func (s *Student) AddScore(points int) {
	if s.score+points > s.MaxScore {
		s.score = s.MaxScore
	} else {
		s.score += points
	}
}

func (s Student) GetScore() int {
	return s.score
}

func (s Student) Info() {
	fmt.Println("Student:", s.name, "Age:", s.age, "Score:", s.score)
}

type Teacher struct {
	User
}

func (t Teacher) GradeStudent(student *Student, points int) {
	student.AddScore(points)
}

func (t Teacher) Info() {
	fmt.Println("Teacher:", t.name, "Age:", t.age)
}

type Course struct {
	Title       string
	MaxStudents int
	Students    []*Student
	Teacher     *Teacher
}

func (c *Course) AddStudent(s *Student) {
	if len(c.Students) >= c.MaxStudents {
		fmt.Println("Нельзя добавить больше студентов")
		return
	}
	c.Students = append(c.Students, s)
}

func (c Course) Info() {
	fmt.Println("Course:", c.Title)
	fmt.Println("Max students:", c.MaxStudents)
	fmt.Println("Students count:", len(c.Students))
}

func AverageScore(students []*Student) float64 {
	if len(students) == 0 {
		return 0
	}
	sum := 0
	for _, s := range students {
		sum += s.score
	}
	return float64(sum) / float64(len(students))
}

func BestStudent(students []*Student) *Student {
	if len(students) == 0 {
		return nil
	}
	best := students[0]
	for _, s := range students {
		if s.score > best.score {
			best = s
		}
	}
	return best
}

func main() {
	course := Course{
		Title:       "Go Course",
		MaxStudents: 3,
	}

	teacher := &Teacher{}
	teacher.SetName("Ivan")
	teacher.SetAge(40)

	s1 := &Student{MaxScore: 100}
	s1.SetName("Alex")
	s1.SetAge(20)

	s2 := &Student{MaxScore: 100}
	s2.SetName("Dima")
	s2.SetAge(21)

	s3 := &Student{MaxScore: 100}
	s3.SetName("Mira")
	s3.SetAge(22)

	course.AddStudent(s1)
	course.AddStudent(s2)
	course.AddStudent(s3)

	course.Info()

	teacher.GradeStudent(s1, 50)
	teacher.GradeStudent(s2, 70)
	teacher.GradeStudent(s3, 30)

	for _, s := range course.Students {
		s.Info()
	}

	fmt.Println("Average score:", AverageScore(course.Students))

	best := BestStudent(course.Students)
	if best != nil {
		fmt.Println("Best student:", best.GetName())
	}
}