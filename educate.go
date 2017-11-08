package educate

// User represents an authenticated user in the system
type User struct {
	ID       int
	Username string
}

// UserService is an interface for interacting with users
type UserService interface {
	User(id int) (*User, error)
	Users() ([]*User, error)
	UserGroup(id int) ([]*User, error)
	CreateUser(u *User) error
	DeleteUser(id int) error
}

// Resource represents a single education resource
type Resource struct {
	ID          int
	Name        string
	Description string
	Slug        string
}

// ResourceService houses the methods for communicating with Resource elements
type ResourceService interface {
	Resource(id int) (*Resource, error)
	Resources() ([]*Resource, error)
	ResourcesGroup() ([]*Resource, error)
	CreateResource(r *Resource) error
	DeleteResource(id int) error
}
