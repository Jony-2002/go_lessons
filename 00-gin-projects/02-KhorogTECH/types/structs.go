package types

type IsOKType struct {
	Result  bool
	Message string
}

type LoginStruct struct {
	Email    string
	Password string
}

type Registration struct {
	FirstName string
	LastName  string
	Age       int
	Address   string
	Email     string
	Password  string
}
