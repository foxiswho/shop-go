package auth

type User interface {
	// Return whether this user_service is logged in or not
	IsAuthenticated() bool

	// Set any flags or extra data that should be available
	Login()

	// Clear any sensitive data out of the user_service
	Logout()

	// Return the unique identifier of this user_service object
	UniqueId() interface{}

	RoleId() int

	//MORE ROLE ID
	RoleExtend() []int

	// Populate this user_service object with values
	GetById(id interface{}) error

	Module() string
}

type AuthAdmin struct {
	User
}

type AuthUser struct {
	User
}
