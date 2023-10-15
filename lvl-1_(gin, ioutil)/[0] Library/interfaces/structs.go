package interfaces

type User struct {
	ID       string
	Name     string
	Surname  string
	Login    string
	Password string
}

type Login struct {
	Login    string
	Password string
}

type Employee struct {
	ID       string
	Name     string
	Surname  string
	Position string
	Salary   string
	ImageUrl string
}
