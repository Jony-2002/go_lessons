package types

type User struct {
	UserName string
	Email    string
	Password string
}

type Response struct {
	Code    int
	Message string
	Data    any
}
