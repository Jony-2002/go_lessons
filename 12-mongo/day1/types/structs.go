package types

type User struct {
	// ID        string
	FirstName string
	LastName  string
	// Avatar    string
	Email     string
	Password  string
}

type Login struct {
	Login    string
	Password string
}
