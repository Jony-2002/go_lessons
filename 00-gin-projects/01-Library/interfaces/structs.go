package interfaces

type ILogin struct {
	Login    string
	Password string
}

type IUser struct {
	FirstName string
	Surname   string
	Phone     string
	Login     string
	Password  string
}

type Book struct {
	ID     int
	Name   string
	Author string
}
