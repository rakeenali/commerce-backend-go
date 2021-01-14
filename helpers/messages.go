package helpers

// const (
// 	ErrNotFound = "models: resource not found"
// )
var (
	ErrNotFound           publicError = "Resource not found"
	ErrInvalidToken       publicError = "Auth Token is not valid"
	ErrAccountUpdate      publicError = "Error occurred while updating account"
	ErrUserExist          publicError = "User with this username already exist"
	ErrInvalidCredentials publicError = "Invalid email or password"

	SucUserCreated    string = "User created successfully"
	SucUserLogin      string = "Login successfull"
	SucAccountUpdated string = "Account Updated"
)

// PublicError interface that implements public error
type PublicError interface {
	error
	Public() string
}

type publicError string

func (e publicError) Error() string {
	return string(e)
}

func (e publicError) Public() string {
	return string(e)
}

type privateError string

func (e privateError) Error() string {
	return string(e)
}
